package cmdsecret

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// secretListCmd represents the secretList command
var secretListCmd = &cobra.Command{
	Use:   "list",
	Short: "List secrets",
	Run: func(cmd *cobra.Command, _ []string) {
		rootcmd.ClientOptions.Select = "keyCredentials,displayName,appId"
		certs, err := rootcmd.Client.GetKeyCredentials(localEndDate, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, kcs := range certs {
			for _, cert := range kcs {
				cmd.Printf("%s\n", rootcmd.OutputJSON(cert))
				//fmt.Printf("%s (%s) ApplicationID: %s\n", cert.KeyID, time.Time(*cert.EndDateTime).String(), appID)
			}
		}
	},
}

func init() {
	secretCmd.AddCommand(secretListCmd)

	secretListCmd.PersistentFlags().StringVar(&localEndDate, "end_date", "", "Limit the secrets to those that expire after this date")

	secretListCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		secretListCmd.Flags().MarkHidden("batch")
		secretListCmd.Flags().MarkHidden("displayname")
		secretListCmd.Flags().MarkHidden("post")
		secretListCmd.Flags().MarkHidden("search")
		secretListCmd.Flags().MarkHidden("select")
		secretListCmd.Flags().MarkHidden("skip")
		secretListCmd.Flags().MarkHidden("skiptoken")
		secretListCmd.Flags().MarkHidden("top")
		// Call parent help func
		secretListCmd.Parent().HelpFunc()(command, strings)
	})

}
