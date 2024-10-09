package serviceprincipalcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// @task: Create a test file for every file in cmd/ that contains a command using serviceprincipalCreate_test.go as a template @all @run

// serviceprincipalDeleteCmd represents the serviceprincipalDelete command
var serviceprincipalDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an serviceprincipal",
	Run: func(cmd *cobra.Command, args []string) {
		err := rootcmd.Client.DeleteServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalDeleteCmd)
}
