package cmd

import (
	"github.com/spf13/cobra"
)

// applicationSAMLClaimUnassignCmd represents the applicationSAMLClaimUnassign command
var applicationSAMLClaimUnassignCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Unassign a claim mapping for a SAML application",
	Long: `Unassign a claim mapping for a SAML application

Exqample:

  ./ecli application saml claim unassign --id 2ac47ba8-f2d2-4c9b-9395-3654fc7d2c3 --claimpolicyid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLClaimUnassign called")
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimUnassignCmd)

	applicationSAMLClaimUnassignCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("batch")
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("search")
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("select")
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("skip")
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLClaimUnassignCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
