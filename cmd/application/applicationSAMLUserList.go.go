package cmdapplication

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// applicationSAMLUserList.goCmd represents the applicationSAMLUserList.go command
var applicationSAMLUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users for a SAML application given its service principal ID",
	Run: func(cmd *cobra.Command, args []string) {
		sp, err := rootcmd.Client.GetServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		for _, user := range sp.AppRoles {
			cmd.Println(rootcmd.OutputJSON(user))
		}
	},
}

func init() {
	applicationSAMLUserCmd.AddCommand(applicationSAMLUserListCmd)
}
