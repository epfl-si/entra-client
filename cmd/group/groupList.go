package groupcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// groupListCmd represents the groupList command
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		groups, _, err := rootcmd.Client.GetGroups(rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		for _, group := range groups {
			cmd.Println(rootcmd.OutputJSON(group))
		}
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
}
