package groupcmd

import (
	rootCmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// groupListCmd represents the groupList command
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		groups, _, err := rootCmd.Client.GetGroups(rootCmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		for _, group := range groups {
			cmd.Println(rootCmd.OutputJSON(group))
		}
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
}
