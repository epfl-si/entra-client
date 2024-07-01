package httpengine

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"epfl-entra/pkg/rest"
	"errors"
	"io"
	"net/http"
)

// CreateUser creates a user and returns an error
func (c *HTTPClient) CreateUser(user *models.User, opts models.ClientOptions) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	resp, err := c.RestClient.Post("/users/"+buildQueryString(opts), u, rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
		return errors.New(resp.Status)
	}

	return nil
}

// DeleteUser deletes a user and returns an error
func (c *HTTPClient) DeleteUser(id string, opts models.ClientOptions) error {
	_, err := c.RestClient.Delete("/users/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	return nil
}

// GetUser returns a user and an error
func (c *HTTPClient) GetUser(id string, opts models.ClientOptions) (*models.User, error) {
	response, err := c.RestClient.Get("/users/"+id+buildQueryString(opts), rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	if err != nil {
		return nil, err
	}

	response.Body.Close()

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUsers returns a list of users, a pagination link and an error
func (c *HTTPClient) GetUsers(opts models.ClientOptions) ([]*models.User, string, error) {
	results := make([]*models.User, 0)

	var response *http.Response
	var err error

	h := c.buildHeaders(opts)

	c.Log.Sugar().Debugf("GetUsers() - 0 - Url: %s\n", "/users"+buildQueryString(opts))
	response, err = c.RestClient.Get("/users"+buildQueryString(opts), h)

	c.Log.Sugar().Debugf("GetUsers() - response: %+v\n", response)
	for {
		if err != nil {
			c.Log.Sugar().Debugf("GetUsers() - 1 - Error: %s\n", err.Error())
			return nil, "", err
		}

		body, err := io.ReadAll(io.Reader(response.Body))
		if err != nil {
			c.Log.Sugar().Debugf("GetUsers() - 2 - Error: %s\n", err.Error())
			return nil, "", err
		}

		response.Body.Close()

		var groupResponse models.UserResponse
		err = json.Unmarshal(body, &groupResponse)
		if err != nil {
			c.Log.Sugar().Debugf("GetUsers() - 3 - Error: %s\n", err.Error())
			return nil, "", err
		}

		results = append(results, groupResponse.Value...)
		c.Log.Sugar().Debugf("GetUsers() - 3.5 - groupResponse: %+v\n", groupResponse)

		if groupResponse.NextLink == "" {
			break
		}

		c.Log.Sugar().Debugf("GetUsers() - 4 - Calling Next: %s\n", groupResponse.NextLink)
		response, err = c.RestClient.Get(groupResponse.NextLink, rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
		c.Log.Sugar().Debugf("GetUsers() - 5 - Next Response: %+v\n", response)
		c.Log.Sugar().Debugf("GetUsers() - 6 - Paging: %#v\n", opts.Paging)

		if opts.Paging {
			break
		}
	}

	return results, "", nil
}

// UpdateUser updates a user and returns an error
func (c *HTTPClient) UpdateUser(user *models.User, opts models.ClientOptions) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	_, err = c.RestClient.Put("/users/"+user.ID, u, rest.Headers{"Authorization": rest.TokenBearerString(c.AccessToken)})
	if err != nil {
		return err
	}

	return nil
}
