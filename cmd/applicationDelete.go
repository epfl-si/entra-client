// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// applicationDeleteCmd represents the applicationDelete command
var applicationDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an application",
	Run: func(cmd *cobra.Command, args []string) {
		err := Client.DeleteApplication(OptID, clientOptions)
		if err != nil {
			printErr(err)
			return
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationDeleteCmd)
}
