package httpengine

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// CreateAuthenticationEventListeners create an authentication event listener
//
// Required permissions: EventListener.ReadWrite.All
//
// Parameters:
//
//	onTokenIssuanceStartListener: Content of the authenticationEventListener object
//	opts: The client options
func (c *HTTPClient) CreateAuthenticationEventListeners(onTokenIssuanceStartListener *models.OnTokenIssuanceStartListener, opts models.ClientOptions) error {
	u, err := json.Marshal(onTokenIssuanceStartListener)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Post("/identity/authenticationEventListeners"+buildQueryString(opts), u, h)
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
