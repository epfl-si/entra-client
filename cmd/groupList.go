// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// groupListCmd represents the groupList command
var groupListCmd = &cobra.Command{
	Use:   "list",
	Short: "List groups",
	Run: func(cmd *cobra.Command, args []string) {
		groups, _, err := Client.GetGroups(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, group := range groups {
			cmd.Println(OutputJSON(group))
		}
	},
}

func init() {
	groupCmd.AddCommand(groupListCmd)
}
