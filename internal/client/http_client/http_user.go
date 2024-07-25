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
//
// Required permissions: DeviceManagementApps.ReadWrite.All
// Required permissions: DeviceManagementConfiguration.ReadWrite.All
// Required permissions: DeviceManagementManagedDevices.ReadWrite.All
// Required permissions: DeviceManagementServiceConfig.ReadWrite.All
// Required permissions: Directory.ReadWrite.All
// Required permissions: User.ReadWrite.All
//
// Parameters:
//
//	user: The user to create
//	options: The client options
func (c *HTTPClient) CreateUser(user *models.User, opts models.ClientOptions) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Post("/users/"+buildQueryString(opts), u, h)

	if err != nil {
		return err
	}
	if response.StatusCode != 201 {
		return errors.New(response.Status)
	}

	return nil
}

// DeleteUser deletes a user by its id and returns an error
//
// Required permissions: DeviceManagementApps.ReadWrite.All
// Required permissions: DeviceManagementConfiguration.ReadWrite.All
// Required permissions: DeviceManagementManagedDevices.ReadWrite.All
// Required permissions: DeviceManagementServiceConfig.Read.All
// Required permissions: User.ReadWrite.All
//
// Parameters:
//
//	id: The ID of the user to delete
//	options: The client options
func (c *HTTPClient) DeleteUser(id string, opts models.ClientOptions) error {
	h := c.buildHeaders(opts)

	response, err := c.RestClient.Delete("/users/"+id, h)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}

// GetUser returns a user and an error
//
// Required permissions: User.ReadBasic.All
// Required permissions: DeviceManagementApps.Read.All
// Required permissions: DeviceManagementApps.ReadWrite.All
// Required permissions: DeviceManagementConfiguration.Read.All
// Required permissions: DeviceManagementConfiguration.ReadWrite.All
// Required permissions: DeviceManagementManagedDevices.Read.All
// Required permissions: DeviceManagementManagedDevices.ReadWrite.All
// Required permissions: DeviceManagementServiceConfig.Read.All
// Required permissions: DeviceManagementServiceConfig.ReadWrite.All
// Required permissions: Directory.Read.All
// Required permissions: Directory.ReadWrite.All
// Required permissions: User.Read
// Required permissions: User.Read.All
//
// Parameters:
//
//	user: The user to create
//	options: The client options
func (c *HTTPClient) GetUser(id string, opts models.ClientOptions) (*models.User, error) {
	h := c.buildHeaders(opts)

	response, err := c.RestClient.Get("/users/"+id+buildQueryString(opts), h)
	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(io.Reader(response.Body))
	response.Body.Close()
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, errors.New(response.Status)
	}

	var user models.User
	err = json.Unmarshal(body, &user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

// GetUsers returns a list of users, a pagination link and an error
//
// Required permissions: User.ReadBasic.All
// Required permissions: DeviceManagementApps.Read.All
// Required permissions: DeviceManagementApps.ReadWrite.All
// Required permissions: DeviceManagementConfiguration.Read.All
// Required permissions: DeviceManagementConfiguration.ReadWrite.All
// Required permissions: DeviceManagementManagedDevices.Read.All
// Required permissions: DeviceManagementManagedDevices.ReadWrite.All
// Required permissions: DeviceManagementServiceConfig.Read.All
// Required permissions: DeviceManagementServiceConfig.ReadWrite.All
// Required permissions: Directory.Read.All
// Required permissions: Directory.ReadWrite.All
// Required permissions: User.Read.All
// Required permissions: User.ReadWrite.All
//
// Parameters:
//
//	options: The client options
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
		if response.StatusCode != 200 {
			return nil, "", errors.New(response.Status)
		}

		if opts.Paging {
			break
		}
	}

	return results, "", nil
}

// UpdateUser updates a user and returns an error
//
// Required permissions: ?????
//
// Parameters:
//
//	user: The modified user
//	options: The client options
func (c *HTTPClient) UpdateUser(user *models.User, opts models.ClientOptions) error {
	u, err := json.Marshal(user)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)

	response, err := c.RestClient.Put("/users/"+user.ID, u, h)
	if err != nil {
		return err
	}
	if response.StatusCode != 204 {
		return errors.New(response.Status)
	}

	return nil
}
