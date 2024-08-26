package usercmd

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// userListCmd represents the userList command
var userListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all users",
	Run: func(cmd *cobra.Command, args []string) {
		users, _, err := rootcmd.Client.GetUsers(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, user := range users {
			cmd.Println(rootcmd.OutputJSON(user))
		}
	},
}

func init() {
	userCmd.AddCommand(userListCmd)
}
