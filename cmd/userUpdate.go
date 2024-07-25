// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// userUpdateCmd represents the userUpdate command
var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("userUpdate called")
		var app models.User
		err := json.Unmarshal([]byte(OptPostData), &app)
		if err != nil {
			panic(err)
		}
		err = Client.UpdateUser(&app, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	userCmd.AddCommand(userUpdateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// userUpdateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// userUpdateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
