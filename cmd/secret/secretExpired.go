package cmdsecret

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// secretExpiredCmd represents the secretExpired command
var secretExpiredCmd = &cobra.Command{
	Use:   "expired",
	Short: "List expired secrets",
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
	secretCmd.AddCommand(secretExpiredCmd)

	secretExpiredCmd.PersistentFlags().StringVar(&localEndDate, "end_date", "", "Limit the secrets to those that expire after this date")

	secretExpiredCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		secretExpiredCmd.Flags().MarkHidden("batch")
		secretExpiredCmd.Flags().MarkHidden("displayname")
		secretExpiredCmd.Flags().MarkHidden("post")
		secretExpiredCmd.Flags().MarkHidden("search")
		secretExpiredCmd.Flags().MarkHidden("select")
		secretExpiredCmd.Flags().MarkHidden("skip")
		secretExpiredCmd.Flags().MarkHidden("skiptoken")
		secretExpiredCmd.Flags().MarkHidden("top")
		// Call parent help func
		secretExpiredCmd.Parent().HelpFunc()(command, strings)
	})

}
