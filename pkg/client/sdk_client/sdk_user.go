package sdkengine

import (
	"entra-client/pkg/client/models"

	"errors"
)

// CreateUser creates a user
func (c *SDKClient) CreateUser(user *models.User, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

// DeleteUser deletes a user
func (c *SDKClient) DeleteUser(id string, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

func (c *SDKClient) GetUser(id string, opts models.ClientOptions) (*models.User, error) {
	return nil, errors.New("not implemented")
}

func (c *SDKClient) GetUsers(opts models.ClientOptions) ([]*models.User, string, error) {
	return nil, "", errors.New("not implemented")
}

// UpdateUser updates a user
func (c *SDKClient) UpdateUser(user *models.User, opts models.ClientOptions) error {
	return errors.New("not implemented")
}
