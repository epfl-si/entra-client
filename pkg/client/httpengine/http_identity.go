package httpengine

import (
	"encoding/json"
	"errors"
	"io"
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
func (c *HTTPClient) CreateAuthenticationEventListeners(onTokenIssuanceStartListener *models.OnTokenIssuanceStartListener, opts models.ClientOptions) (*models.AuthenticationEventListener, error) {
	u, err := json.Marshal(onTokenIssuanceStartListener)
	if err != nil {
		return nil, err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Post("/identity/authenticationEventListeners"+buildQueryString(opts), u, h)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusCreated {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Log.Sugar().Debugf("CreateAuthenticationEventListeners() - Body read error: %s\n", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	var ael *models.AuthenticationEventListener
	err = json.Unmarshal(body, &ael)
	c.Log.Sugar().Debugf("CreateAuthenticationEventListeners() - Body: %s\n", string(body))
	if err != nil {
		c.Log.Sugar().Debugf("CreateAuthenticationEventListeners() - Response unmarshall error: %s\n", err.Error())
		return nil, err
	}

	return ael, nil
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
