package usercmd

import (
	"encoding/json"
	rootcmd "epfl-entra/cmd"
	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

// userCreateCmd represents the userCreate command
var userCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a user",
	Long: `Create a user whose JSON is passed as argument with --post
	
Example:
  ecli user create --post '{"displayName": "test user AA"}'
`,
	Run: func(cmd *cobra.Command, args []string) {
		var app models.User
		err := json.Unmarshal([]byte(rootcmd.OptPostData), &app)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		err = rootcmd.Client.CreateUser(&app, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	userCmd.AddCommand(userCreateCmd)

	userCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		userCreateCmd.Flags().MarkHidden("batch")
		userCreateCmd.Flags().MarkHidden("search")
		userCreateCmd.Flags().MarkHidden("select")
		userCreateCmd.Flags().MarkHidden("skip")
		userCreateCmd.Flags().MarkHidden("skiptoken")
		userCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
