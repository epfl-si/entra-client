package groupcmd

import (
	"fmt"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptGroupID is the group ID for app role assignments
var OptGroupID string

// groupGetAppRoleCmd represents the get-approle command
var groupGetAppRoleCmd = &cobra.Command{
	Use:   "get-approle",
	Short: "Get app role assignments for a group",
	Long: `List all app role assignments granted to a group.

Requires --groupid parameter to specify the group.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if OptGroupID == "" {
			return fmt.Errorf("--groupid is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		assignments, _, err := rootcmd.Client.GetGroupAppRoleAssignments(OptGroupID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, assignment := range assignments {
			cmd.Println(rootcmd.OutputJSON(assignment))
		}
	},
}

func init() {
	groupCmd.AddCommand(groupGetAppRoleCmd)

	groupGetAppRoleCmd.Flags().StringVar(&OptGroupID, "groupid", "", "Group ID (required)")
}
