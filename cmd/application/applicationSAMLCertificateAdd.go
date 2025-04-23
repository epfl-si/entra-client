package cmdapplication

import (
	"time"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// OptKeyName is associated with the --keyname flag
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

		if rootcmd.OptID == "" {
			cmd.PrintErr("Service Principal ID is required (use --id)\n")
			return
		}
		if OptKeyName == "" {
			cmd.PrintErr("Service key name is required (use --keyname)\n")
			return
		}

		sp, err := rootcmd.Client.GetServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		spKeyCredentials := []models.KeyCredential{}

		startDateTime := time.Now().UTC()
		endDateTime := startDateTime.AddDate(0, 6, 0).UTC()
		cStartDateTime := models.CustomTime(startDateTime)
		cEndDateTime := models.CustomTime(endDateTime)

		// Build new KeyCredential
		newCredential := models.KeyCredential{
			CustomKeyIdentifier: OptKeyName,
			EndDateTime:         &cEndDateTime,
			KeyID:               uuid.Must(uuid.NewRandom()).String(),
			StartDateTime:       &cStartDateTime,
			DisplayName:         OptKeyName,
			Usage:               "Verify",
			Type:                "AsymmetricX509Cert",
			// Key:                 []byte("base64MIIC8DCCAdigAwIBAgIQYuAOdippIbJMaUAsghvhUDANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQDEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yNDA2MDYxNDAwNDRaFw0yNzA2MDYxNDAwNDNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQgU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyoHs70fxBrDS1aKCQITj3CzJLJQPnDhv3GX35k/gUXYSrG4JJjREaxqoHWx5tu4CVVlsCckD4OYMUIHm0dNgy2Dkrq84BR7N7BTwnuqUg5hNiofE0zWH+6ovFtjJ9Bci5ai7LglgmH1JXh8kJVR1bptrZskXR47Yae8xnKrZ3rh/1ym/nwsTwz//PKBThIhIyJ6vQlq1Wnn2ZkH9AqpLURiAagbfX+hZpcJAgFLnay9OCrFg/d5ykefzlROKeIBf71ySPViciVzDWmJGgJffYaKMsQAjTOdJi+J4WZ9cwaSMda/G5aZ3/2+DjcZM5P2Tmtr3HH/xLaOqcV7NnxOoTQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQABS37CTtFRhLFcXhcfuNGeTILTx8vUogBdPxlccC9e7+TyyKEKDKuyK3C0MeTh7Dl/6oKACcOQC1Ygcwzc0KzsBPrJBubbhmNpzwZsIhYO1mvxgkcDKmaYrP5G/5dvonM2+V0xc8XzZe5Hm/dwlPNdmCRKIRR9o350ariB4p56fWcJYIK0MnQXHpMsZqlos5a7a3Zmm3HPfanZnHaw1rVHkOrYKkT3/eQf7oCalQ7D8iFBs3K5n2MtSf1QSsUbQcAN1BF1NIfZA+hqAlARf9kc1dmE/6o+dXncDizLgGWuFoGEP6FOao53gaC8OefSjiOd1idB3C2iQv7NnaP9GDiC"),
			Key: "base64MIIC8DCCAdigAwIBAgIQYuAOdippIbJMaUAsghvhUDANBgkqhkiG9w0BAQsFADA0MTIwMAYDVQQDEylNaWNyb3NvZnQgQXp1cmUgRmVkZXJhdGVkIFNTTyBDZXJ0aWZpY2F0ZTAeFw0yNDA2MDYxNDAwNDRaFw0yNzA2MDYxNDAwNDNaMDQxMjAwBgNVBAMTKU1pY3Jvc29mdCBBenVyZSBGZWRlcmF0ZWQgU1NPIENlcnRpZmljYXRlMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyoHs70fxBrDS1aKCQITj3CzJLJQPnDhv3GX35k/gUXYSrG4JJjREaxqoHWx5tu4CVVlsCckD4OYMUIHm0dNgy2Dkrq84BR7N7BTwnuqUg5hNiofE0zWH+6ovFtjJ9Bci5ai7LglgmH1JXh8kJVR1bptrZskXR47Yae8xnKrZ3rh/1ym/nwsTwz//PKBThIhIyJ6vQlq1Wnn2ZkH9AqpLURiAagbfX+hZpcJAgFLnay9OCrFg/d5ykefzlROKeIBf71ySPViciVzDWmJGgJffYaKMsQAjTOdJi+J4WZ9cwaSMda/G5aZ3/2+DjcZM5P2Tmtr3HH/xLaOqcV7NnxOoTQIDAQABMA0GCSqGSIb3DQEBCwUAA4IBAQABS37CTtFRhLFcXhcfuNGeTILTx8vUogBdPxlccC9e7+TyyKEKDKuyK3C0MeTh7Dl/6oKACcOQC1Ygcwzc0KzsBPrJBubbhmNpzwZsIhYO1mvxgkcDKmaYrP5G/5dvonM2+V0xc8XzZe5Hm/dwlPNdmCRKIRR9o350ariB4p56fWcJYIK0MnQXHpMsZqlos5a7a3Zmm3HPfanZnHaw1rVHkOrYKkT3/eQf7oCalQ7D8iFBs3K5n2MtSf1QSsUbQcAN1BF1NIfZA+hqAlARf9kc1dmE/6o+dXncDizLgGWuFoGEP6FOao53gaC8OefSjiOd1idB3C2iQv7NnaP9GDiC",
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

		err = rootcmd.Client.PatchServicePrincipal(rootcmd.OptID, &patchedServicePrincipal, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		// Get the updated Service Principal
		sp, err = rootcmd.Client.GetServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Printf("Service Principal %s updated with new certificate and keyCredential %+v", sp.DisplayName, sp.KeyCredentials)

		// Activate the certificate by its keyId
		err = rootcmd.Client.PatchServicePrincipal(rootcmd.OptID, &models.ServicePrincipal{PreferredTokenSigningKeyThumbprint: OptKeyName}, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
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
