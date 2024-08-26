// Package serviceprincipalcmd is used for service principal commandsj
package serviceprincipalcmd

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// OptPrincipalID is the ID of the principal
var OptPrincipalID string

// OptAppRoleID is the ID of the AppRole
var OptAppRoleID string

// serviceprincipalCmd represents the serviceprincipal command
var serviceprincipalCmd = &cobra.Command{
	Use:   "serviceprincipal",
	Short: "Manage service principals",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("serviceprincipal called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(serviceprincipalCmd)

	serviceprincipalCmd.PersistentFlags().StringVar(&OptPrincipalID, "principalid", "", "ID of the principal")
	serviceprincipalCmd.PersistentFlags().StringVar(&OptAppRoleID, "approleid", "", "ID of the AppRole")
}
