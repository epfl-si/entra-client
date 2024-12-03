// Package cmdclaimsmappingpolicy is used for claims mapping policy commands
package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptCmpID is associated with the --cmpid flag
var OptCmpID string

// claimCmd represents the claim command
var claimCmd = &cobra.Command{
	Use:     "claimsmappingpolicy",
	Aliases: []string{"cmp"},
	Short:   "Manage claims mapping policies",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("claim called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(claimCmd)

	claimCmd.PersistentFlags().StringVar(&OptCmpID, "cmpid", "", "Claims Mapping Policy Id to use")
}
