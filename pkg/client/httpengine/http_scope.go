package httpengine

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// PatchRequiredResourceAccess patches an application and returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application ID
//	app: The application modifications
//	opts: The client options
func (c *HTTPClient) PatchRequiredResourceAccess(id string, app *models.Application, opts models.ClientOptions) error {

	// Filter only allowed scopes
	requiredResourceAccess, selectedScopeNames, err := c.FilterAllowedRequiredResource(app, opts)
	if err != nil {
		return err
	}

	appWithMandatoryFields := models.ApplicationWithMandatoryFields{
		Application:            *app,
		RequiredResourceAccess: requiredResourceAccess,
	}

	u, err := json.Marshal(appWithMandatoryFields)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Patch("/applications/"+id, u, h)
	if err != nil {
		return fmt.Errorf("rest Patch %s: %w", id, err)
	}
	defer response.Body.Close()

	if response.StatusCode != 204 {
		fmt.Printf("PatchRequiredResourceAccess() - Body: %#v\n", getBody(response))
		fmt.Printf("PatchRequiredResourceAccess() - Payload: %s\n", u)
		return errors.New("rest Patch " + id + " unexpected status code " + response.Status)
	}

	if len(requiredResourceAccess) > 0 {
		// Give service principal
		sp, err := c.GetServicePrincipalByAppID(app.AppID, opts)
		if err != nil {
			return err
		}

		// Give consent to the application
		err = c.PatchConsentToApplication(sp.ID, selectedScopeNames, opts)
		if err != nil {
			return err
		}
	}

	return nil
}

// FilterAllowedRequiredResourceAccess filter only allowed RequiredResource
//
// Parameters:
//
//	app: The application modifications
//	opts: The client options
func (c *HTTPClient) FilterAllowedRequiredResource(app *models.Application, opts models.ClientOptions) ([]models.RequiredResource, []string, error) {

	requiredResourceAccess := []models.RequiredResource{}
	selectedScopeNames := []string{}
	allowedScopes := map[string]map[string]string{
		c.EntraConfig.Get("MSGRAPH_API_RESOURCE_APP_ID"): {
			c.EntraConfig.Get("MSGRAPH_USER_READ_RESOURCE_ID"):      "User.Read",
			c.EntraConfig.Get("MSGRAPH_OPENID_RESOURCE_ID"):         "openid",
			c.EntraConfig.Get("MSGRAPH_PROFILE_RESOURCE_ID"):        "profile",
			c.EntraConfig.Get("MSGRAPH_EMAIL_RESOURCE_ID"):          "email",
			c.EntraConfig.Get("MSGRAPH_OFFLINE_ACCESS_RESOURCE_ID"): "offline_access",
		},
	}
	if len(app.RequiredResourceAccess) > 0 {
		for _, rr := range app.RequiredResourceAccess {
			allowedResourceScopes, hasWhitelist := allowedScopes[rr.ResourceAppID]
			if !hasWhitelist {
				requiredResourceAccess = append(requiredResourceAccess, models.RequiredResource{
					ResourceAppID:  rr.ResourceAppID,
					ResourceAccess: rr.ResourceAccess,
				})
				continue
			}
			currentRequiredResourceAccess := models.RequiredResource{
				ResourceAppID:  rr.ResourceAppID,
				ResourceAccess: []models.ResourceAccess{},
			}
			for _, ra := range rr.ResourceAccess {
				if ra.Type != "Scope" {
					continue
				}
				scopeName, scopeAllowed := allowedResourceScopes[ra.ID]
				if !scopeAllowed {
					return nil, nil, fmt.Errorf("Scope not allowed: %s (resource: %s)", ra.ID, rr.ResourceAppID)
				}
				currentRequiredResourceAccess.ResourceAccess = append(currentRequiredResourceAccess.ResourceAccess, ra)
				selectedScopeNames = append(selectedScopeNames, scopeName)
			}
			if len(currentRequiredResourceAccess.ResourceAccess) > 0 {
				requiredResourceAccess = append(requiredResourceAccess, currentRequiredResourceAccess)
			}
		}
	}
	return requiredResourceAccess, selectedScopeNames, nil
}
