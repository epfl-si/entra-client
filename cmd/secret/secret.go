// Package cmdsecret is used for secret commands
package cmdsecret

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

var localEndDate string

// secretCmd represents the secret command
var secretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Manage secrets",
	Run: func(cmd *cobra.Command, _ []string) {
		cmd.Println("secret called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(secretCmd)

	secretCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		secretCmd.Flags().MarkHidden("batch")
		secretCmd.Flags().MarkHidden("displayname")
		secretCmd.Flags().MarkHidden("post")
		secretCmd.Flags().MarkHidden("search")
		secretCmd.Flags().MarkHidden("select")
		secretCmd.Flags().MarkHidden("skip")
		secretCmd.Flags().MarkHidden("skiptoken")
		secretCmd.Flags().MarkHidden("top")
		// Call parent help func
		secretCmd.Parent().HelpFunc()(command, strings)
	})
}
