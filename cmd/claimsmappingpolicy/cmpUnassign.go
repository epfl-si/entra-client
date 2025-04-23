package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimUnassignCmd represents the claims mapping policy assign command
var claimUnassignCmd = &cobra.Command{
	Use:   "unassign",
	Short: "Unassign the claims mapping policy linked to a service principal",
	Long: `This command enables you to unassign the claims mapping policy linked to a service principal.

	Example:
		./ecli claim unassign --spid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --cmpid 76cbc426-f398-44e5-993b-b07186066505
	
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

		err := rootcmd.Client.UnassignClaimsMappingPolicy(rootcmd.OptSpID, OptCmpID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

	},
}

func init() {
	claimCmd.AddCommand(claimUnassignCmd)

	claimUnassignCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimUnassignCmd.Flags().MarkHidden("batch")
		claimUnassignCmd.Flags().MarkHidden("search")
		claimUnassignCmd.Flags().MarkHidden("select")
		claimUnassignCmd.Flags().MarkHidden("skip")
		claimUnassignCmd.Flags().MarkHidden("skiptoken")
		claimUnassignCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
