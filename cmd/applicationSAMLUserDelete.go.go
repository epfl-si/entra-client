// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationSAMLUserDelete.goCmd represents the applicationSAMLUserDelete.go command
var applicationSAMLUserDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a user from a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationSAMLUserDelete called")
	},
}

func init() {
	applicationSAMLUserCmd.AddCommand(applicationSAMLUserDeleteCmd)
}
