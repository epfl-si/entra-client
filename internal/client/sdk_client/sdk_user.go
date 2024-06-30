package sdk_client

import (
	"epfl-entra/internal/models"
	"strconv"

	"errors"

	graphusers "github.com/microsoftgraph/msgraph-sdk-go/users"
)

// CreateUser creates a user
func (c *SDKClient) CreateUser(user *models.User, opts models.ClientOptions) error {
	return errors.New("Not implemented")
}

// DeleteUser deletes a user
func (c *SDKClient) DeleteUser(id string, opts models.ClientOptions) error {
	return errors.New("Not implemented")
}

func (c *SDKClient) GetUser(id string, opts models.ClientOptions) (*models.User, error) {
	return nil, errors.New("Not implemented")
}

func (c *SDKClient) GetUsers(opts models.ClientOptions) ([]*models.User, string, error) {
	return nil, "", errors.New("Not implemented")
}

// UpdateUser updates a user
func (c *SDKClient) UpdateUser(user *models.User, opts models.ClientOptions) error {
	return errors.New("Not implemented")
}

func getUserRequestConfiguration(opts models.ClientOptions) (*graphusers.UsersRequestBuilderGetRequestConfiguration, error) {
	requestParameters := &graphusers.UsersRequestBuilderGetQueryParameters{}

	if opts.Top != "" {
		parsed, err := (strconv.ParseInt(opts.Top, 10, 32))
		if err != nil {
			return nil, errors.New("Top parameter error: " + err.Error())
		}
		requestTop := int32(parsed)

		requestParameters.Top = &requestTop
	}

	configuration := &graphusers.UsersRequestBuilderGetRequestConfiguration{
		QueryParameters: requestParameters,
	}

	return configuration, nil
}
