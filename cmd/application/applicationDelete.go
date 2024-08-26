package cmdapplication

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// @task: Create a test file for every file in cmd/ that contains a command using applicationCreate_test.go as a template @all @run

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
