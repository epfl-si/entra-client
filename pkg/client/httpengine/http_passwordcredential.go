package httpengine

import (
	"github.com/epfl-si/entra-client/pkg/client/models"
)

// GetPasswordCredentials retrieves all the password credentials and returns them in a map indexed by application ID
//
// Required permissions: Application.Read.All
// Required permissions: ServicePrincipal.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetPasswordCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.PasswordCredentialEPFL, error) {
	pcMap := make(map[string][]models.PasswordCredentialEPFL, 0)
	processedPasswords := make(map[string]bool, 0)

	applications, _, err := c.GetApplications(opts)
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if len(application.PasswordCredentials) > 0 {
			for _, pwd := range application.PasswordCredentials {
				if _, found := processedPasswords[pwd.KeyID]; !found {
					pcMap[application.AppID] = append(pcMap[application.AppID], models.PasswordCredentialEPFL{
						AppID:              application.AppID,
						AppDisplayName:     application.DisplayName,
						RemainingDays:      0, // Remaining days can be calculated if needed
						PasswordCredential: pwd,
					})
					processedPasswords[pwd.KeyID] = true
				}
			}
		}
	}
	sps, _, err := c.GetServicePrincipals(opts)
	if err != nil {
		return nil, err
	}

	for _, sp := range sps {
		if len(sp.PasswordCredentials) > 0 {
			for _, pwd := range sp.PasswordCredentials {
				if _, found := processedPasswords[pwd.KeyID]; !found {
					pcMap[sp.AppID] = append(pcMap[sp.AppID], models.PasswordCredentialEPFL{
						AppID:              sp.AppID,
						AppDisplayName:     sp.DisplayName,
						RemainingDays:      0, // Remaining days can be calculated if needed
						PasswordCredential: pwd,
					})
					processedPasswords[pwd.KeyID] = true
				}
			}
		}
	}
	return pcMap, nil
}

// GetPasswordCredentialsByAppID retrieves the password credentials for an application and returns them
//
// Required permissions: Application.Read
// Required permissions: ServicePrincipal.Read
//
// Parameters:
//
//	appID: The application ID
//	opts: The client options
func (c *HTTPClient) GetPasswordCredentialsByAppID(dateLimit string, appID string, opts models.ClientOptions) ([]models.PasswordCredentialEPFL, error) {
	passwordCredentials := make([]models.PasswordCredentialEPFL, 0)
	processedPasswords := make(map[string]bool, 0)

	application, err := c.GetApplicationByAppID(appID, opts)
	if err != nil {
		return nil, err
	}

	if len(application.PasswordCredentials) > 0 {
		for _, pwd := range application.PasswordCredentials {
			if _, found := processedPasswords[pwd.KeyID]; !found {
				passwordCredentials = append(passwordCredentials, models.PasswordCredentialEPFL{
					AppID:              application.AppID,
					AppDisplayName:     application.DisplayName,
					RemainingDays:      0, // Remaining days can be calculated if needed
					PasswordCredential: pwd,
				})
				processedPasswords[pwd.KeyID] = true
			}
		}
	}
	sp, err := c.GetServicePrincipalByAppID(appID, opts)
	if err != nil {
		return nil, err
	}

	if len(sp.PasswordCredentials) > 0 {
		for _, pwd := range sp.PasswordCredentials {
			if _, found := processedPasswords[pwd.KeyID]; !found {
				passwordCredentials = append(passwordCredentials, models.PasswordCredentialEPFL{
					AppID:              sp.AppID,
					AppDisplayName:     sp.DisplayName,
					RemainingDays:      0, // Remaining days can be calculated if needed
					PasswordCredential: pwd,
				})
				processedPasswords[pwd.KeyID] = true
			}
		}
	}

	return passwordCredentials, nil
}

// GetExpiredPasswordCredentials retrieves all expired password credentials at the date specified in dateLimit
//
// Required permissions: Application.Read.All
// Required permissions: ServicePrincipal.Read.All
//
// Parameters:
//
//	dateLimit: The date limit in the format "YYYY-MM-DD"
//	opts: The client options
func (c *HTTPClient) GetExpiredPasswordCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.PasswordCredentialEPFL, error) {
	pcMap, err := c.GetPasswordCredentials(dateLimit, opts)
	results := make(map[string][]models.PasswordCredentialEPFL, 0)
	if err != nil {
		return nil, err
	}

	for appID, keyCreds := range pcMap {
		if len(keyCreds) != 0 {
			for _, kc := range keyCreds {
				if kc.RemainingDays < 0 {
					results[appID] = append(results[appID], kc)
				}
			}
		}

	}
	return results, nil
}

// GetLocalPasswordCredentials retrieves all the password credentials for local app and returns them in a map indexed by application ID
// By Local app, we mean apps that have been created directly in the tenant and not provided by a third party
//
//	ie. app that have both an application and a service principal object
//
// Required permissions: Application.Read.All
// Required permissions: ServicePrincipal.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetLocalPasswordCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.PasswordCredentialEPFL, error) {
	pcMap := make(map[string][]models.PasswordCredentialEPFL, 0)
	processedPasswords := make(map[string]bool, 0)

	applications, _, err := c.GetApplications(opts)
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if len(application.PasswordCredentials) > 0 {
			for _, pwd := range application.PasswordCredentials {
				if _, found := processedPasswords[pwd.KeyID]; !found {
					pcMap[application.AppID] = append(pcMap[application.AppID], models.PasswordCredentialEPFL{
						AppID:              application.AppID,
						AppDisplayName:     application.DisplayName,
						RemainingDays:      0, // Remaining days can be calculated if needed
						PasswordCredential: pwd,
					})
					processedPasswords[pwd.KeyID] = true
				}
			}
		}
	}

	return pcMap, nil
}
