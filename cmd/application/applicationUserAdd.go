package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationUserAdd.goCmd represents the applicationUserAdd.go command
var applicationUserAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a user/group to the authorized list for an application",
	Long: `Add a user/group to the authorized list for an application

	Syntax:
	  ./ecli application user add [--spid <service_principal_id>|--appid <application_id>] --userid <user_id/display_name>

	Example:
	  ./ecli application user add --spid a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d --userid 0f2a8e2d-9c45-4c55-acd9-49c8e278f706
	  ./ecli application user add --appid 17c28601-f637-405a-88c0-591322fd5437 --userid "AAD_All Hosts Users"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptSpID == "" && rootcmd.OptAppID == "" {
			rootcmd.PrintErrString("Service Principal or application ID is required (use --spid or --appid)")
			return
		}

		if OptUserID == "" {
			rootcmd.PrintErrString("UserID is required (use --userid)")
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

		var id string
		if rootcmd.OptAppID != "" {
			id = rootcmd.OptAppID
			sp, err := rootcmd.Client.GetServicePrincipalByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
			if err != nil {
				rootcmd.PrintErrString("No service principal found for this appID: " + err.Error())
				return
			}
			id = sp.ID
		} else {
			id = rootcmd.OptSpID
		}

		err := rootcmd.Client.AddGroupToServicePrincipal(id, OptUserID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserAddCmd)

	applicationUserAddCmd.MarkFlagsMutuallyExclusive("spid", "appid")

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
