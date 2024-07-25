// Package cmd provides the commands for the command line apptemplate
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// apptemplateInstantiateCmd represents the apptemplateInstantiate command
var apptemplateInstantiateCmd = &cobra.Command{
	Use:   "instantiate",
	Short: "Instatitate an application template",
	Long: `Instantiate an application  by providing an application template ID and the name of the application.

`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("apptemplateInstantiate called")
		app, sp, err := Client.InstantiateApplicationTemplate(OptID, OptDisplayName, clientOptions)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Applicaiton: %s\n", OutputJSON(app))
		fmt.Printf("ServicePrincipal: %s\n", OutputJSON(sp))
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateInstantiateCmd)
}
