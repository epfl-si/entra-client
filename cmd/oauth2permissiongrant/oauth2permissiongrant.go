// Package oauth2permissiongrantcmd provides commands for OAuth2 permission grants
package oauth2permissiongrantcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptResourceID is the resource ID (service principal ID of the API)
var OptResourceID string

// oauth2permissiongrantCmd represents the oauth2permissiongrant command
var oauth2permissiongrantCmd = &cobra.Command{
	Use:     "oauth2permissiongrant",
	Aliases: []string{"oauth2pg"},
	Short:   "Manage OAuth2 permission grants",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("oauth2permissiongrant called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(oauth2permissiongrantCmd)

	oauth2permissiongrantCmd.PersistentFlags().StringVar(&OptResourceID, "resourceid", "", "Resource ID (service principal ID of the API)")
}
