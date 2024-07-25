// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// userListCmd represents the userList command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userList called")
		users, _, err := Client.GetUsers(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, user := range users {
			fmt.Printf("User: %s\n", OutputJSON(user))
		}
	},
}

func init() {
	userCmd.AddCommand(userListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
