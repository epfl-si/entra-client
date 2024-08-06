// Package cmd provides the commands for the command line application
package cmd

import (
	"epfl-entra/internal/models"
	"fmt"

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
			panic("Name is required (use --displayname)")
		}

		// Configure app
		app.DisplayName = OptDisplayName
		if OptHomeURI != "" {
			app.Web = &models.WebSection{HomePageURL: OptHomeURI}
		}

		newApp, err := Client.CreateApplication(&app, clientOptions)
		if err != nil {
			panic(err)
		}

		err = Client.WaitApplication(newApp.ID, 60, clientOptions)
		if err != nil {
			panic(err)
		}

		sp, err := Client.CreateServicePrincipal(&models.ServicePrincipal{
			AppID: newApp.AppID,
			Tags: []string{
				"HideApp",
				"WindowsAzureActiveDirectoryIntegratedApp",
			},
			ServicePrincipalType: "Application"}, clientOptions)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Application created: %+v\n", sp)
	},
}

func init() {
	applicationCmd.AddCommand(applicationCreateCmd)
}
