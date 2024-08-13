// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// applicationSAMLUser.goCmd represents the applicationSAMLUser.go command
var applicationSAMLUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Handle users for a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLUser called")
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLUserCmd)
}
