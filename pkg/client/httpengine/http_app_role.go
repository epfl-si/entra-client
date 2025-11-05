package httpengine

import (
	"encoding/json"

	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/google/uuid"
)

// CreateAppRole creates an application role and returns an error
//
// Required permissions: Application.ReadWrite.All
//
// Parameters:
//
//	appID: The application ID where the role will be created
//	appRole: The application role to be created
//	opts: The client options
func (c *HTTPClient) CreateAppRoleByAppID(appID string, appRole *models.AppRole, opts models.ClientOptions) (string, error) {
	if opts.Default {
		appRole = &models.AppRole{
			ID:          uuid.NewString(),
			DisplayName: "DefaultAccess",
			Description: "Default App Role",
			Value:       "default_access",
			AllowedMemberTypes: []string{
				"User",
				"Application",
			},
			IsEnabled: true,
			Origin:    "Application",
		}
	}

	// If no ID is provided, generate a new UUID
	if appRole.ID == "" {
		appRole.ID = uuid.NewString()
	}

	// Get the application by its appID using c.RestClient
	app, err := c.GetApplicationByAppID(appID, opts)
	if err != nil {
		return "", err
	}

	// Append the new appRole to the existing AppRoles
	app.AppRoles = append(app.AppRoles, *appRole)

	// Marshal the updated application back to JSON
	u, err := json.Marshal(map[string]interface{}{
		"appRoles": app.AppRoles,
	})
	if err != nil {
		return "", err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	resp, err := c.RestClient.Patch("/applications/"+app.ID, u, h)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return appRole.ID, nil
}
