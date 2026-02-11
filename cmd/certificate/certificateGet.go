package cmdcertificate

import (
	"time"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// certificateGetCmd represents the certificateGetByAppID command
var certificateGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get certificates for a specific application",
	Long: `Get all key credentials (certificates) for a specific application by AppID

Example:
  ecli certificate get --appid <app-id>
  ecli certificate get --appid <app-id> --end_date 2024-12-31`,
	Run: func(cmd *cobra.Command, _ []string) {
		if rootcmd.OptAppID == "" {
			cmd.PrintErr("Application ID is required (use --appid)\n")
			return
		}

		certs, err := rootcmd.Client.GetKeyCredentialsByAppID("", rootcmd.OptAppID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
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
			cmd.Printf("%s\n", rootcmd.OutputJSON(cert))
		}
	},
}

func init() {
	certificateCmd.AddCommand(certificateGetCmd)

	certificateGetCmd.PersistentFlags().StringVar(&localEndDate, "end_date", "", "Limit the certificates to those that expire before this date (YYYY-MM-DD)")

	certificateGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		certificateGetCmd.Flags().MarkHidden("batch")
		certificateGetCmd.Flags().MarkHidden("id")
		certificateGetCmd.Flags().MarkHidden("post")
		certificateGetCmd.Flags().MarkHidden("search")
		certificateGetCmd.Flags().MarkHidden("select")
		certificateGetCmd.Flags().MarkHidden("skip")
		certificateGetCmd.Flags().MarkHidden("skiptoken")
		certificateGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
