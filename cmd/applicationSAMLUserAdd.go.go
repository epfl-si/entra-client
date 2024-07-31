// Package cmd provides the commands for the command line application
package cmd

import (
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// OptUserID is associated with the --userID flag
var OptUserID string

// applicationSAMLUserAdd.goCmd represents the applicationSAMLUserAdd.go command
var applicationSAMLUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user to a SAML application",
	Long: `Add a user to a SAML application

	Syntax:
	  ./ecli application saml user add --id <service_principal_id> --userID <user_id>

	Example:
	  ./ecli application saml user add --id a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d --userID 0f2a8e2d-9c45-4c55-acd9-49c8e278f706
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationSAMLUserAdd called")
		if OptID == "" {
			panic("Service Principal ID is required (use --id)")
		}

		if OptUserID == "" {
			panic("UserID is required (use --userID)")
		}

		assignment := models.AppRoleAssignment{
			AppRoleID:   "8b292bda-39b6-4b77-849e-887565235bb0",
			PrincipalID: OptUserID,
			// PrincipalType: "User",
			ResourceID: OptID,
		}

		err := Client.AssignAppRoleToServicePrincipal(&assignment, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationSAMLUserCmd.AddCommand(applicationSAMLUserAddCmd)

	applicationSAMLUserAddCmd.Flags().StringVar(&OptUserID, "userID", "", "ID of the user to add to the SAML application")
}
