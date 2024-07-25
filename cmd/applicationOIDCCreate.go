// Package cmd provides the commands for the command line application
package cmd

import (
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// applicationOIDCCreateCmd represents the applicationOIDCCreate command
var applicationOIDCCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationOIDCCreate called")
		if OptAppName == "" {
			panic("Name is required (use --name)")
		}
		client, err := httpengine.New()
		if err != nil {
			panic(err)
		}
		opts := models.ClientOptions{}
		app, sp, err := client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", OptAppName, opts)
		err = client.WaitServicePrincipal(sp.ID, 60, opts)
		if err != nil {
			panic(err)
		}
		err = client.WaitApplication(app.ID, 60, opts)
		if err != nil {
			panic(err)
		}
		// Add redirect uri
		// By default, use grant type: Authorization Code Flow with PKCE.

		// Configure supported account types

		// Configure claims

	},
}

func init() {
	applicationOIDCCmd.AddCommand(applicationOIDCCreateCmd)
}
