package httpengine

import (
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
func (c *HTTPClient) GetKeyCredentialsByAppID(appID, dateLimit string, opts models.ClientOptions) ([]models.KeyCredentialEPFL, error) {
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
