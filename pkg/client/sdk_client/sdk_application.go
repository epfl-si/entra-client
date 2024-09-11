package sdkengine

import (
	"context"
	"entra-client/pkg/client/models"
	"errors"
	"fmt"

	mm "github.com/microsoftgraph/msgraph-sdk-go/models"
)

// DeleteApplication deletes an application and returns an error
func (c *SDKClient) DeleteApplication(id string, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

// CreateApplication creates an application and returns an error
// Beware: There's no unicity check on the DisplayName
// No value returned (only error)
func (c *SDKClient) CreateApplication(app *models.Application, opts models.ClientOptions) error {
	requestBody := mm.NewApplication()
	appToRequestBody(app, requestBody)

	// In case of success applications is nil (!?!)
	_, err := c.APIClient.Applications().Post(context.Background(), requestBody, nil)
	if err != nil {
		c.Log.Sugar().Debugf("CreateApplication() - 0 - Error: %s", err.Error())
		return err
	}
	return nil
}

// GetApplication returns an application by ID
func (c *SDKClient) GetApplication(id string, opts models.ClientOptions) (*models.Application, error) {
	return nil, errors.New("not implemented")
}

// GetApplications returns a list of applications
func (c *SDKClient) GetApplications(opts models.ClientOptions) ([]*models.Application, string, error) {
	fmt.Printf("GetApplications() - 0 - Token: %s\n", c.AccessToken)
	results, err := c.APIClient.Applications().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("GetApplications() - 1 - Error: %s\n", err.Error())
		return nil, "", err
	}

	return toApplications(results.GetValue()), safeString(results.GetOdataNextLink()), nil
}

// PatchApplication patches an application
func (c *SDKClient) PatchApplication(id string, app *models.Application, opts models.ClientOptions) error {
	return errors.New("not implemented")
}

func appToRequestBody(app *models.Application, requestBody *mm.Application) {
	if app.DisplayName != "" {
		requestBody.SetDisplayName(&app.DisplayName)
	}
}

func toApplications(apps []mm.Applicationable) []*models.Application {

	return []*models.Application{}
}
