// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// serviceprincipalCmd represents the serviceprincipal command
var serviceprincipalCmd = &cobra.Command{
	Use:   "serviceprincipal",
	Short: "Manage service principasl",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("serviceprincipal called")
	},
}

func init() {
	rootCmd.AddCommand(serviceprincipalCmd)
}
