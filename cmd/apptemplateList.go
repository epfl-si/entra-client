// Package cmd provides the commands for the command line apptemplate
package cmd

import (
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
		apptemplates, _, err := Client.GetApplicationTemplates(clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		for _, apptemplate := range apptemplates {
			cmd.Printf("%s\n", OutputJSON(apptemplate))
		}
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateListCmd)
}
