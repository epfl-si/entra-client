// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OptPrincipalID is the ID of the principal
var OptPrincipalID string

// OptAppRoleID is the ID of the AppRole
var OptAppRoleID string

// serviceprincipalAssociateCmd represents the serviceprincipalAssociate command
var serviceprincipalAssociateCmd = &cobra.Command{
	Use:   "associate",
	Short: "Associate various data to service principal",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serviceprincipalAssociate called")
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalAssociateCmd)

	rootCmd.PersistentFlags().StringVar(&OptPrincipalID, "principalid", "", "ID of the principal")
	rootCmd.PersistentFlags().StringVar(&OptAppRoleID, "approleid", "", "ID of the AppRole")
}
