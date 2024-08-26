package serviceprincipalcmd

import (
	"github.com/spf13/cobra"
)

// serviceprincipalAssociateCmd represents the serviceprincipalAssociate command
var serviceprincipalAssociateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate various data to service principal",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("serviceprincipalAssociate called")
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalAssociateCmd)

	serviceprincipalAssociateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalAssociateCmd.Flags().MarkHidden("batch")
		serviceprincipalAssociateCmd.Flags().MarkHidden("search")
		serviceprincipalAssociateCmd.Flags().MarkHidden("select")
		serviceprincipalAssociateCmd.Flags().MarkHidden("skip")
		serviceprincipalAssociateCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalAssociateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
