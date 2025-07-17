// Package cmdcertificate is used for certificate commands
package cmdcertificate

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

var localEndDate string

// certificateCmd represents the certificate command
var certificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Manage certificates",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println("certificate called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(certificateCmd)

	certificateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		certificateCmd.Flags().MarkHidden("batch")
		certificateCmd.Flags().MarkHidden("displayname")
		certificateCmd.Flags().MarkHidden("post")
		certificateCmd.Flags().MarkHidden("search")
		certificateCmd.Flags().MarkHidden("select")
		certificateCmd.Flags().MarkHidden("skip")
		certificateCmd.Flags().MarkHidden("skiptoken")
		certificateCmd.Flags().MarkHidden("top")
		// Call parent help func
		certificateCmd.Parent().HelpFunc()(command, strings)
	})
}
