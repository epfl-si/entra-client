package httpengine

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/epfl-si/entra-client/pkg/client/models"

	"io"
)

// Ressources:
// https://learn.microsoft.com/en-us/entra/identity-platform/reference-claims-customization
// https://learn.microsoft.com/en-us/graph/api/resources/claimsmappingpolicyschema?view=graph-rest-1.0

// DefaultClaimsMappingPolicy is the default claims mapping policy for EPFL
//
//	uniqueid -> employeeId (Sciper)
//	gaspar -> onPremisesSamAccountName (Gaspar username)
//	given_name -> givenName (First name)
//	family_name -> surname (Last name)
//	email -> mail (Email)
var DefaultClaimsMappingPolicy = []string{
	`{	"ClaimsMappingPolicy": { "Version":1, "IncludeBasicClaimSet":"false", "ClaimsSchema":[ 
{"Source":"user", "ID": "employeeId", "JwtClaimType": "uniqueid"},
{"Source":"user", "ID": "onPremisesSamAccountName", "JwtClaimType": "gaspar"},
{"Source":"user", "ID": "givenName", "JwtClaimType": "given_name"},
{"Source":"user", "ID": "surname", "JwtClaimType": "family_name"},
{"Source":"user", "ID": "mail", "JwtClaimType": "mail"}
]}}`}

//{"Source":"user", "ID": "groups", "JwtClaimType": "groups"}

// AssignClaimsMappingPolicy assigns a claims mapping policy and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.ReadWrite.ApplicationConfiguration
//
// Parameters:
//
//	claimspolicy: The claims mapping policy to assign
//	opts: The client options
func (c *HTTPClient) AssignClaimsMappingPolicy(cmpID, spID string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	payload := []byte(`
{
    "@odata.id": "https://graph.microsoft.com/v1.0/policies/claimsMappingPolicies/` + cmpID + `"
}
	`)

	response, err := c.RestClient.Post(fmt.Sprintf("/servicePrincipals/%s/claimsMappingPolicies/$ref", spID), payload, h)
	if err != nil {
		c.Log.Sugar().Debugf("AssignClaimsMappingPolicy() - Body: %+v\n", response)
		return err
	}

	if response.StatusCode != 204 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.Log.Sugar().Debugf("AssignClaimsMappingPolicy() - Error reading response body: %s\n", err.Error())
			return err
		}
		response.Body.Close()
		c.Log.Sugar().Debugf("AssignClaimsMappingPolicy() - Body: %s\n", string(body))
		return errors.New(response.Status)
	}

	return nil
}

// CreateClaimsMappingPolicy creates a claims mapping policy and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.ReadWrite.ApplicationConfiguration
//
// Parameters:
//
//	claimspolicy: The claims mapping policy to create
//	opts: The client options
func (c *HTTPClient) CreateClaimsMappingPolicy(claimspolicy *models.ClaimsMappingPolicy, opts models.ClientOptions) (string, error) {
	if opts.Default {
		claimspolicy.Definition = DefaultClaimsMappingPolicy
		claimspolicy.DisplayName = "EPFL Default Claims Policy"
		claimspolicy.IsOrganizationDefault = true
	}

	u, err := json.Marshal(claimspolicy)
	if err != nil {
		c.Log.Sugar().Debugf("CreateClaimsMappingPolicy() - Error marshalling claims: %s\n", err.Error())
		return "", err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	// response, err := c.RestClient.Post("/claimsmappingpolicies"+buildQueryString(opts), u, h)
	response, err := c.RestClient.Post("/policies/claimsmappingpolicies", u, h)
	if err != nil {
		c.Log.Sugar().Debugf("CreateClaimsMappingPolicy() - Body: %+v\n", response)
		return "", err
	}
	if response.StatusCode != 201 {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.Log.Sugar().Debugf("DeleteClaimsMappingPolicy() - Error reading response body: %s\n", err.Error())
			return "", err
		}
		c.Log.Sugar().Debugf("CreateClaimsMappingPolicy() - Body: %s\n", string(body))
		return "", errors.New(response.Status)
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return "", err
	}

	response.Body.Close()

	var claims models.ClaimsMappingPolicy
	err = json.Unmarshal(body, &claims)
	if err != nil {
		c.Log.Sugar().Debugf("CreateClaimsMappingPolicy() - Body: %s\n", string(body))
		return "", err
	}

	if opts.Debug {
		c.Log.Sugar().Debugf("CreateClaimsMappingPolicy() - Response: %+v\n", claims)
	}

	return claims.ID, nil
}

// DeleteClaimsMappingPolicy deletes a claims mapping policy and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.ReadWrite.ApplicationConfiguration
//
// Parameters:
//
//	id: The application ID
//	opts: The client options
func (c *HTTPClient) DeleteClaimsMappingPolicy(id string, opts models.ClientOptions) error {
	if id == "" {
		return errors.New("ID missing")
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Delete("/policies/claimsmappingPolicies/"+id, h)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		c.Log.Sugar().Debugf("DeleteClaimsMappingPolicy() - Response: %#v\n", response)
		body, err := io.ReadAll(response.Body)
		if err != nil {
			c.Log.Sugar().Debugf("UnassignClaimsMappingPolicy() - Error reading response body: %s\n", err.Error())
			return err
		}
		c.Log.Sugar().Debugf("DeleteClaimsMappingPolicy() - Body: %s\n", string(body))

		return errors.New(response.Status)
	}

	return nil
}

// GetClaimsMappingPolicy gets a claims mapping policy  by its Id and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.Read.All
//
// Parameters:
//
//	id: The application ID
//	opts: The client options
func (c *HTTPClient) GetClaimsMappingPolicy(cmpID string, opts models.ClientOptions) (*models.ClaimsMappingPolicy, error) {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Get("/policies/claimsmappingpolicies/"+cmpID+buildQueryString(opts), h)
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

	var cmp models.ClaimsMappingPolicy
	err = json.Unmarshal(body, &cmp)
	if err != nil {
		return nil, err
	}

	return &cmp, nil
}

// GetClaimsMappingPolicies gets all claims mapping policies and returns a slice of claims mapping policies, a pagination link and an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetClaimsMappingPolicies(opts models.ClientOptions) ([]*models.ClaimsMappingPolicy, string, error) {
	c.Log.Sugar().Debugf("GetClaimsMappingPolicies() - Started\n")
	results := make([]*models.ClaimsMappingPolicy, 0)
	var claimsMappingPolicyResponse models.ClaimsMappingPolicyResponse
	var err error

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get("/policies/claimsmappingPolicies"+buildQueryString(opts), h)

	for {
		if err != nil {
			c.Log.Sugar().Debugf("GetClaimsMappingPolicies() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetClaimsMappingPolicies() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		err = json.Unmarshal(body, &claimsMappingPolicyResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetClaimsMappingPolicies() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var cmpResponse models.ClaimsMappingPolicyResponse
		err = json.Unmarshal(body, &cmpResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetClaimsMappingPolicys() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, cmpResponse.Value...)

		if cmpResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetClaimsMappingPolicys() - 4 - Calling Next: %s\n", cmpResponse.NextLink)
		response, err = c.RestClient.Get(cmpResponse.NextLink, h)
		if response.StatusCode != 200 {
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			break
		}
	}

	return results, claimsMappingPolicyResponse.NextLink, nil
}

// PatchClaimsMappingPolicy patches a claims mapping policy and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.ReadWrite.ApplicationConfiguration
//
// Parameters:
//
//	cmid: The claims mapping policy ID
//	app: The claims mapping policy modification
//	opts: The client options
func (c *HTTPClient) PatchClaimsMappingPolicy(cmpid string, cmp *models.ClaimsMappingPolicy, opts models.ClientOptions) error {
	if opts.Default {
		cmp.Definition = DefaultClaimsMappingPolicy
		cmp.DisplayName = "EPFL Default Claims Policy"
		cmp.IsOrganizationDefault = true
	}

	// c.Log.Sugar().Debugf("PatchClaimsMappingPolicy() - cmp: %#v\n", cmp)

	u, err := json.Marshal(cmp)
	if err != nil {
		return err
	}
	// c.Log.Sugar().Debugf("PatchClaimsMappingPolicy() - Payload: %s\n", string(u))

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Patch("/policies/claimsmappingpolicies/"+cmpid, u, h)
	c.Log.Sugar().Debugf("PatchClaimsMappingPolicy() - Response: %#v\n", response)
	body, err := io.ReadAll(io.Reader(response.Body))
	c.Log.Sugar().Debugf("PatchClaimsMappingPolicy() - Body: %s\n", string(body))
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

// UpdateClaimsMappingPolicy updates a claims mapping policy and returns an error
// func (c *HTTPClient) UpdateClaimsMappingPolicy(app *models.ClaimsMappingPolicy, options models.ClientOptions) (err error) {
// 	return errors.New("not implemented")
// }

// WaitClaimsMappingPolicy waits for a claims mapping policy to be created and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.Read.All
//
// Parameters:
//
//	id: The claims mapping policy ID
//	timeout: The timeout in seconds to wait for the claims mapping policy before returning an error
//	opts: The client options
func (c *HTTPClient) WaitClaimsMappingPolicy(id string, timeout int, options models.ClientOptions) (err error) {
	duration := 0
	_, err = c.GetClaimsMappingPolicy(id, options)
	for err != nil && duration < timeout {
		time.Sleep(2 * time.Second)
		duration = duration + 2
		_, err = c.GetClaimsMappingPolicy(id, options)
		c.Log.Sugar().Debugf("WaitClaimsMappingPolicy() - Duration: %d - Error: %s\n", duration, err)
	}

	if duration >= timeout {
		return errors.New("timeout")
	}

	return nil

}

// ListUsageClaimsMappingPolicy gets all the usage of given claims mapping policies
//
//	returns a slice of claims mapping policies, a pagination link and an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) ListUsageClaimsMappingPolicy(cmpID string, opts models.ClientOptions) ([]*models.ClaimsMappingPolicy, error) {
	c.Log.Sugar().Debugf("GetClaimsMappingPolicys() - Started\n")
	results := make([]*models.ClaimsMappingPolicy, 0)
	var claimsMappingPolicyResponse models.ClaimsMappingPolicyResponse
	var err error

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get(fmt.Sprintf("/policies/claimsMappingPolicies/%s/appliesTo%v", cmpID, buildQueryString(opts)), h)

	for {
		if err != nil {
			c.Log.Sugar().Debugf("ListUsageClaimsMappingPolicy() - 1 - Error: %s\n", err.Error())
			return nil, err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("ListUsageClaimsMappingPolicy() - 2 - Error: %s\n", err.Error())
			return nil, err
		}

		err = json.Unmarshal(body, &claimsMappingPolicyResponse)
		if err != nil {
			c.Log.Sugar().Debugf("ListUsageClaimsMappingPolicy() - 3 - Error: %s\n", err.Error())
			return nil, err
		}

		response.Body.Close()

		var cmpResponse models.ClaimsMappingPolicyResponse
		err = json.Unmarshal(body, &cmpResponse)
		if err != nil {
			c.Log.Sugar().Debugf("ListUSageClaimsMappingPolicy() - 3 - Error: %s\n", err.Error())
			return nil, err
		}

		results = append(results, cmpResponse.Value...)

		if cmpResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("ListUsageClaimsMappingPolicy() - 4 - Calling Next: %s\n", cmpResponse.NextLink)
		response, err = c.RestClient.Get(cmpResponse.NextLink, h)
		if response.StatusCode != 200 {
			return nil, errors.New(response.Status)
		}
	}

	return results, nil
}

// UnassignClaimsMappingPolicy unassigns a claims mapping policy and returns an error
//
// Required permissions: Policy.Read.ApplicationConfiguration
// Required permissions: Policy.ReadWrite.ApplicationConfiguration
//
// Parameters:
//
//	claimspolicy: The claims mapping policy to assign
//	opts: The client options
func (c *HTTPClient) UnassignClaimsMappingPolicy(spID, cmpID string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)

	response, err := c.RestClient.Delete(fmt.Sprintf("/servicePrincipals/%s/claimsMappingPolicies/%s/$ref", spID, cmpID), h)
	if err != nil {
		c.Log.Sugar().Debugf("UnassignClaimsMappingPolicy() - Body: %+v\n", response)
		return err
	}

	if response.StatusCode != 204 {
		body, _ := io.ReadAll(io.Reader(response.Body))
		response.Body.Close()
		c.Log.Sugar().Debugf("UnassignClaimsMappingPolicy() - Body: %s\n", string(body))
		return errors.New(response.Status)
	}

	return nil
}
