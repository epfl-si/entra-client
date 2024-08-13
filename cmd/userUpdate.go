// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

// userUpdateCmd represents the userUpdate command
var userUpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update a user",
	Run: func(cmd *cobra.Command, args []string) {
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

	userUpdateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		userUpdateCmd.Flags().MarkHidden("batch")
		userUpdateCmd.Flags().MarkHidden("search")
		userUpdateCmd.Flags().MarkHidden("select")
		userUpdateCmd.Flags().MarkHidden("skip")
		userUpdateCmd.Flags().MarkHidden("skiptoken")
		userUpdateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
