// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// serviceprincipalGetCmd represents the serviceprincipalGet command
var serviceprincipalGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a service principal by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if OptID == "" {
			panic("Service Principal ID is required (use --id)")
		}
		user, err := Client.GetServicePrincipal(OptID, clientOptions)
		if err != nil {
			panic(err)
		}

		cmd.Println(OutputJSON(user))
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalGetCmd)

	serviceprincipalGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalGetCmd.Flags().MarkHidden("batch")
		serviceprincipalGetCmd.Flags().MarkHidden("search")
		serviceprincipalGetCmd.Flags().MarkHidden("select")
		serviceprincipalGetCmd.Flags().MarkHidden("skip")
		serviceprincipalGetCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
