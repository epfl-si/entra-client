// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// userListCmd represents the userList command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		users, _, err := Client.GetUsers(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, user := range users {
			cmd.Println(OutputJSON(user))
		}
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}
