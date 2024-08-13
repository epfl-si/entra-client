package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

// CreateApplication create an application
func createApplication(name string, clientOptions models.ClientOptions) (*models.Application, *models.ServicePrincipal, error) {
	app := &models.Application{
		DisplayName: name,
	}

	newApp, err := Client.CreateApplication(app, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	err = Client.WaitApplication(newApp.ID, 60, clientOptions)
	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Application created")

	sp, err := Client.CreateServicePrincipal(&models.ServicePrincipal{
		AppID: newApp.AppID,
		Tags: []string{
			// "HideApp",
			"WindowsAzureActiveDirectoryIntegratedApp",
		},
		ServicePrincipalType: "Application"}, clientOptions)

	if err != nil {
		return nil, nil, err
	}

	fmt.Println("Service Principal created")

	return newApp, sp, nil
}

// NormalizeURI performs some modification (required by Microsoft QPI) on URI
//   - removes the trailing slash from a string
//   - replace http with api
func NormalizeURI(s string) string {
	var n string
	if len(s) > 0 && s[len(s)-1] == '/' {
		n = s[:len(s)-1]
		s = n
	}

	if len(s) > 5 && s[:5] == "http:" {
		n = "https:" + s[5:]
		s = n
	}

	return s
}

// OutputJSON returns a JSON representation of the data
func OutputJSON(data interface{}) string {
	jdata, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	return string(jdata)
}

// addCertificate adds a certificate to a Service Principal
//   - spID: the Service Principal ID
//   - certUsage: the certificate usage (e.g. 'Verify'/'Sign')
//   - certType: the certificate type	(e.g. 'AsymmetricX509Cert')
//   - certBase64: the certificate in base64 format
func addCertificate(spID string, certUsage, certType, certBase64 string, clientOptions models.ClientOptions) error {

	sp, err := Client.GetServicePrincipal(spID, clientOptions)
	if err != nil {
		return err
	}

	spKeyCredentials := []models.KeyCredential{}

	startDateTime := time.Now()
	startDateTime, _ = time.Parse(time.RFC3339, startDateTime.String())
	endDateTime := startDateTime.AddDate(0, 0, 364)
	endDateTime, _ = time.Parse(time.RFC3339, endDateTime.String())
	// Format date to this format: "2024-01-11T15:31:26Z
	// Weird bug due to Timezone, can make end date off of few hours

	// Build new KeyCredential
	// keyID := uuid.Must(uuid.NewRandom()).String()
	newCredential := models.KeyCredential{
		// KeyIdentifier: "key1",
		// EndDateTime: endDateTime,
		// KeyId:         keyID,
		// StartDateTime: startDateTime,
		DisplayName: sp.DisplayName + " " + certUsage + "ing certificate",
		Usage:       certUsage,
		Type:        certType,
		Key:         certBase64,
	}

	// if sp.KeyCredentials != nil {
	// 	spKeyCredentials = sp.KeyCredentials

	// } else {
	// 	spKeyCredentials = []models.KeyCredential{}
	// }

	spKeyCredentials = append(spKeyCredentials, newCredential)

	patchedServicePrincipal := models.ServicePrincipal{
		KeyCredentials: spKeyCredentials,
	}

	err = Client.PatchServicePrincipal(spID, &patchedServicePrincipal, clientOptions)
	if err != nil {
		return err
	}

	// Get the updated Service Principal
	sp, err = Client.GetServicePrincipal(spID, clientOptions)
	if err != nil {
		return err
	}

	// fmt.Printf("Service Principal %s updated with new certificate and keyCredential %+v", sp.DisplayName, sp.KeyCredentials)

	// Activate the certificate by its keyId
	err = Client.PatchServicePrincipal(OptID, &models.ServicePrincipal{PreferredTokenSigningKeyThumbprint: OptKeyName}, clientOptions)
	if err != nil {
		return err
	}

	return nil
}

func hideInCommand(c *cobra.Command, flags ...string) {
	c.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		for _, flag := range flags {
			fmt.Println("Hiding flag: ", flag)
			c.Flags().MarkHidden(flag)
		}
		// Call parent help func
		c.Parent().HelpFunc()(command, strings)
	})

}
