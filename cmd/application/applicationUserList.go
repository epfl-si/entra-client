package cmdapplication

import (
	rootcmd "entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationUserList.goCmd represents the applicationUserList.go command
var applicationUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users for a SAML application given its service principal ID",
	Run: func(cmd *cobra.Command, args []string) {
		sp, err := rootcmd.Client.GetAssignedAppRoles(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		for _, user := range sp {
			cmd.Println(rootcmd.OutputJSON(user))
		}
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserListCmd)
}
