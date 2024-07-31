// Package cmd provides the commands for the command line application
package cmd

import (
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// serviceprincipalAssociateUserCmd represents the serviceprincipalAssociateUser command
var serviceprincipalAssociateUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Associate an user to a service principal",
	Long: `Associate an user to a service principal:

Example:
  ./ecli serviceprincipal associate user --principalid "0f2a8e2d-9c45-4c55-acd9-49c8e278f706" --approleid "3a84e31e-bffa-470f-b9e6-754a61e4dc63" --id c40b288c-0239-48d2-993b-892bdc721c00

  Where --principalid is the ID of the user, --id is the ID of the service principal and "3a84e31e-bffa-470f-b9e6-754a61e4dc63" is the ID of the AppRole (in this case, the AppRole is "Admin,WAAD").
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serviceprincipalAssociateUser called")
		if OptPrincipalID == "" {
			panic("PrincipalID is required (use --principalid)")
		}
		if OptAppRoleID == "" {
			panic("AppRoleID is required (use --approleid)")
		}
		if OptID == "" {
			panic("ID is required (use --id)")
		}

		assignment := &models.AppRoleAssignment{
			PrincipalID:   OptPrincipalID,
			AppRoleID:     OptAppRoleID,
			PrincipalType: "User",
			ResourceID:    OptID,
		}

		err := Client.AssignAppRoleToServicePrincipal(assignment, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	serviceprincipalAssociateCmd.AddCommand(serviceprincipalAssociateUserCmd)
}
