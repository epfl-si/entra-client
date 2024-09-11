package usercmd

import (
	rootcmd "entra-client/cmd"

	"github.com/spf13/cobra"
)

// userGetCmd represents the userGet command
var userGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a user by ID",
	Run: func(cmd *cobra.Command, args []string) {
		user, err := rootcmd.Client.GetUser(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(user))
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
