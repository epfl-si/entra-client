package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimListCmd represents the claimList command
var claimListCmd = &cobra.Command{
	Use:   "list",
	Short: "List claims mapping policies",
	Run: func(cmd *cobra.Command, args []string) {
		claims, _, err := rootcmd.Client.GetClaimsMappingPolicies(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, claim := range claims {
			cmd.Println(rootcmd.OutputJSON(claim))
		}
	},
}

func init() {
	claimCmd.AddCommand(claimListCmd)
}