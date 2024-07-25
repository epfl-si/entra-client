// Package cmd provides the commands for the command line apptemplate
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apptemplateListCmd represents the apptemplateList command
var apptemplateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List apptemplate templates",
	Long: `List application templates
	
	Example:
	  ./ecli apptemplate list --select id,displayname
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apptemplateList called")
		apptemplates, _, err := Client.GetApplicationTemplates(clientOptions)
		if err != nil {
			panic(err)
		}

		for _, apptemplate := range apptemplates {
			fmt.Printf("%s\n", OutputJSON(apptemplate))
		}
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateListCmd)
}
