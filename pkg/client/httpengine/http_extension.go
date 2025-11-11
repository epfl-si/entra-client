package httpengine

// Currently extension are defined in app 7f3a3b77-684c-447c-8a26-b18917abfed2
// (Tenant Schema Extension App)

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/epfl-si/entra-client/pkg/client/models"

	"io"
)

// CreateExtension creates a Extension and returns an error
//
// Required permissions: Directory.Read.All
// Required permissions: ExtensionMember.Read.All
//
// Parameters:
//
//	extension: The extension to be created
//	opts: The client options
func (c *HTTPClient) CreateExtension(extension *models.ExtensionProperty, opts models.ClientOptions) error {
	u, err := json.Marshal(extension)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	resp, err := c.RestClient.Post("/extensions/"+buildQueryString(opts), u, h)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// DeleteExtension deletes a extension by its Id and returns an error
//
// Required permissions: Extension.ReadWrite.All
//
// Parameters:
//
//	id: The extension ID
//	opts: The client options
func (c *HTTPClient) DeleteExtension(id string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Delete("/extensions/"+id, h)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

// GetExtension returns a extension (by its Id) and an error
//
// Required permissions: Directory.Read.All
// Required permissions: ExtensionMember.Read.All
//
// Parameters:
//
//	id: The extension ID
//	opts: The client options
func (c *HTTPClient) GetExtension(id string, opts models.ClientOptions) (*models.ExtensionProperty, error) {
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Get("/extensions/"+id+buildQueryString(opts), h)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		c.Log.Sugar().Debugf("GetExtension() - QueryString: %s\n", buildQueryString(opts))
		c.Log.Sugar().Debugf("GetExtension() - Response body: %s\n", getBody(response))
		return nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var extension models.ExtensionProperty
	err = json.Unmarshal(body, &extension)
	if err != nil {
		return nil, err
	}

	return &extension, nil
}

// GetExtensions gets all extensions and returns a slice of extensions and an error
//
// Required permissions: Directory.Read.All
// Required permissions: ExtensionMember.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetExtensions(opts models.ClientOptions) ([]*models.ExtensionProperty, error) {
	results := make([]*models.ExtensionProperty, 0)

	var response *http.Response
	var err error

	h := c.buildHeaders(opts)

	response, err = c.RestClient.Post("/directoryObjects/getAvailableExtensionProperties", []byte("{}"), h)
	if err != nil {
		c.Log.Sugar().Debugf("GetExtensions - POST Error: %s\n", err.Error())
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		c.Log.Sugar().Debugf("GetExtensions - Error: %s\n", err.Error())
		return nil, err
	}

	response.Body.Close()

	var extensionPropertyResponse models.ExtensionPropertyResponse
	err = json.Unmarshal(body, &extensionPropertyResponse)
	if err != nil {
		c.Log.Sugar().Debugf("GetExtensions - Error: %s\n", err.Error())
		return nil, err
	}

	results = append(results, extensionPropertyResponse.Value...)

	return results, nil
}

// UpdateExtension updates a extension and returns an error
//
// Required permissions: ???????
//
// Parameters:
//
//	extension: The extension to be updated
//	opts: The client options
func (c *HTTPClient) UpdateExtension(extension *models.ExtensionProperty, opts models.ClientOptions) error {
	u, err := json.Marshal(extension)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Put("/extensions/"+extension.ID, u, h)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	return nil
}
