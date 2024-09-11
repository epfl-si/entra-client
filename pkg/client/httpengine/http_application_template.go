package httpengine

import (
	"encoding/json"
	"entra-client/pkg/client/models"
	"errors"

	"io"
)

type applicationTemplatePost struct {
	ID          string `json:"id,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

type instantiateResponse struct {
	Context          string                   `json:"@odata.context"`
	Application      *models.Application      `json:"application,omitempty"`
	ServicePrincipal *models.ServicePrincipal `json:"servicePrincipal,omitempty"`
}

// InstantiateApplicationTemplate instantiates an application template and returns a service principal and an error
//
// Required permissions: Application.ReadWrite
// Required permissions: Directory.ReadWrite.All
//
// Parameters:
//
//	id: The application template ID
//	name: The name of the application to create
//	opts: The client options
func (c *HTTPClient) InstantiateApplicationTemplate(id, name string, opts models.ClientOptions) (*models.Application, *models.ServicePrincipal, error) {
	if id == "" {
		return nil, nil, errors.New("ID missing")
	}
	if name == "" {
		return nil, nil, errors.New("Name missing")
	}
	u, err := json.Marshal(applicationTemplatePost{ID: id, DisplayName: name})
	if err != nil {
		return nil, nil, err
	}

	c.Log.Sugar().Debug("InstantiateApplicationTemplate() - 1 - U: " + string(u))
	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Post("/applicationTemplates/"+id+"/instantiate", u, h)
	if err != nil {
		return nil, nil, err
	}
	if response.StatusCode != 201 {
		return nil, nil, errors.New(response.Status)
	}

	body, err := io.ReadAll(response.Body)
	if err != nil {
		c.Log.Sugar().Debugf("InstantiateApplicationTemplate() - Body read error: %s\n", err.Error())
		return nil, nil, err
	}

	defer response.Body.Close()
	var ir *instantiateResponse
	err = json.Unmarshal(body, &ir)
	c.Log.Sugar().Debugf("InstantiateApplicationTemplate() - Body: %s\n", string(body))
	if err != nil {
		c.Log.Sugar().Debugf("InstantiateApplicationTemplate() - Response unmarshall error: %s\n", err.Error())
		return nil, nil, err
	}

	return ir.Application, ir.ServicePrincipal, nil
}

// GetApplicationTemplate gets an application template by its Id and returns the application template and an error
//
// Required permissions: ???
//
// Parameters:
//
//	id: The application template ID
//	opts: The client options
func (c *HTTPClient) GetApplicationTemplate(id string, opts models.ClientOptions) (*models.ApplicationTemplate, error) {
	h := c.buildHeaders(opts)
	// No options for ApplicationTemplate as it's not collection objects
	response, err := c.RestClient.Get("/applicationTemplates/"+id, h)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	response.Body.Close()

	var application models.ApplicationTemplate
	err = json.Unmarshal(body, &application)
	if err != nil {
		return nil, err
	}

	return &application, nil
}

// GetApplicationTemplates gets all application templates and returns a slice of application templates, a pagination link and an error
//
// Required permissions: ???
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetApplicationTemplates(opts models.ClientOptions) ([]*models.ApplicationTemplate, string, error) {
	results := make([]*models.ApplicationTemplate, 0)
	var applicationResponse models.ApplicationTemplateResponse
	var err error

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get("/applicationTemplates"+buildQueryString(opts), h)

	for {

		if err != nil {
			c.Log.Sugar().Debugf("GetApplicationTemplates() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetApplicationTemplates() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		err = json.Unmarshal(body, &applicationResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetApplicationTemplates() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var applicationResponse models.ApplicationTemplateResponse
		err = json.Unmarshal(body, &applicationResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetApplicationTemplates() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, applicationResponse.Value...)

		if applicationResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", applicationResponse.NextLink)
		response, err = c.RestClient.Get(applicationResponse.NextLink, h)

		if opts.Paging {
			break
		}
	}

	return results, applicationResponse.NextLink, nil
}
