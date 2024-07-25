package sdkengine

import (
	"epfl-entra/internal/models"
	"errors"
)

// CreateClaimsMappingPolicy creates an serviceprincipal and returns an error
func (c *SDKClient) CreateClaimsMappingPolicy(app *models.ClaimsMappingPolicy, opts models.ClientOptions) (string, error) {
	return "", errors.New("not implemented")
}

// DeleteClaimsMappingPolicy deletes an serviceprincipal and returns an error
func (c *SDKClient) DeleteClaimsMappingPolicy(id string, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

func (c *SDKClient) GetClaimsMappingPolicy(id string, opts models.ClientOptions) (*models.ClaimsMappingPolicy, error) {
	return nil, errors.New("not implemented")
}

func (c *SDKClient) GetClaimsMappingPolicies(opts models.ClientOptions) ([]*models.ClaimsMappingPolicy, string, error) {
	return nil, "", errors.New("not implemented")
}

// PatchClaimsMappingPolicy patches an serviceprincipal and returns an error
func (c *SDKClient) PatchClaimsMappingPolicy(id string, app *models.ClaimsMappingPolicy, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

func (c *SDKClient) UpdateClaimsMappingPolicy(app *models.ClaimsMappingPolicy, options models.ClientOptions) (err error) {
	return errors.New("not implemented")
}
