package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/spf13/cobra"
)

// OptUserAll is a global variable to store the --all flag
var OptUserAll bool

// applicationUserDelete.goCmd represents the applicationUserDelete.go command
var applicationUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user/group from the authorized list for an application",
	Long: `Delete a user/group from the authorized list for an application
	
	Syntax:
	  ./ecli application user delete [--spid <service_principal_id>|--appid <application_id>] [--userid <user_id>|--all]

	Example:
	  ./ecli application user delete --spid a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d --userid 0f2a8e2d-9c45-4c55-acd9-49c8e278f706
	  ./ecli application user delete --appid 17c28601-f637-405a-88c0-591322fd5437 --userid "AAD_All Hosts Users"
	  ./ecli application user delete --appid 17c28601-f637-405a-88c0-591322fd5437 --all
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptSpID == "" && rootcmd.OptAppID == "" {
			cmd.PrintErr("Service Principal or application ID is required (use --spid or --appid)\n")
			return
		}

		if OptUserID == "" && !OptUserAll {
			cmd.PrintErr("UserID (use --userid) OR --all required\n")
			return
		}

		var userID string
		if OptUserAll {
			userID = ""
		} else {
			userID = OptUserID
		}

		var id string
		if rootcmd.OptAppID != "" {
			id = rootcmd.OptAppID
			sp, err := rootcmd.Client.GetServicePrincipalByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
			if err != nil {
				rootcmd.PrintErr("No service principal found for this appID: " + err.Error() + "\n")
				return
			}
			id = sp.ID
		} else {
			id = rootcmd.OptSpID
		}

		err := rootcmd.Client.RemoveGroupFromServicePrincipal(id, userID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserDeleteCmd)
	applicationUserDeleteCmd.Flags().BoolVar(&OptUserAll, "all", false, "Select all users")

	applicationUserDeleteCmd.MarkFlagsMutuallyExclusive("spid", "appid")

	applicationUserDeleteCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationUserDeleteCmd.Flags().MarkHidden("batch")
		applicationUserDeleteCmd.Flags().MarkHidden("search")
		applicationUserDeleteCmd.Flags().MarkHidden("select")
		applicationUserDeleteCmd.Flags().MarkHidden("skip")
		applicationUserDeleteCmd.Flags().MarkHidden("skiptoken")
		applicationUserDeleteCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
