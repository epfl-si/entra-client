// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// manifestGetCmd represents the manifestGet command
var manifestGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application manifest by its AppId",
	Run: func(cmd *cobra.Command, args []string) {
		if OptID == "" {
			printErrString("Application AppID is required (use --id)")
			return
		}
		clientOptions.Filter = "appId%20eq%20'" + OptID + "'"
		sps, _, err := Client.GetServicePrincipals(clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		if len(sps) != 1 {
			if len(sps) == 0 {
				printErrString("No service principals found")
			} else {
				printErrString("Ambiguous service principal ID")
			}
			return
		}

		apps, _, err := Client.GetApplications(clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		if len(apps) != 1 {
			if len(apps) == 0 {
				printErrString("No applications found")
			} else {
				printErrString("Ambiguous application ID")
			}
			return
		}

		cmd.Println("{\n \"application\": \"" + OutputJSON(apps[0]) + "\",\n \"servicePrincipal\": \"" + OutputJSON(sps[0]) + "\"\n}")
	},
}

func init() {
	manifestCmd.AddCommand(manifestGetCmd)
	// manifestGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
	// 	// Hide flags for this command
	// 	manifestGetCmd.Flags().MarkHidden("batch")
	// 	manifestGetCmd.Flags().MarkHidden("displayname")
	// 	manifestGetCmd.Flags().MarkHidden("post")
	// 	manifestGetCmd.Flags().MarkHidden("search")
	// 	manifestGetCmd.Flags().MarkHidden("select")
	// 	manifestGetCmd.Flags().MarkHidden("skip")
	// 	manifestGetCmd.Flags().MarkHidden("skiptoken")
	// 	manifestGetCmd.Flags().MarkHidden("top")
	// 	// Call parent help func
	// 	command.Parent().HelpFunc()(command, strings)
	// })
}
