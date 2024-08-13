package cmd

import (
	"epfl-entra/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

var OptKeyName string

// applicationSAMLCertificateAddCmd represents the applicationSAMLCertificateAdd command
var applicationSAMLCertificateAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a certificqte to a SAML application",
	Long: `Add a certificate to a SAML application

Example:

  ./ecli application saml certificate add --id 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --keyname key1
`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLCertificateAdd called")

		if OptID == "" {
			panic("Service Principal ID is required (use --id)")
		}
		if OptKeyName == "" {
			panic("Service key name is required (use --keyname)")
		}

		sp, err := Client.GetServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		spKeyCredentials := []models.KeyCredential{}

		startDateTime := time.Now()
		endDateTime := startDateTime.AddDate(1, 0, 0)

		// Build new KeyCredential
		newCredential := models.KeyCredential{
			KeyIdentifier: OptKeyName,
			EndDateTime:   &endDateTime,
			KeyId:         uuid.Must(uuid.NewRandom()).String(),
			StartDateTime: &startDateTime,
			DisplayName:   OptKeyName,
			Usage:         "Verify",
			Type:          "AsymmetricX509Cert",
			Key:           "MIIC8DCCAdigAwIBAgIQYuAOdippIbJMaUAsghvhUDANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQDEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yNDA2MDYxNDAwNDRaFw0yNzA2MDYxNDAwNDNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQgU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyoHs70fxBrDS1aKCQITj3CzJLJQPnDhv3GX35k/gUXYSrG4JJjREaxqoHWx5tu4CVVlsCckD4OYMUIHm0dNgy2Dkrq84BR7N7BTwnuqUg5hNiofE0zWH+6ovFtjJ9Bci5ai7LglgmH1JXh8kJVR1bptrZskXR47Yae8xnKrZ3rh/1ym/nwsTwz//PKBThIhIyJ6vQlq1Wnn2ZkH9AqpLURiAagbfX+hZpcJAgFLnay9OCrFg/d5ykefzlROKeIBf71ySPViciVzDWmJGgJffYaKMsQAjTOdJi+J4WZ9cwaSMda/G5aZ3/2+DjcZM5P2Tmtr3HH/xLaOqcV7NnxOoTQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQABS37CTtFRhLFcXhcfuNGeTILTx8vUogBdPxlccC9e7+TyyKEKDKuyK3C0MeTh7Dl/6oKACcOQC1Ygcwzc0KzsBPrJBubbhmNpzwZsIhYO1mvxgkcDKmaYrP5G/5dvonM2+V0xc8XzZe5Hm/dwlPNdmCRKIRR9o350ariB4p56fWcJYIK0MnQXHpMsZqlos5a7a3Zmm3HPfanZnHaw1rVHkOrYKkT3/eQf7oCalQ7D8iFBs3K5n2MtSf1QSsUbQcAN1BF1NIfZA+hqAlARf9kc1dmE/6o+dXncDizLgGWuFoGEP6FOao53gaC8OefSjiOd1idB3C2iQv7NnaP9GDiC",
		}

		if sp.KeyCredentials != nil {
			spKeyCredentials = sp.KeyCredentials

		} else {
			spKeyCredentials = []models.KeyCredential{}
		}

		spKeyCredentials = append(spKeyCredentials, newCredential)

		patchedServicePrincipal := models.ServicePrincipal{
			KeyCredentials: spKeyCredentials,
		}

		err = Client.PatchServicePrincipal(OptID, &patchedServicePrincipal, clientOptions)
		if err != nil {
			panic(err)
		}

		// Get the updated Service Principal
		sp, err = Client.GetServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		cmd.Printf("Service Principal %s updated with new certificate and keyCredential %+v", sp.DisplayName, sp.KeyCredentials)

		// Activate the certificate by its keyId
		err = Client.PatchServicePrincipal(OptID, &models.ServicePrincipal{PreferredTokenSigningKeyThumbprint: OptKeyName}, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationSAMLCertificateCmd.AddCommand(applicationSAMLCertificateAddCmd)
	// applicationSAMLCertificateAddCmd.Flags().StringVar(&OptKeyName, "keyname", "", "Name of the added certificate")
	applicationSAMLCertificateAddCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLCertificateAddCmd.Flags().MarkHidden("top")
		applicationSAMLCertificateAddCmd.Flags().MarkHidden("skip")
		applicationSAMLCertificateAddCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLCertificateAddCmd.Flags().MarkHidden("select")
		applicationSAMLCertificateAddCmd.Flags().MarkHidden("search")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
