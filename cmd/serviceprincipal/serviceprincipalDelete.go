package serviceprincipalcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// serviceprincipalDeleteCmd represents the serviceprincipalDelete command
var serviceprincipalDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an serviceprincipal",
	Run: func(cmd *cobra.Command, args []string) {
		err := rootcmd.Client.DeleteServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalDeleteCmd)
}
