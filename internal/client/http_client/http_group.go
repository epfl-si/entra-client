package http_client

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"errors"
	"net/http"

	"io"
)

// CreateGroup creates a group and returns an error
func (c *HTTPClient) CreateGroup(group *models.Group, opts models.ClientOptions) error {
	u, err := json.Marshal(group)
	if err != nil {
		return err
	}

	resp, err := c.RestClient.Post("/groups/"+buildQueryString(opts), u, rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// DeleteGroup deletes a group and returns an error
func (c *HTTPClient) DeleteGroup(id string, opts models.ClientOptions) error {
	_, err := c.RestClient.Delete("/groups/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	return nil
}

// GetGroup returns a group and an error
func (c *HTTPClient) GetGroup(id string, opts models.ClientOptions) (*models.Group, error) {
	response, err := c.RestClient.Get("/groups/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var group models.Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (c *HTTPClient) GetGroups(opts models.ClientOptions) ([]*models.Group, string, error) {
	results := make([]*models.Group, 0)

	var response *http.Response
	var err error

	h := c.buildHeaders(opts)

	response, err = c.RestClient.Get("/groups"+buildQueryString(opts), h)

	for {
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var groupResponse models.GroupResponse
		err = json.Unmarshal(body, &groupResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetGroups() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, groupResponse.Value...)

		if groupResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetGroups() - 4 - Calling Next: %s\n", groupResponse.NextLink)
		response, err = c.RestClient.Get(groupResponse.NextLink, h)

		if opts.Paging {
			break
		}
	}

	return results, "", nil
}

// UpdateGroup updates a group and returns an error
func (c *HTTPClient) UpdateGroup(group *models.Group, opts models.ClientOptions) error {
	u, err := json.Marshal(group)
	if err != nil {
		return err
	}

	_, err = c.RestClient.Put("/groups/"+group.ID, u, rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	return nil
}
