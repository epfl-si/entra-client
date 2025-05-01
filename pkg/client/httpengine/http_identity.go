package httpengine

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// AddApplicationToAuthenticationEventListeners add an application to a authenticationEventListeners object
//
// Required permissions: EventListener.ReadWrite.All
//
// Parameters:
//
//	AuthenticationEventListenersId: ID of the authenticationEventListeners
//	IncludeApplications: App id to add
//	opts: The client options
func (c *HTTPClient) AddApplicationToAuthenticationEventListeners(AuthenticationEventListenersId string, IncludeApplications *models.IdentityAuthenticationEventListenersIncludeApplicationsBody, opts models.ClientOptions) error {
	u, err := json.Marshal(IncludeApplications)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Post("/identity/authenticationEventListeners/"+AuthenticationEventListenersId+"/conditions/applications/includeApplications"+buildQueryString(opts), u, h)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// RemoveApplicationToAuthenticationEventListeners remove an application to a authenticationEventListeners object
//
// Required permissions: EventListener.ReadWrite.All
//
// Parameters:
//
//	AuthenticationEventListenersId: ID of the authenticationEventListeners
//	appId: App id to add
//	opts: The client options
func (c *HTTPClient) RemoveApplicationToAuthenticationEventListeners(AuthenticationEventListenersId string, appId string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Delete("/identity/authenticationEventListeners/"+AuthenticationEventListenersId+"/conditions/applications/includeApplications/"+appId+buildQueryString(opts), h)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusNoContent {
		return errors.New(resp.Status)
	}

	return nil
}
