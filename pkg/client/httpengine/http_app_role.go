package httpengine

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// CreateAppRole creates an application role and returns an error
//
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	appID: The application ID where the role will be created
//	appRole: The application role to be created
//	opts: The client options
func (c *HTTPClient) CreateAppRoleByAppID(appID string, appRole *models.AppRole, opts models.ClientOptions) error {
	// Get the application by its appID using c.RestClient
	app, err := c.GetApplicationByAppID(appID, opts)
	if err != nil {
		return err
	}

	// Append the new appRole to the existing AppRoles
	app.AppRoles = append(app.AppRoles, *appRole)

	// Marshal the updated application back to JSON
	u, err := json.Marshal(map[string]interface{}{
		"appRoles": app.AppRoles,
	})
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Patch("/applications/"+app.ID+buildQueryString(opts), u, h)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(resp.Status)
	}

	return nil
}
