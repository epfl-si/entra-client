package httpengine

import (
	"encoding/json"
	"errors"
	"fmt"
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

// GetAuthenticationEventListener retrieves an authentication event listener by ID
//
// Required permissions: EventListener.Read.All or EventListener.ReadWrite.All
//
// Parameters:
//
//	listenerID: The ID of the authentication event listener to retrieve
//	opts: The client options
func (c *HTTPClient) GetAuthenticationEventListener(listenerID string, opts models.ClientOptions) (*models.AuthenticationEventListener, error) {
	h := c.buildHeaders(opts)

	// Use PATCH to update the conditions
	resp, err := c.RestClient.Get("/identity/authenticationEventListeners/"+listenerID, h)
	if err != nil {
		return nil, err
	}

	// Expect 200 OK or 204 No Content for successful updates
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Body read error: %s\n", err.Error())
		return nil, err
	}

	defer resp.Body.Close()
	var ael *models.AuthenticationEventListener
	err = json.Unmarshal(body, &ael)
	c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Body: %s\n", string(body))

	if err != nil {
		c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Response unmarshall error: %s\n", err.Error())
		return nil, err
	}

	return ael, nil
}

// AddApplicationToAuthenticationEventListener adds an application to an authentication event listener's include list
//
// Required permissions: EventListener.ReadWrite.All
//
// Parameters:
//
//	listenerId: The ID of the authentication event listener
//	appId: The application ID to add to the listener's include applications list
//	opts: The client options
func (c *HTTPClient) AddApplicationToAuthenticationEventListener(listenerID string, appID string, opts models.ClientOptions) error {
	currentListener, err := c.GetAuthenticationEventListener(listenerID, opts)
	if err != nil {
		return fmt.Errorf("failed to get current listener: %w", err)
	}

	if currentListener.Conditions != nil &&
		currentListener.Conditions.Applications.IncludeApplications != nil {

		for _, existingApp := range currentListener.Conditions.Applications.IncludeApplications {
			if existingApp.AppId == appID {
				c.Log.Sugar().Debugf("AddApplicationToAuthenticationEventListener() - App %s already exists in listener %s", appID, listenerID)
				return nil // App already exists, no need to add
			}
		}
	}

	requestBody := map[string]string{
		"@odata.type": "#microsoft.graph.authenticationConditionApplication",
		"appId":       appID,
	}

	u, err := json.Marshal(requestBody)
	if err != nil {
		return fmt.Errorf("failed to marshal request body: %w", err)
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	// Use the specific endpoint for adding applications
	endpoint := fmt.Sprintf("/identity/authenticationEventListeners/%s/conditions/applications/includeApplications", listenerID)

	resp, err := c.RestClient.Post(endpoint, u, h)
	if err != nil {
		return fmt.Errorf("failed to add application: %w", err)
	}

	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	c.Log.Sugar().Debugf("AddApplicationToAuthenticationEventListener() - Response: %s", string(body))

	// This endpoint returns 201 Created on success
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %s", resp.Status)
	}

	c.Log.Sugar().Debugf("AddApplicationToAuthenticationEventListener() - Successfully added app %s to listener %s", appID, listenerID)
	return nil
}
