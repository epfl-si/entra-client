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
	h["Content-Type"] = "application/json"

	resp, err := c.RestClient.Post("/identity/authenticationEventListeners"+buildQueryString(opts), u, h)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Log.Sugar().Debugf("CreateAuthenticationEventListeners() - Body read error: %s\n", err.Error())
		return nil, err
	}

	c.Log.Sugar().Debugf("CreateAuthenticationEventListeners() - Response: %s", string(body))

	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	var ael *models.AuthenticationEventListener
	if err := json.Unmarshal(body, &ael); err != nil {
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
	defer resp.Body.Close()

	// Expect 200 OK or 204 No Content for successful updates
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Body read error: %s\n", err.Error())
		return nil, err
	}
	var ael *models.AuthenticationEventListener
	err = json.Unmarshal(body, &ael)
	c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Body: %s\n", string(body))

	if err != nil {
		c.Log.Sugar().Debugf("GetAuthenticationEventListener() - Response unmarshall error: %s\n", err.Error())
		return nil, err
	}

	return ael, nil
}

// IsApplicationInAuthenticationEventListener checks whether an application is in an authentication event listener's include list
//
// Required permissions: EventListener.Read.All or EventListener.ReadWrite.All
//
// Parameters:
//
//	listenerID: The ID of the authentication event listener
//	appID: The application ID to check
//	opts: The client options
func (c *HTTPClient) IsApplicationInAuthenticationEventListener(listenerID string, appID string, opts models.ClientOptions) (bool, error) {
	h := c.buildHeaders(opts)

	endpoint := fmt.Sprintf("/identity/authenticationEventListeners/%s/conditions/applications/includeApplications/%s", listenerID, appID)

	resp, err := c.RestClient.Get(endpoint, h)
	if err != nil {
		return false, fmt.Errorf("failed to check application in listener: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		return true, nil
	}

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	}

	return false, fmt.Errorf("unexpected status code: %s", resp.Status)
}

// GetAuthenticationEventListeners retrieves authentication event listeners matching the given options.
//
// Required permissions: EventListener.Read.All or EventListener.ReadWrite.All
// Parameters:
//
//	opts: The client options containing query
func (c *HTTPClient) GetAuthenticationEventListeners(opts models.ClientOptions) ([]models.AuthenticationEventListener, error) {
	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Get("/identity/authenticationEventListeners"+buildQueryString(opts), h)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	var listResp models.AuthenticationEventListenerListResponse
	if err := json.Unmarshal(body, &listResp); err != nil {
		return nil, fmt.Errorf("failed to unmarshal response: %w", err)
	}

	return listResp.Value, nil
}

// DeleteAuthenticationEventListener deletes an authentication event listener by ID.
//
// Required permissions: EventListener.ReadWrite.All
// Parameters:
//
//	listenerID: The ID of the authentication event listener
//	opts: The client options
func (c *HTTPClient) DeleteAuthenticationEventListener(listenerID string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Delete("/identity/authenticationEventListeners/"+listenerID, h)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("%s: %s", resp.Status, string(body))
	}

	return nil
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
	exists, err := c.IsApplicationInAuthenticationEventListener(listenerID, appID, opts)
	if err != nil {
		return fmt.Errorf("failed to check if application exists in listener: %w", err)
	}
	if exists {
		c.Log.Sugar().Debugf("AddApplicationToAuthenticationEventListener() - App %s already exists in listener %s", appID, listenerID)
		return nil
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

// RemoveApplicationFromAuthenticationEventListener removes an application from an authentication event listener's include list
//
// Required permissions: EventListener.ReadWrite.All
//
// Parameters:
//
//	listenerID: The ID of the authentication event listener
//	appID: The application ID to remove from the listener's include applications list
//	opts: The client options
func (c *HTTPClient) RemoveApplicationFromAuthenticationEventListener(listenerID string, appID string, opts models.ClientOptions) error {
	exists, err := c.IsApplicationInAuthenticationEventListener(listenerID, appID, opts)
	if err != nil {
		return fmt.Errorf("failed to check if application exists in listener: %w", err)
	}
	if !exists {
		c.Log.Sugar().Debugf("RemoveApplicationFromAuthenticationEventListener() - App %s not found in listener %s, nothing to remove", appID, listenerID)
		return nil
	}

	h := c.buildHeaders(opts)

	endpoint := fmt.Sprintf("/identity/authenticationEventListeners/%s/conditions/applications/includeApplications/%s", listenerID, appID)

	resp, err := c.RestClient.Delete(endpoint, h)
	if err != nil {
		return fmt.Errorf("failed to remove application: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("unexpected status code %s", resp.Status)
	}

	c.Log.Sugar().Debugf("RemoveApplicationFromAuthenticationEventListener() - Successfully removed app %s from listener %s", appID, listenerID)
	return nil
}
