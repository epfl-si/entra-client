// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

// serviceprincipalPatchCmd represents the serviceprincipalPatch command
var serviceprincipalPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch a ServicePrincipal",
	Run: func(cmd *cobra.Command, args []string) {
		var app models.ServicePrincipal
		err := json.Unmarshal([]byte(OptPostData), &app)
		if err != nil {
			panic(err)
		}
		err = Client.PatchServicePrincipal(OptID, &app, clientOptions)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serviceprincipalPatchCmd)

	serviceprincipalPatchCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalPatchCmd.Flags().MarkHidden("batch")
		serviceprincipalPatchCmd.Flags().MarkHidden("search")
		serviceprincipalPatchCmd.Flags().MarkHidden("select")
		serviceprincipalPatchCmd.Flags().MarkHidden("skip")
		serviceprincipalPatchCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalPatchCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
