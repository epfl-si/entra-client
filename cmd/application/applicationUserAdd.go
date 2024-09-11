package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptUserID is associated with the --userID flag
var OptUserID string

// applicationUserAdd.goCmd represents the applicationUserAdd.go command
var applicationUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user to a SAML application",
	Long: `Add a user to a SAML application

	Syntax:
	  ./ecli application saml user add --id <service_principal_id> --userID <user_id>

	Example:
	  ./ecli application saml user add --id a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d --userID 0f2a8e2d-9c45-4c55-acd9-49c8e278f706
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ID is required (use --id)")
			return
		}

		if OptUserID == "" {
			rootcmd.PrintErrString("UserID is required (use --userID)")
			return
		}

		// assignment := models.AppRoleAssignment{
		// 	AppRoleID:     "00000000-0000-0000-0000-000000000000",
		// 	PrincipalID:   OptUserID,
		// 	PrincipalType: "Group",
		// 	ResourceID:    rootcmd.OptID,
		// }

		// err := rootcmd.Client.AssignAppRoleToServicePrincipal(&assignment, rootcmd.ClientOptions)
		// if err != nil {
		// 	rootcmd.PrintErr(err)
		// 	return
		// }
		err := rootcmd.Client.AddGroupToServicePrincipal(rootcmd.OptID, OptUserID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserAddCmd)

	applicationUserAddCmd.Flags().StringVar(&OptUserID, "userID", "", "ID of the user to add to the SAML application")

	applicationUserAddCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationUserAddCmd.Flags().MarkHidden("batch")
		applicationUserAddCmd.Flags().MarkHidden("search")
		applicationUserAddCmd.Flags().MarkHidden("select")
		applicationUserAddCmd.Flags().MarkHidden("skip")
		applicationUserAddCmd.Flags().MarkHidden("skiptoken")
		applicationUserAddCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
