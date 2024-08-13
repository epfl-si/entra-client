// Package cmd provides the commands for the command line apptemplate
package cmd

import (
	"github.com/spf13/cobra"
)

// apptemplateInstantiateCmd represents the apptemplateInstantiate command
var apptemplateInstantiateCmd = &cobra.Command{
	Use:   "instantiate",
	Short: "Instatitate an application template",
	Long: `Instantiate an application  by providing an application template ID and the name of the application.

`,
	Run: func(cmd *cobra.Command, args []string) {
		app, sp, err := Client.InstantiateApplicationTemplate(OptID, OptDisplayName, clientOptions)
		if err != nil {
			panic(err)
		}
		cmd.Printf("Application: %s\n", OutputJSON(app))
		cmd.Printf("ServicePrincipal: %s\n", OutputJSON(sp))
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateInstantiateCmd)

	apptemplateInstantiateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		apptemplateInstantiateCmd.Flags().MarkHidden("batch")
		apptemplateInstantiateCmd.Flags().MarkHidden("search")
		apptemplateInstantiateCmd.Flags().MarkHidden("select")
		apptemplateInstantiateCmd.Flags().MarkHidden("skip")
		apptemplateInstantiateCmd.Flags().MarkHidden("skiptoken")
		apptemplateInstantiateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
