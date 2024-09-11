// Package cmdclaim is used for claim commands
package cmdclaim

import (
	rootcmd "entra-client/cmd"

	"github.com/spf13/cobra"
)

// claimCmd represents the claim command
var claimCmd = &cobra.Command{
	Use:   "claim",
	Short: "Manage SAML2 claims",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("claim called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(claimCmd)
}
