// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationList.goCmd represents the applicationList.go command
var applicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	Run: func(cmd *cobra.Command, args []string) {
		applications, _, err := Client.GetApplications(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, application := range applications {
			fmt.Printf("%s\n", OutputJSON(application))
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationListCmd)
}
