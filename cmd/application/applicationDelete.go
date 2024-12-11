package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationDeleteCmd represents the applicationDelete command
var applicationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an application",
	Run: func(cmd *cobra.Command, args []string) {
		err := rootcmd.Client.DeleteApplication(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationDeleteCmd)
}
