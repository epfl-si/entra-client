// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

// groupCreateCmd represents the groupCreate command
var groupCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a group",
	Long: `Create a group whose JSON is passed as argument with --post
	
Example:
  ecli group create --post '{"displayName": "test group AA"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		var group models.Group
		err := json.Unmarshal([]byte(OptPostData), &group)
		if err != nil {
			printErr(err)
			return
		}
		err = Client.CreateGroup(&group, clientOptions)
		if err != nil {
			printErr(err)
			return
		}
	},
}

func init() {
	groupCmd.AddCommand(groupCreateCmd)

	groupCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		groupCreateCmd.Flags().MarkHidden("batch")
		groupCreateCmd.Flags().MarkHidden("search")
		groupCreateCmd.Flags().MarkHidden("select")
		groupCreateCmd.Flags().MarkHidden("skip")
		groupCreateCmd.Flags().MarkHidden("skiptoken")
		groupCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
