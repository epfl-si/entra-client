// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceprincipalListCmd represents the serviceprincipalList command
var serviceprincipalListCmd = &cobra.Command{
	Use:   "list",
	Short: "List ServicePrincipals",
	Run: func(cmd *cobra.Command, args []string) {
		sps, _, err := Client.GetServicePrincipals(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, sp := range sps {
			cmd.Println(OutputJSON(sp))
		}
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalListCmd)
}
