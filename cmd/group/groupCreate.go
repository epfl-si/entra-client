package groupcmd

import (
	"encoding/json"
	rootCmd "epfl-entra/cmd"
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
		err := json.Unmarshal([]byte(rootCmd.OptPostData), &group)
		if err != nil {
			rootCmd.PrintErr(err)
			return
		}
		err = rootCmd.Client.CreateGroup(&group, rootCmd.ClientOptions)
		if err != nil {
			rootCmd.PrintErr(err)
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
