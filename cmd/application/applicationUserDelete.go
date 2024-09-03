package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationUserDelete.goCmd represents the applicationUserDelete.go command
var applicationUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user from a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationUserDelete called")
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserDeleteCmd)

	applicationUserDeleteCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationUserDeleteCmd.Flags().MarkHidden("batch")
		applicationUserDeleteCmd.Flags().MarkHidden("search")
		applicationUserDeleteCmd.Flags().MarkHidden("select")
		applicationUserDeleteCmd.Flags().MarkHidden("skip")
		applicationUserDeleteCmd.Flags().MarkHidden("skiptoken")
		applicationUserDeleteCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
