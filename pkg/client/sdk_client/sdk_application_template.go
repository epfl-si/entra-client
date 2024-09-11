package sdkengine

import (
	"context"
	"entra-client/internal/models"
	"errors"
	"fmt"

	mm "github.com/microsoftgraph/msgraph-sdk-go/models"
)

func (c *SDKClient) InstantiateApplicationTemplate(id, name string, opts models.ClientOptions) (*models.Application, *models.ServicePrincipal, error) {
	return nil, nil, errors.New("not implemented")
}

func (c *SDKClient) GetApplicationTemplate(id string, opts models.ClientOptions) (*models.ApplicationTemplate, error) {
	return nil, errors.New("not implemented")
}

func (c *SDKClient) GetApplicationTemplates(opts models.ClientOptions) ([]*models.ApplicationTemplate, string, error) {
	fmt.Printf("GetApplicationTemplates() - 0 - Token: %s\n", c.AccessToken)
	results, err := c.APIClient.ApplicationTemplates().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("GetApplicationTemplates() - 1 - Error: %s\n", err.Error())
		return nil, "", err
	}

	return toApplicationTemplates(results.GetValue()), safeString(results.GetOdataNextLink()), nil
}

func toApplicationTemplates(apps []mm.ApplicationTemplateable) []*models.ApplicationTemplate {

	return []*models.ApplicationTemplate{}
}
