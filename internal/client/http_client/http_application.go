package httpengine

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"errors"
	"time"

	"io"
)

// AddPasswordToApplication adds a password/secret to an application and returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application ID
//	keyName: The password nane
//	opts: The client options
func (c *HTTPClient) AddPasswordToApplication(id, keyName string, opts models.ClientOptions) (*models.PasswordCredential, error) {
	if id == "" {
		return nil, errors.New("ID missing")
	}

	if keyName == "" {
		return nil, errors.New("key name missing")
	}

	pc := &models.PasswordCredential{
		DisplayName: keyName,
	}

	u, err := json.Marshal(pc)
	if err != nil {
		return nil, err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Post("/applications/"+id+"/addPassword", u, h)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	body := getBody(response)

	var result models.PasswordCredential
	err = json.Unmarshal([]byte(body), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateApplication creates an application and returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	app: The application to create
//	opts: The client options
func (c *HTTPClient) CreateApplication(app *models.Application, opts models.ClientOptions) (*models.Application, error) {
	u, err := json.Marshal(app)
	if err != nil {
		return nil, err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Post("/applications", u, h)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 201 {
		c.Log.Sugar().Debugf("CreateApplication() - Response: %s\n", getBody(response))
		return nil, errors.New(response.Status)
	}

	var resultApp models.Application
	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &resultApp)
	if err != nil {
		c.Log.Sugar().Debugf("CreateApplication() - Body: %s\n", string(body))
		return nil, err
	}

	return &resultApp, nil
}

// DeleteApplication deletes an application and returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application ID
//	opts: The client options
func (c *HTTPClient) DeleteApplication(id string, opts models.ClientOptions) error {
	if id == "" {
		return errors.New("ID missing")
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Delete("/applications/"+id, h)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		c.Log.Sugar().Debugf("DeleteApplication() - Response: %#v\n", response)
		body, _ := io.ReadAll(io.Reader(response.Body))
		c.Log.Sugar().Debugf("DeleteApplication() - Body: %s\n", string(body))

		return errors.New(response.Status)
	}

	return nil
}

// GetApplication gets an application by its Id and returns an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application ID
//	opts: The client options
func (c *HTTPClient) GetApplication(id string, opts models.ClientOptions) (*models.Application, error) {
	if id == "" {
		return nil, errors.New("ID missing")
	}
	h := c.buildHeaders(opts)
	response, err := c.RestClient.Get("/applications/"+id+buildQueryString(opts), h)
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

	var application models.Application
	err = json.Unmarshal(body, &application)
	if err != nil {
		c.Log.Sugar().Debugf("GetApplication() - Body: %s\n", string(body))
		return nil, err
	}

	return &application, nil
}

// GetApplications returns all applications and an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetApplications(opts models.ClientOptions) ([]*models.Application, string, error) {
	results := make([]*models.Application, 0)
	var applicationResponse models.ApplicationResponse
	var err error

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get("/applications"+buildQueryString(opts), h)

	for {

		if err != nil {
			c.Log.Sugar().Debugf("GetApplications() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetApplications() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		err = json.Unmarshal(body, &applicationResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetApplications() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var applicationResponse models.ApplicationResponse
		err = json.Unmarshal(body, &applicationResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetApplications() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, applicationResponse.Value...)

		if applicationResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", applicationResponse.NextLink)
		response, err = c.RestClient.Get(applicationResponse.NextLink, h)
		if response.StatusCode != 200 {
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			break
		}
	}

	return results, applicationResponse.NextLink, nil
}

// PatchApplication patches an application and returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application ID
//	app: The application modifications
//	opts: The client options
func (c *HTTPClient) PatchApplication(id string, app *models.Application, opts models.ClientOptions) error {
	u, err := json.Marshal(app)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Patch("/applications/"+id, u, h)
	if err != nil {
		return err
	}

	if response.StatusCode != 204 {
		c.Log.Sugar().Debugf("PatchApplication() - Response: %#v\n", response)
		c.Log.Sugar().Debugf("PatchApplication() - Body: %s\n", getBody(response))
		return errors.New(response.Status)
	}

	return nil
}

// WaitApplication waits for an application to be created and returns an error
//
// Required permissions: Application.Read.All
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	id: The application to create
//	timeout: The time to wait before returning an error if not present
//	opts: The client options
func (c *HTTPClient) WaitApplication(id string, timeout int, options models.ClientOptions) (err error) {
	duration := 0
	_, err = c.GetApplication(id, options)
	for err != nil && duration < timeout {
		time.Sleep(2 * time.Second)
		duration = duration + 2
		_, err = c.GetApplication(id, options)
		c.Log.Sugar().Debugf("WaitApplication() - Duration: %d - Error: %s\n", duration, err)
	}

	c.Log.Sugar().Debugf("WaitApplication() - ID: %d \n", id)
	if duration >= timeout {
		return errors.New("timeout")
	}

	return nil
}

// UpdateApplication updates an application and returns an error
func (c *HTTPClient) UpdateApplication(app *models.Application, options models.ClientOptions) (err error) {
	return errors.New("not implemented")
}
