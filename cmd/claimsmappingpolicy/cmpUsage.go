package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimUsageCmd represents the claims mapping policy usage command
var claimUsageCmd = &cobra.Command{
	Use:   "usage",
	Short: "List usage of a claims mapping policy",
	Long: `This command enables you to list usage of a claims mapping policy.

	Example:
		./ecli claim usage  --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
	
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" {
			rootcmd.PrintErrString("ID is required (use --cmpid)")
			return
		}

		cmps, err := rootcmd.Client.ListUsageClaimsMappingPolicy(OptCmpID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		if len(cmps) == 0 {
			cmd.Println("No usage found for the specified claims mapping policy.")
			return
		}

		for _, cmp := range cmps {
			cmd.Println(rootcmd.OutputJSON(cmp))
		}

	},
}

func init() {
	claimCmd.AddCommand(claimUsageCmd)

	claimUsageCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimUsageCmd.Flags().MarkHidden("batch")
		claimUsageCmd.Flags().MarkHidden("search")
		claimUsageCmd.Flags().MarkHidden("select")
		claimUsageCmd.Flags().MarkHidden("skip")
		claimUsageCmd.Flags().MarkHidden("skiptoken")
		claimUsageCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
