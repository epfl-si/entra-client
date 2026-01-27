package serviceprincipalcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// serviceprincipalGetAppRoleCmd represents the get-approle command
var serviceprincipalGetAppRoleCmd = &cobra.Command{
	Use:   "get-approle",
	Short: "Get app role assignments for a service principal",
	Long: `List all app role assignments granted to a service principal.

Requires either --spid (service principal object ID) or --appid (application ID).

Examples:
  ./ecli serviceprincipal get-approle --spid a8ff0bc1-3046-43d8-a4b1-d8c42fd6623d
  ./ecli serviceprincipal get-approle --appid 00000000-0000-0000-0000-000000000000`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptSpID == "" && rootcmd.OptAppID == "" {
			cmd.PrintErr("Service Principal ID or Application ID is required (use --spid or --appid)\n")
			return
		}

		var id string
		if rootcmd.OptAppID != "" {
			sp, err := rootcmd.Client.GetServicePrincipalByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
			if err != nil {
				cmd.PrintErr("No service principal found for this appID: " + err.Error() + "\n")
				return
			}
			id = sp.ID
		} else {
			id = rootcmd.OptSpID
		}

		assignments, err := rootcmd.Client.GetAssignmentsFromServicePrincipal(id, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, assignment := range assignments {
			cmd.Println(rootcmd.OutputJSON(assignment))
		}
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalGetAppRoleCmd)
	serviceprincipalGetAppRoleCmd.MarkFlagsMutuallyExclusive("spid", "appid")
}
