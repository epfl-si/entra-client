package http_client

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"errors"

	"io"
)

func (c *HTTPClient) CreateApplication(app *models.Application, options models.ClientOptions) (err error) {
	return errors.New("Not implemented")
}

// DeleteApplication deletes an application and returns an error
func (c *HTTPClient) DeleteApplication(id string, opts models.ClientOptions) error {
	if id == "" {
		return errors.New("ID missing")
	}
	_, err := c.RestClient.Delete("/applications/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	return nil
}

func (c *HTTPClient) GetApplication(id string, opts models.ClientOptions) (*models.Application, error) {
	response, err := c.RestClient.Get("/applications/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var application models.Application
	err = json.Unmarshal(body, &application)
	if err != nil {
		return nil, err
	}

	return &application, nil
}

func (c *HTTPClient) GetApplications(opts models.ClientOptions) ([]*models.Application, string, error) {
	results := make([]*models.Application, 0)
	var applicationResponse models.ApplicationResponse

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

		if opts.Paging {
			break
		}
	}

	return results, applicationResponse.NextLink, nil
}

func (c *HTTPClient) UpdateApplication(app *models.Application, options models.ClientOptions) (err error) {
	return errors.New("Not implemented")
}
