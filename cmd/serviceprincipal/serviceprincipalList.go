package serviceprincipalcmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// serviceprincipalListCmd represents the serviceprincipalList command
var serviceprincipalListCmd = &cobra.Command{
	Use:   "list",
	Short: "List ServicePrincipals",
	Run: func(cmd *cobra.Command, args []string) {
		sps, _, err := rootcmd.Client.GetServicePrincipals(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, sp := range sps {
			cmd.Println(rootcmd.OutputJSON(sp))
		}
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalListCmd)
}
