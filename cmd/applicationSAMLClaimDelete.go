package cmd

import (
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
			panic("Claim Policy ID is required (use --claimpolicyid)")
		}

		err := Client.DeleteClaimsMappingPolicy(OptClaimPolicyID, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimDeleteCmd)
}
