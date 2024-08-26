package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationSAMLUserDelete.goCmd represents the applicationSAMLUserDelete.go command
var applicationSAMLUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user from a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLUserDelete called")
	},
}

func init() {
	applicationSAMLUserCmd.AddCommand(applicationSAMLUserDeleteCmd)

	applicationSAMLUserDeleteCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("batch")
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("search")
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("select")
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("skip")
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLUserDeleteCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
