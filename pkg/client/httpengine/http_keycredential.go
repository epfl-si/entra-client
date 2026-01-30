package httpengine

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/epfl-si/entra-client/pkg/client/models"
)

// GetKeyCredentials retrieves all the key credentials and returns them in a map indexed by application ID
//
// Required permissions: Application.Read.All
// Required permissions: ServicePrincipal.Read.All
//
// Parameters:
//
//	opts: The client options
func (c *HTTPClient) GetKeyCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.KeyCredentialEPFL, error) {
	kcMap := make(map[string][]models.KeyCredentialEPFL, 0)
	processedCerts := make(map[string]bool, 0)
	var expirationDate time.Time

	if dateLimit != "" {
		var err error
		expirationDate, err = time.Parse("2006-01-02", dateLimit)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %v", err)
		}
	} else {
		expirationDate = time.Now()
	}

	applications, _, err := c.GetApplications(opts)
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if len(application.KeyCredentials) > 0 {
			for _, cert := range application.KeyCredentials {
				if _, found := processedCerts[cert.KeyID]; !found {
					nbDaysRemaining, err := nbDaysBetweenDates(expirationDate.Format("2006-01-02"), time.Time(*cert.EndDateTime).Format("2006-01-02"))
					if err != nil {
						return nil, err
					}
					kcMap[application.AppID] = append(kcMap[application.AppID], models.KeyCredentialEPFL{
						AppID:          application.AppID,
						AppDisplayName: application.DisplayName,
						RemainingDays:  nbDaysRemaining,
						KeyCredential:  cert,
					})
					processedCerts[cert.KeyID] = true
				}
			}
		}
	}
	sps, _, err := c.GetServicePrincipals(opts)
	if err != nil {
		return nil, err
	}

	for _, sp := range sps {
		if len(sp.KeyCredentials) > 0 {
			for _, cert := range sp.KeyCredentials {
				nbDaysRemaining, err := nbDaysBetweenDates(expirationDate.Format("2006-01-02"), time.Time(*cert.EndDateTime).Format("2006-01-02"))
				if err != nil {
					return nil, err
				}
				if _, found := processedCerts[cert.KeyID]; !found {
					kcMap[sp.AppID] = append(kcMap[sp.AppID], models.KeyCredentialEPFL{
						AppID:          sp.AppID,
						AppDisplayName: sp.DisplayName,
						RemainingDays:  nbDaysRemaining,
						KeyCredential:  cert,
					})
					processedCerts[cert.KeyID] = true
				}
			}
		}
	}
	return kcMap, nil
}

// GetKeyCredentials retrieves the key credentials for an application and returns them
//
// Required permissions: Application.Read
// Required permissions: ServicePrincipal.Read
//
// Parameters:
//
//	appID: The application ID
//	opts: The client options
func (c *HTTPClient) GetKeyCredentialsByAppID(dateLimit, appID string, opts models.ClientOptions) ([]models.KeyCredentialEPFL, error) {
	keyCredentials := make([]models.KeyCredentialEPFL, 0)
	processedCerts := make(map[string]bool, 0)

	application, err := c.GetApplicationByAppID(appID, opts)
	if err != nil {
		return nil, err
	}

	if len(application.KeyCredentials) > 0 {
		for _, cert := range application.KeyCredentials {
			nbDaysRemaining, err := nbDaysBetweenDates(time.Now().Format("2006-01-02"), time.Time(*cert.EndDateTime).Format("2006-01-02"))
			if err != nil {
				return nil, err
			}
			if _, found := processedCerts[cert.KeyID]; !found {
				keyCredentials = append(keyCredentials, models.KeyCredentialEPFL{
					AppID:          application.AppID,
					AppDisplayName: application.DisplayName,
					RemainingDays:  nbDaysRemaining,
					KeyCredential:  cert,
				})
				processedCerts[cert.KeyID] = true
			}
		}
	}
	sp, err := c.GetServicePrincipalByAppID(appID, opts)
	if err != nil {
		return nil, err
	}

	if len(sp.KeyCredentials) > 0 {
		for _, cert := range sp.KeyCredentials {
			nbDaysRemaining, err := nbDaysBetweenDates(time.Now().Format("2006-01-02"), time.Time(*cert.EndDateTime).Format("2006-01-02"))
			if err != nil {
				return nil, err
			}
			if _, found := processedCerts[cert.KeyID]; !found {
				keyCredentials = append(keyCredentials, models.KeyCredentialEPFL{
					AppID:          sp.AppID,
					AppDisplayName: sp.DisplayName,
					RemainingDays:  nbDaysRemaining,
					KeyCredential:  cert,
				})
				processedCerts[cert.KeyID] = true
			}
		}
	}

	return keyCredentials, nil
}

// GetExpiredKeyCredentials retrieves all expired key credentials at the date specified in dateLimit
//
// Required permissions: Application.Read.All
// Required permissions: ServicePrincipal.Read.All
//
// Parameters:
//
//	dateLimit: The date limit in the format "YYYY-MM-DD"
//	opts: The client options
func (c *HTTPClient) GetExpiredKeyCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.KeyCredentialEPFL, error) {
	kcMap, err := c.GetKeyCredentials(dateLimit, opts)
	results := make(map[string][]models.KeyCredentialEPFL, 0)
	if err != nil {
		return nil, err
	}

	for appID, keyCreds := range kcMap {
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

// nbDaysBetweenDates returns date2-date1 in days
func nbDaysBetweenDates(date1, date2 string) (int64, error) {
	startDate, err := time.Parse("2006-01-02", date1)
	if err != nil {
		return 0, err
	}
	endDate, err := time.Parse("2006-01-02", date2)
	if err != nil {
		return 0, err
	}

	return int64(endDate.Sub(startDate).Hours() / 24), nil
}

// GetLocalKeyCredentials retrieves all the key credentials for local app and returns them in a map indexed by application ID
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
func (c *HTTPClient) GetLocalKeyCredentials(dateLimit string, opts models.ClientOptions) (map[string][]models.KeyCredentialEPFL, error) {
	kcMap := make(map[string][]models.KeyCredentialEPFL, 0)
	processedCerts := make(map[string]bool, 0)
	var expirationDate time.Time

	if dateLimit != "" {
		var err error
		expirationDate, err = time.Parse("2006-01-02", dateLimit)
		if err != nil {
			return nil, fmt.Errorf("invalid date format: %v", err)
		}
	} else {
		expirationDate = time.Now()
	}

	applications, _, err := c.GetApplications(opts)
	if err != nil {
		return nil, err
	}

	for _, application := range applications {
		if len(application.KeyCredentials) > 0 {
			for _, cert := range application.KeyCredentials {
				if _, found := processedCerts[cert.KeyID]; !found {
					nbDaysRemaining, err := nbDaysBetweenDates(expirationDate.Format("2006-01-02"), time.Time(*cert.EndDateTime).Format("2006-01-02"))
					if err != nil {
						return nil, err
					}
					kcMap[application.AppID] = append(kcMap[application.AppID], models.KeyCredentialEPFL{
						AppID:          application.AppID,
						AppDisplayName: application.DisplayName,
						RemainingDays:  nbDaysRemaining,
						KeyCredential:  cert,
					})
					processedCerts[cert.KeyID] = true
				}
			}
		}
	}

	return kcMap, nil
}

// AddKeyCredentialToApplication adds a key credential to an application or returns an error
//
// Required permissions: Application.ReadWrite
//
// Parameters:
//
//	appID: The application ID
//	displayName: The certificate description
//	startDateTime: The certificate start date
//	endDateTime: The certificate end date
//	key: The base64 certificate
//	opts: The client options
func (c *HTTPClient) AddKeyCredentialToApplication(appID, displayName, startDateTime, endDateTime, key string, opts models.ClientOptions) (err error) {
	if appID == "" {
		return fmt.Errorf("appID missing")
	}
	if displayName == "" {
		return fmt.Errorf("displayName missing")
	}
	if startDateTime == "" {
		return fmt.Errorf("startDateTime missing")
	}
	if endDateTime == "" {
		return fmt.Errorf("endDateTime missing")
	}
	if key == "" {
		return fmt.Errorf("key missing")
	}

	application, err := c.GetApplicationByAppID(appID, opts)
	if err != nil {
		return err
	}

	dStartDateTime, err := time.Parse("2006-01-02T15:04:05Z", startDateTime)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	dEndDateTime, err := time.Parse("2006-01-02T15:04:05Z", endDateTime)
	if err != nil {
		return fmt.Errorf("invalid date format: %v", err)
	}
	cStartDateTime := models.CustomTime(dStartDateTime)
	cEndDateTime := models.CustomTime(dEndDateTime)

	keyCredentials := append(application.KeyCredentials, models.KeyCredential{
		StartDateTime: &cStartDateTime,
		EndDateTime:   &cEndDateTime,
		Type:          "AsymmetricX509Cert",
		Usage:         "Verify",
		Key:           key,
		DisplayName:   displayName,
	})

	applicationKeyCredentials := &models.Application{
		KeyCredentials: keyCredentials,
	}

	u, err := json.Marshal(applicationKeyCredentials)
	if err != nil {
		return err
	}

	h := c.buildHeaders(opts)
	h["Content-Type"] = "application/json"

	response, err := c.RestClient.Patch("/applications/"+application.ID, u, h)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 204 {
		return fmt.Errorf("Error while sending certificate to Entra: %v", response.Status)
	}

	return nil
}
