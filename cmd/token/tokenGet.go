// Package cmdtoken is used for token commands
package cmdtoken

import (
	"fmt"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	client "github.com/epfl-si/entra-client/pkg/client"

	"github.com/spf13/cobra"
)

var OptRestricted bool

// tokenCmd represents the token command
var tokenGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Returns a valid token",
	Long: `Get an access token to use with the API

Usage:
  ./ecli token get [--restricted]

Example:
  ./ecli token get               # Will get a token with full API access (but not fully JWT compliant)
  ./ecli token get --restricted  # Will get a token with restricted scope (JWT compliant)
  `,
	Run: func(cmd *cobra.Command, args []string) {

		token, err := client.GetToken(rootcmd.Client.GetClientID(), rootcmd.Client.GetSecret(), rootcmd.Client.GetTenant(), OptRestricted)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		fmt.Println(token)
	},
}

func init() {
	tokenCmd.AddCommand(tokenGetCmd)

	tokenGetCmd.PersistentFlags().BoolVar(&OptRestricted, "restricted", false, "Restricted token")
}
