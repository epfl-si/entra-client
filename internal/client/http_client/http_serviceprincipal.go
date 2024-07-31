package httpengine

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"errors"
	"time"

	"io"
)

// AssociateAppRoleToServicePrincipal associates a serviceprincipal to a group and returns an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	assignment: The app role assignment
//
//	     ResourceID: The service principal ID
//	     AppRoleID: The app role ID
//	     PrincipalID: The principal ID
//
//	opts: The client options
func (c *HTTPClient) AssociateAppRoleToServicePrincipal(assignment *models.AppRoleAssignment, opts models.ClientOptions) error {
	// TODO: see https://learn.microsoft.com/en-us/entra/identity/enterprise-apps/assign-user-or-group-access-portal?pivots=ms-graph#assign-users-and-groups-to-an-application-using-microsoft-graph-api to simplify appRole selection and using default one

	c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - called\n")
	u, err := json.Marshal(assignment)
	if err != nil {
		c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - Error marshalling assignment: %+v\n", err)
		return err
	}
	c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - Assignment: %s\n", string(u))

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Post("/servicePrincipals/"+assignment.ResourceID+"/appRoleAssignments", u, h)
	if err != nil {
		c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - Error: %+v\n", response)
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 201 {
		c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - Unexpected response code: %+v\n", response)
		c.Log.Sugar().Debugf("AssociateAppRoleToServicePrincipal() - Body: %s\n", getBody(response))
		return errors.New(response.Status)
	}

	return nil
}

// AssociateClaimsPolicyToServicePrincipal associates a Claims Policy to a serviceprincipal and returns an error
//
// Required permissions: Policy.Read.All
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	claimsPolicyID: The claims policy ID
//	servicePrincipalID: The service principal ID
func (c *HTTPClient) AssociateClaimsPolicyToServicePrincipal(claimsPolicyID, servicePrincipalID string) error {
	body := []byte(`{"@odata.id":"https://graph.microsoft.com/v1.0/policies/claimsMappingPolicies/` + claimsPolicyID + `"}`)

	h := c.buildHeaders(models.ClientOptions{})
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Post("/servicePrincipals/"+servicePrincipalID+"/claimsMappingPolicies", body, h)
	defer response.Body.Close()
	if err != nil {
		c.Log.Sugar().Debugf("AssociateClaimsPolicyToServicePrincipal() - Error: %+v\n", response)
		return err
	}
	if response.StatusCode != 204 {
		c.Log.Sugar().Debugf("AssociateClaimsPolicyToServicePrincipal() - Body: %s\n", getBody(response))
		return errors.New(response.Status)
	}

	return nil
}

// CreateServicePrincipal creates an serviceprincipal and returns an error
//
// Required permissions: Application.ReadWrite.All
// Required permissions: Directory.ReadWrite.All
//
// Parameters:
//
//	 app: The service principal to be created
//		opts: The client options
func (c *HTTPClient) CreateServicePrincipal(app *models.ServicePrincipal, opts models.ClientOptions) error {
	u, err := json.Marshal(app)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Post("/serviceprincipals"+buildQueryString(opts), u, h)
	defer response.Body.Close()
	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return errors.New(response.Status)
	}

	return nil
}

// DeleteServicePrincipal deletes an serviceprincipal and returns an error
//
// Required permissions: Application.ReadWrite.All
// Required permissions: Directory.ReadWrite.All
//
// Parameters:
//
//	id: The service principal ID
//	opts: The client options
func (c *HTTPClient) DeleteServicePrincipal(id string, opts models.ClientOptions) error {
	if id == "" {
		return errors.New("ID missing")
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Delete("/serviceprincipals/"+id, h)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 204 {
		c.Log.Sugar().Debugf("DeleteServicePrincipal() - Response: %#v\n", response)
		c.Log.Sugar().Debugf("DeleteServicePrincipal() - Body: %s\n", getBody(response))

		return errors.New(response.Status)
	}

	return nil
}

// GetServicePrincipal gets an serviceprincipal by its Id and returns the serviceprincipal and an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	id: The service principal ID
//	opts: The client options
func (c *HTTPClient) GetServicePrincipal(id string, opts models.ClientOptions) (*models.ServicePrincipal, error) {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Get("/serviceprincipals/"+id+buildQueryString(opts), h)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var serviceprincipal models.ServicePrincipal
	err = json.Unmarshal(body, &serviceprincipal)
	if err != nil {
		return nil, err
	}

	return &serviceprincipal, nil
}

// GetServicePrincipals gets a list of serviceprincipals and returns a slice of serviceprincipals, a pagination link and an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetServicePrincipals(opts models.ClientOptions) ([]*models.ServicePrincipal, string, error) {
	results := make([]*models.ServicePrincipal, 0)
	var serviceprincipalResponse models.ServicePrincipalResponse
	var err error

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get("/serviceprincipals"+buildQueryString(opts), h)

	for {

		if err != nil {
			c.Log.Sugar().Debugf("GetServicePrincipals() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetServicePrincipals() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		err = json.Unmarshal(body, &serviceprincipalResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetServicePrincipals() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var serviceprincipalResponse models.ServicePrincipalResponse
		err = json.Unmarshal(body, &serviceprincipalResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetServicePrincipals() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, serviceprincipalResponse.Value...)

		if serviceprincipalResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", serviceprincipalResponse.NextLink)
		response, err = c.RestClient.Get(serviceprincipalResponse.NextLink, h)
		if response.StatusCode != 200 {
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			break
		}
	}

	return results, serviceprincipalResponse.NextLink, nil
}

// PatchServicePrincipal patches an serviceprincipal and returns an error
//
// Required permissions: Application.ReadWrite.All
// Required permissions: Directory.ReadWrite.All
//
// Parameters:
//
//	id: The service principal ID
//	app: The service principal to be patched
//	opts: The client options
func (c *HTTPClient) PatchServicePrincipal(id string, app *models.ServicePrincipal, opts models.ClientOptions) error {
	u, err := json.Marshal(app)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Patch("/servicePrincipals/"+id, u, h)
	c.Log.Sugar().Debugf("PatchServicePrincipal() - Response: %#v\n", response)
	body, err := io.ReadAll(io.Reader(response.Body))
	c.Log.Sugar().Debugf("PatchServicePrincipal() - Response: %s\n", string(body))
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

// UpdateServicePrincipal updates an serviceprincipal and returns an error
func (c *HTTPClient) UpdateServicePrincipal(app *models.ServicePrincipal, options models.ClientOptions) (err error) {
	return errors.New("not implemented")
}

// WaitServicePrincipal waits for an serviceprincipal to be created and returns an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	id: The service principal ID
//	timeout: The timeout in seconds to wait for the service principal availability in the API before returning an error
//	options: The client options
func (c *HTTPClient) WaitServicePrincipal(id string, timeout int, options models.ClientOptions) (err error) {
	duration := 0
	_, err = c.GetServicePrincipal(id, options)
	for err != nil && duration < timeout {
		time.Sleep(2 * time.Second)
		duration = duration + 2
		_, err = c.GetServicePrincipal(id, options)
		c.Log.Sugar().Debugf("WaitServicePrincipal() - Duration: %d - Error: %s\n", duration, err)
	}

	if duration >= timeout {
		return errors.New("timeout")
	}

	return nil

}