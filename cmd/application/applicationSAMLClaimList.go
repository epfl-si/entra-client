package cmdapplication

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// applicationSAMLClaimListCmd represents the applicationSAMLClaimList command
var applicationSAMLClaimListCmd = &cobra.Command{
	Use:   "list",
	Short: "List claims mapping for a SAML application",
	Long: `List claims mapping for a SAML application

Example:

	  ./ecli application saml claim list --id 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ID is required (use --id)")
			return
		}

		cmps, _, err := rootcmd.Client.GetClaimsMappingPoliciesForServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, cmp := range cmps {
			cmd.Println(rootcmd.OutputJSON(cmp))
		}
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimListCmd)
}