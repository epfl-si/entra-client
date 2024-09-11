// Package cmdtoken is used for token commands
package cmdtoken

import (
	"fmt"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// tokenCmd represents the token command
var tokenCmd = &cobra.Command{
	Use:   "token",
	Short: "Returns a valid token",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(rootcmd.Client.GetToken())
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(tokenCmd)
}
