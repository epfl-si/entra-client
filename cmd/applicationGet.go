// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application by ID",
	Run: func(cmd *cobra.Command, args []string) {
		application, err := Client.GetApplication(OptID, clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		cmd.Println(OutputJSON(application))
	},
}

func init() {
	applicationCmd.AddCommand(applicationGetCmd)
}
