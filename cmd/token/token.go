// Package cmdtoken is used for token commands
package cmdtoken

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Handles token commands",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("token called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(tokenCmd)
}
