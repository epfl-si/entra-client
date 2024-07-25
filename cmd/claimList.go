// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// claimListCmd represents the claimList command
var claimListCmd = &cobra.Command{
	Use:   "list",
	Short: "List claims mapping policies",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("claimList called")
		claims, _, err := Client.GetClaimsMappingPolicies(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, claim := range claims {
			fmt.Printf("%s\n", OutputJSON(claim))
		}
	},
}

func init() {
	claimCmd.AddCommand(claimListCmd)
}
