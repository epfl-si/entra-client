// Package cmdclaimsmappingpolicy is used for claims mapping policy commands
package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptCmpID is associated with the --cmpid flag
var OptCmpID string

// OptDefault is associated with the --default flag
var OptDefault = false

// claimCmd represents the claim command
var claimCmd = &cobra.Command{
	Use:     "claimsmappingpolicy",
	Aliases: []string{"cmp"},
	Short:   "Manage claims mapping policies",
	Long: `This command enables you to manage claims mapping policies.
This command also accepts the alias 'cmp'.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("claim called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(claimCmd)

	claimCmd.PersistentFlags().BoolVar(&OptDefault, "default", false, "Use a default EPFL's value")
	claimCmd.PersistentFlags().StringVar(&OptCmpID, "cmpid", "", "Claims Mapping Policy Id to use")
}
