package sdkengine

import (
	"context"
	"errors"
	"fmt"

	"github.com/epfl-si/entra-client/pkg/client/models"

	mm "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func spToRequestBody(sp *models.ServicePrincipal, requestBody *mm.ServicePrincipal) {
	if sp.DisplayName != "" {
		requestBody.SetDisplayName(&sp.DisplayName)
	}
}

// AssociateAppRoleToServicePrincipal associates a serviceprincipal to a group and returns an error
func (c *SDKClient) AssociateAppRoleToServicePrincipal(assignment *models.AppRoleAssignment, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

// AssociateClaimsPolicyToServicePrincipal associates a Claims Policy to a serviceprincipal and returns an error
func (c *SDKClient) AssociateClaimsPolicyToServicePrincipal(claimsPolicyID, servicePrincipalID string) error {
	return errors.New("not implemented")
}

// DeleteServicePrincipal deletes an application and returns an error
func (c *SDKClient) DeleteServicePrincipal(id string, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

// CreateServicePrincipal creates an application and returns an error
// Beware: There's no unicity check on the DisplayName
// No value returned (only error)
func (c *SDKClient) CreateServicePrincipal(sp *models.ServicePrincipal, opts models.ClientOptions) error {
	requestBody := mm.NewServicePrincipal()
	spToRequestBody(sp, requestBody)

	// In case of success applications is nil (!?!)
	_, err := c.APIClient.ServicePrincipals().Post(context.Background(), requestBody, nil)
	if err != nil {
		c.Log.Sugar().Debugf("CreateServicePrincipal() - 0 - Error: %s", err.Error())
		return err
	}
	return nil
}

func (c *SDKClient) GetServicePrincipal(id string, opts models.ClientOptions) (*models.ServicePrincipal, error) {
	return nil, errors.New("not implemented")
}
func (c *SDKClient) GetServicePrincipals(opts models.ClientOptions) ([]*models.ServicePrincipal, string, error) {
	fmt.Printf("GetServicePrincipals() - 0 - Token: %s\n", c.AccessToken)
	results, err := c.APIClient.ServicePrincipals().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("GetServicePrincipals() - 1 - Error: %s\n", err.Error())
		return nil, "", err
	}

	return toServicePrincipals(results.GetValue()), safeString(results.GetOdataNextLink()), nil
}

// PatchServicePrincipal patches a Service principal
func (c *SDKClient) PatchServicePrincipal(id string, sp *models.ServicePrincipal, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

func toServicePrincipals(sps []mm.ServicePrincipalable) []*models.ServicePrincipal {

	return []*models.ServicePrincipal{}
}
