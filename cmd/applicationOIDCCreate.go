// Package cmd provides the commands for the command line application
package cmd

import (
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// applicationOIDCCreateCmd represents the applicationOIDCCreate command
//
// Ref: https://learn.microsoft.com/en-us/entra/identity-platform/v2-protocols-oidc
var applicationOIDCCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an OIDC application",
	Long: `Create an OIDC application

Usage:
  ./ecli application oidc create --displayname "<Application name>" --redirect_uri "<Redirect URI>"

Example:
  ./ecli application oidc create --displayname "AA OIDC provisioning 1" --redirect_uri "https://aaoidcprovisioning1/redirect"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptDisplayName == "" {
			panic("Name is required (use --displayname)")
		}
		if OptRedirectURI == "" {
			panic("Callback URL is required (use --redirect_uri)")
		}
		client, err := httpengine.New()
		if err != nil {
			panic(err)
		}
		opts := models.ClientOptions{}

		// TODO: What is the utility of the app returned as first value?
		_, sp, err := client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", OptDisplayName, opts)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Application SP: %+v\n", sp)

		err = client.WaitServicePrincipal(sp.ID, 60, opts)
		if err != nil {
			panic(err)
		}

		// Customize application
		spPatch := &models.ServicePrincipal{}
		sp.Homepage = "https://www.epfl.ch"
		// spPatch.ReplyUrls = []string{OptRedirectURI}
		spPatch.Tags = []string{"WindowsAzureActiveDirectoryIntegratedApp"}

		err = client.PatchServicePrincipal(sp.ID, spPatch, opts)
		if err != nil {
			panic(err)
		}
		// By default, use grant type: Authorization Code Flow with PKCE.

		// Configure supported account types

		// Configure claims

	},
}

func init() {
	applicationOIDCCmd.AddCommand(applicationOIDCCreateCmd)
}
