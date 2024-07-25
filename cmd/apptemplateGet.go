// Package cmd provides the commands for the command line apptemplate
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apptemplateGetCmd represents the apptemplateGet command
var apptemplateGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application template by ID",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apptemplateGet called")
		application, err := Client.GetApplicationTemplate(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		fmt.Printf("ApplicationTemplate: %s\n", OutputJSON(application))
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateGetCmd)
}
