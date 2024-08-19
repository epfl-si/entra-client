// Package cmd provides the commands for the command line application
package cmd

import (
	"epfl-entra/internal/models"
	"errors"

	"github.com/spf13/cobra"
)

// applicationCreateCmd represents the applicationCreate command
var applicationCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an application",
	Long: `Create an application
	
Example:
  ./ecli application create --displayname "AA test app create" --home_uri "https://www.epfl.ch"
`,
	Run: func(cmd *cobra.Command, args []string) {
		var app models.Application
		if OptDisplayName == "" {
			printErr(errors.New("Name is required (use --displayname)"))
			return
		}

		// Configure app
		app.DisplayName = OptDisplayName
		if OptHomeURI != "" {
			app.Web = &models.WebSection{HomePageURL: OptHomeURI}
		}

		newApp, err := Client.CreateApplication(&app, clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		err = Client.WaitApplication(newApp.ID, 60, clientOptions)
		if err != nil {
			printErr(err)
			return
		}

		sp, err := Client.CreateServicePrincipal(&models.ServicePrincipal{
			AppID: newApp.AppID,
			Tags: []string{
				"HideApp",
				"WindowsAzureActiveDirectoryIntegratedApp",
			},
			ServicePrincipalType: "Application"}, clientOptions)
		if err != nil {
			printErr(err)
			return
		}
		cmd.Printf("Application created: %+v\n", sp)
	},
}

func init() {
	applicationCmd.AddCommand(applicationCreateCmd)

	// hideInCommand(applicationCreateCmd, "top', 'skip', 'skiptoken', 'select', 'search")
	applicationCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationCreateCmd.Flags().MarkHidden("batch")
		applicationCreateCmd.Flags().MarkHidden("search")
		applicationCreateCmd.Flags().MarkHidden("select")
		applicationCreateCmd.Flags().MarkHidden("skip")
		applicationCreateCmd.Flags().MarkHidden("skiptoken")
		applicationCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
