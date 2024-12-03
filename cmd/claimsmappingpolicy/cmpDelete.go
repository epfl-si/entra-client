package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimDeleteCmd represents the claims mapping policy delete command
var claimDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a claim mapping policy",
	Long: `This command enables you to delete a claims mapping policy.

	Example:
		./ecli claim delete --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" {
			rootcmd.PrintErrString("ID is required (use --cmpid)")
			return
		}

		err := rootcmd.Client.DeleteClaimsMappingPolicy(OptCmpID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println("Claims mapping policy deleted successfully.")
	},
}

func init() {
	claimCmd.AddCommand(claimDeleteCmd)

	claimDeleteCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimDeleteCmd.Flags().MarkHidden("batch")
		claimDeleteCmd.Flags().MarkHidden("search")
		claimDeleteCmd.Flags().MarkHidden("select")
		claimDeleteCmd.Flags().MarkHidden("skip")
		claimDeleteCmd.Flags().MarkHidden("skiptoken")
		claimDeleteCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
