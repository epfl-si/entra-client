package cmd

import (
	"fmt"

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
		fmt.Println("applicationSAMLClaimUnassign called")
	},
}

func init() {
	applicationSAMLClaimCmd.AddCommand(applicationSAMLClaimUnassignCmd)
}
