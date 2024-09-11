package cmdapplication

import (
	rootcmd "entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationSAMLClaimDeleteCmd represents the applicationSAMLClaimDelete command
var applicationSAMLClaimDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a claim mapping for a SAML application",
	Long: `Delete a claim mapping for a SAML application

Exqample:

  ./ecli application saml claim delete --claimpolicyid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptClaimPolicyID == "" {
			rootcmd.PrintErrString("Claim Policy ID is required (use --claimpolicyid)")
			return
		}

		err := rootcmd.Client.DeleteClaimsMappingPolicy(OptClaimPolicyID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimDeleteCmd)
	applicationSAMLClaimDeleteCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("batch")
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("search")
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("select")
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("skip")
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLClaimDeleteCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
