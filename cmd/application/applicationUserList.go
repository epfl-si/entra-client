package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationUserList.goCmd represents the applicationUserList.go command
var applicationUserListCmd = &cobra.Command{
	Use:   "list",
	Short: "List users/groups authorized for an application given its service principal ID or application ID",
	Long: `List users/groups authorized for an application

	Syntax:
	  ./ecli application user list [--spid <service_principal_id>|--appid <application_id>]

	Example:
	  ./ecli application user list --spid a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d
	  ./ecli application user list --spid 806b30e8-1e91-4410-80c8-265ddfc03748
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptSpID == "" && rootcmd.OptAppID == "" {
			cmd.PrintErr("Service Principal or application ID is required (use --spid or --appid)\n")
			return
		}

		var id string
		if rootcmd.OptAppID != "" {
			id = rootcmd.OptAppID
			sp, err := rootcmd.Client.GetServicePrincipalByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
			if err != nil {
				cmd.PrintErr("No service principal found for this appID: " + err.Error() + "\n")
				return
			}
			id = sp.ID
		} else {
			id = rootcmd.OptSpID
		}

		gps, err := rootcmd.Client.GetGroupsFromServicePrincipal(id, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		for _, user := range gps {
			cmd.Println(rootcmd.OutputJSON(user))
		}
	},
}

func init() {
	applicationUserCmd.AddCommand(applicationUserListCmd)
	applicationUserListCmd.MarkFlagsMutuallyExclusive("spid", "appid")
}
