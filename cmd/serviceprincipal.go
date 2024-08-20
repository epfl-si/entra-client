// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceprincipalCmd represents the serviceprincipal command
var serviceprincipalCmd = &cobra.Command{
	Use:   "serviceprincipal",
	Short: "Manage service principals",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("serviceprincipal called")
	},
}

func init() {
	rootCmd.AddCommand(serviceprincipalCmd)

	serviceprincipalCmd.PersistentFlags().StringVar(&OptPrincipalID, "principalid", "", "ID of the principal")
	serviceprincipalCmd.PersistentFlags().StringVar(&OptAppRoleID, "approleid", "", "ID of the AppRole")
}
