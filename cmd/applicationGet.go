// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationGet called")
		application, err := Client.GetApplication(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Application: %s\n", OutputJSON(application))
	},
}

func init() {
	applicationCmd.AddCommand(applicationGetCmd)
}
