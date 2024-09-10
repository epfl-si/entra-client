package serviceprincipalcmd

import (
	rootcmd "epfl-entra/cmd"
	"epfl-entra/pkg/entra-client/models"

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
		if OptPrincipalID == "" {
			rootcmd.PrintErrString("PrincipalID is required (use --principalid)")
			return
		}
		if OptAppRoleID == "" {
			rootcmd.PrintErrString("AppRoleID is required (use --approleid)")
			return
		}
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("ID is required (use --id)")
			return
		}

		assignment := &models.AppRoleAssignment{
			PrincipalID:   OptPrincipalID,
			AppRoleID:     OptAppRoleID,
			PrincipalType: "User",
			ResourceID:    rootcmd.OptID,
		}

		err := rootcmd.Client.AssignAppRoleToServicePrincipal(assignment, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	serviceprincipalAssociateCmd.AddCommand(serviceprincipalAssociateUserCmd)

	serviceprincipalAssociateUserCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("batch")
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("search")
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("select")
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("skip")
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalAssociateUserCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
