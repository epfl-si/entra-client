package cmd

import (
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
		if OptID == "" {
			panic("Service Principal ID is required (use --id)")
		}

		cmps, _, err := Client.GetClaimsMappingPoliciesForServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		for _, cmp := range cmps {
			cmd.Println(OutputJSON(cmp))
		}
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimListCmd)
}
