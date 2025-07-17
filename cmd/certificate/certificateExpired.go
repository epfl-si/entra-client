package cmdcertificate

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// certificateExpiredCmd represents the certificateExpired command
var certificateExpiredCmd = &cobra.Command{
	Use:   "expired",
	Short: "List expired certificates",
	Run: func(cmd *cobra.Command, _ []string) {
		rootcmd.ClientOptions.Select = "keyCredentials,displayName,appId"
		certs, err := rootcmd.Client.GetExpiredKeyCredentials(localEndDate, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, kcs := range certs {
			for _, cert := range kcs {
				cmd.Printf("%s\n", rootcmd.OutputJSON(cert))
				//fmt.Printf("%s (%s) in %d days: %s (%s)\n", cert.KeyID, time.Time(*cert.EndDateTime).String(), cert.RemainingDays, cert.AppDisplayName, cert.AppID)
			}
		}
	},
}

func init() {
	certificateCmd.AddCommand(certificateExpiredCmd)

	certificateExpiredCmd.PersistentFlags().StringVar(&localEndDate, "end_date", "", "Limit the certificates to those that expire after this date")

	certificateExpiredCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		certificateExpiredCmd.Flags().MarkHidden("batch")
		certificateExpiredCmd.Flags().MarkHidden("displayname")
		certificateExpiredCmd.Flags().MarkHidden("post")
		certificateExpiredCmd.Flags().MarkHidden("search")
		certificateExpiredCmd.Flags().MarkHidden("select")
		certificateExpiredCmd.Flags().MarkHidden("skip")
		certificateExpiredCmd.Flags().MarkHidden("skiptoken")
		certificateExpiredCmd.Flags().MarkHidden("top")
		// Call parent help func
		certificateExpiredCmd.Parent().HelpFunc()(command, strings)
	})

}
