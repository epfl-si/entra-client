package cmdclaimsmappingpolicy

import (
	"encoding/json"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// claimPatchCmd represents the claims mapping policy delete command
var claimPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch/modify a claims mapping policy",
	Long: `This command enables you to patch a claims mapping policy.

	Example:
		./ecli claim patch --cmpid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --default
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptCmpID == "" {
			cmd.PrintErr("ID is required (use --cmpid)")
			return
		}
		var cmp models.ClaimsMappingPolicy

		if rootcmd.OptDefault {
			rootcmd.ClientOptions.Default = true
		}

		if rootcmd.OptPostData != "" {
			err := json.Unmarshal([]byte(rootcmd.OptPostData), &cmp)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}
		}

		err := rootcmd.Client.PatchClaimsMappingPolicy(OptCmpID, &cmp, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Println("Claims mapping policy patched successfully.")
	},
}

func init() {
	claimCmd.AddCommand(claimPatchCmd)

	claimPatchCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimPatchCmd.Flags().MarkHidden("batch")
		claimPatchCmd.Flags().MarkHidden("search")
		claimPatchCmd.Flags().MarkHidden("select")
		claimPatchCmd.Flags().MarkHidden("skip")
		claimPatchCmd.Flags().MarkHidden("skiptoken")
		claimPatchCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
