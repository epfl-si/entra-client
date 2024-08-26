package cmdcertificate

import (
	"epfl-entra/internal/models"
	"fmt"
	"time"

	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

var localEndDate string

// certificateListCmd represents the certificateList command
var certificateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List certificates",
	Run: func(cmd *cobra.Command, args []string) {
		rootcmd.ClientOptions.Select = "keyCredentials"
		certs := make(map[string]models.KeyCredential, 0)
		apps := make(map[string]string, 0)

		applications, _, err := rootcmd.Client.GetApplications(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, application := range applications {
			if len(application.KeyCredentials) > 0 {
				for _, cert := range application.KeyCredentials {
					if _, ok := certs[cert.KeyID]; ok {
						fmt.Println("Skipping duplicate certificate for application " + cert.KeyID)
						continue
					}
					certs[cert.KeyID] = cert
					apps[cert.KeyID] = application.DisplayName
				}
			}
		}
		sps, _, err := rootcmd.Client.GetServicePrincipals(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, sp := range sps {
			if len(sp.KeyCredentials) > 0 {
				for _, cert := range sp.KeyCredentials {
					if _, ok := certs[cert.KeyID]; ok {
						fmt.Println("Skipping duplicate certificate for service principal " + cert.KeyID)
						continue
					}
					certs[cert.KeyID] = cert
					apps[cert.KeyID] = sp.DisplayName
				}
			}
		}
		for _, cert := range certs {
			if localEndDate != "" {
				endDate, err := time.Parse("2006-01-02", localEndDate)
				if err != nil {
					rootcmd.PrintErr(err)
					return
				}
				if cert.EndDateTime == nil || time.Time(*cert.EndDateTime).After(endDate) {
					continue
				}
			}
			fmt.Printf("KeyID: %s (%s) %s\n", cert.KeyID, time.Time(*cert.EndDateTime).String(), apps[cert.KeyID])
		}
	},
}

func init() {
	certificateCmd.AddCommand(certificateListCmd)

	certificateListCmd.PersistentFlags().StringVar(&localEndDate, "end_date", "", "Limit the certificates to those that expire after this date")

	certificateListCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		certificateListCmd.Flags().MarkHidden("batch")
		certificateListCmd.Flags().MarkHidden("displayname")
		certificateListCmd.Flags().MarkHidden("post")
		certificateListCmd.Flags().MarkHidden("search")
		certificateListCmd.Flags().MarkHidden("select")
		certificateListCmd.Flags().MarkHidden("skip")
		certificateListCmd.Flags().MarkHidden("skiptoken")
		certificateListCmd.Flags().MarkHidden("top")
		// Call parent help func
		certificateListCmd.Parent().HelpFunc()(command, strings)
	})

}
