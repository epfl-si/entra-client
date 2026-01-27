package oauth2permissiongrantcmd

import (
	"fmt"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// oauth2permissiongrantListCmd represents the list command
var oauth2permissiongrantListCmd = &cobra.Command{
	Use:   "list",
	Short: "List OAuth2 permission grants",
	Long: `List OAuth2 permission grants with optional filtering.

Use --appid to filter by application ID (will resolve to service principal object ID).
Use --spid to filter by service principal object ID directly.
Use --resourceid to filter by resource ID (service principal ID of the API).

--appid and --spid are mutually exclusive. --resourceid can be combined with either.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if rootcmd.OptAppID != "" && rootcmd.OptSpID != "" {
			return fmt.Errorf("--appid and --spid are mutually exclusive")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		opts := rootcmd.ClientOptions
		var filters []string

		// If appid is provided, resolve it to the service principal object ID
		if rootcmd.OptAppID != "" {
			spID, err := rootcmd.Client.GetServicePrincipalIDByAppID(rootcmd.OptAppID, models.ClientOptions{})
			if err != nil {
				rootcmd.PrintErr(fmt.Errorf("failed to resolve appid to service principal: %w", err))
				return
			}
			filters = append(filters, fmt.Sprintf("clientId eq '%s'", spID))
		}

		// If spid is provided, use it directly for filtering
		if rootcmd.OptSpID != "" {
			filters = append(filters, fmt.Sprintf("clientId eq '%s'", rootcmd.OptSpID))
		}

		// If resourceid is provided, add it to filters
		if OptResourceID != "" {
			filters = append(filters, fmt.Sprintf("resourceId eq '%s'", OptResourceID))
		}

		// Combine filters with 'and'
		if len(filters) > 0 {
			opts.Filter = filters[0]
			for i := 1; i < len(filters); i++ {
				opts.Filter = fmt.Sprintf("%s and %s", opts.Filter, filters[i])
			}
		}

		grants, _, err := rootcmd.Client.GetOAuth2PermissionGrants(opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, grant := range grants {
			cmd.Println(rootcmd.OutputJSON(grant))
		}
	},
}

func init() {
	oauth2permissiongrantCmd.AddCommand(oauth2permissiongrantListCmd)
}
