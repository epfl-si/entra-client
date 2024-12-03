package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimGetCmd represents the claims mapping policy get command
var claimGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a claims mapping policy",
	Long: `This command enables you to get a claims mapping policy.

	Example:
		./ecli claim Get --spid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
	
}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" {
			rootcmd.PrintErrString("ID is required (use --cmpid)")
			return
		}

		cmp, err := rootcmd.Client.GetClaimsMappingPolicy(OptCmpID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("%s\n", rootcmd.OutputJSON(cmp))

	},
}

func init() {
	claimCmd.AddCommand(claimGetCmd)

	claimGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimGetCmd.Flags().MarkHidden("batch")
		claimGetCmd.Flags().MarkHidden("search")
		claimGetCmd.Flags().MarkHidden("select")
		claimGetCmd.Flags().MarkHidden("skip")
		claimGetCmd.Flags().MarkHidden("skiptoken")
		claimGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
