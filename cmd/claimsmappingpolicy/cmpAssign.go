package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimAssignCmd represents the claims mapping policy assign command
var claimAssignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign a claims mapping policy to a service principal",
	Long: `This command enables you to Assign a claims mapping policy.

	Example:
		./ecli claim assign --spid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
	
}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" {
			cmd.PrintErr("ID is required (use --cmpid)")
			return
		}

		if rootcmd.OptSpID == "" {
			cmd.PrintErr("ID is required (use --spid)")
			return
		}

		err := rootcmd.Client.AssignClaimsMappingPolicy(OptCmpID, rootcmd.OptSpID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		cmd.Println("Claims mapping policy assigned successfully.")

	},
}

func init() {
	claimCmd.AddCommand(claimAssignCmd)

	claimAssignCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimAssignCmd.Flags().MarkHidden("batch")
		claimAssignCmd.Flags().MarkHidden("search")
		claimAssignCmd.Flags().MarkHidden("select")
		claimAssignCmd.Flags().MarkHidden("skip")
		claimAssignCmd.Flags().MarkHidden("skiptoken")
		claimAssignCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
