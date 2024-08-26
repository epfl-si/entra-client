package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationSAMLClaimAssignCmd represents the applicationSAMLClaimAssign command
//
// ️⚠⚠⚠
// Once you assign a Claims Mapping Policy to an Enterprise Application using Graph API or Powershell, you can no longer manage SAML claims in the Azure AD portal. Accessing the SAML claims configuration page in the Azure AD portal will result in the following error message:
// "This configuration was overwritten by a claim mapping policy created via Graph/PowerShell."
// (source: https://www.twilio.com/docs/flex/admin-guide/setup/sso-configuration/azure-ad/custom-azure-ad-attributes-as-saml-claims)
var applicationSAMLClaimAssignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a claim mapping for a SAML application",
	Long: `Assign a claim mapping for a SAML application

Example:

  ./ecli application saml claim assign --id 2ac47ba8-f2d2-4c9b-9395-3654fc7d2c3 --claimpolicyid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLClaimAssign called")
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimAssignCmd)
	applicationSAMLClaimAssignCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLClaimAssignCmd.Flags().MarkHidden("top")
		applicationSAMLClaimAssignCmd.Flags().MarkHidden("skip")
		applicationSAMLClaimAssignCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLClaimAssignCmd.Flags().MarkHidden("select")
		applicationSAMLClaimAssignCmd.Flags().MarkHidden("search")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
