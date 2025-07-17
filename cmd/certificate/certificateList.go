package cmdcertificate

import (
	"time"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// certificateListCmd represents the certificateList command
var certificateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List certificates",
	Run: func(cmd *cobra.Command, _ []string) {
		rootcmd.ClientOptions.Select = "keyCredentials,displayName,appId"
		certs, err := rootcmd.Client.GetKeyCredentials("", rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, kcs := range certs {
			for _, cert := range kcs {
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
				cmd.Printf("%s\n", rootcmd.OutputJSON(cert))
				//fmt.Printf("%s (%s) ApplicationID: %s\n", cert.KeyID, time.Time(*cert.EndDateTime).String(), appID)
			}
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
