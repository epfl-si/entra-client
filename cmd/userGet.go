// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// userGetCmd represents the userGet command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user by ID",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := Client.GetUser(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		cmd.Println(OutputJSON(user))
	},
}

func init() {
	userCmd.AddCommand(userGetCmd)

	userGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		userGetCmd.Flags().MarkHidden("batch")
		userGetCmd.Flags().MarkHidden("search")
		userGetCmd.Flags().MarkHidden("select")
		userGetCmd.Flags().MarkHidden("skip")
		userGetCmd.Flags().MarkHidden("skiptoken")
		userGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
