package cmdapplication

// @task write a test file applicationOIDCCreate_test.go that test the command applicationOIDCCreate @run

import (
	rootcmd "epfl-entra/cmd"
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"

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
		if rootcmd.OptDisplayName == "" {
			rootcmd.PrintErrString("Name is required (use --displayname)")
			return
		}
		if OptRedirectURI == "" {
			rootcmd.PrintErrString("Callback URL is required (use --redirect_uri)")
			return
		}
		client, err := httpengine.New()
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		bootstrApp := &models.Application{
			DisplayName: rootcmd.OptDisplayName,
			Web: &models.WebSection{
				RedirectURISettings: []models.URI{{URI: OptRedirectURI, Index: 1}},
			},
		}

		opts := models.ClientOptions{}

		app, sp, err := rootcmd.CreateApplication(bootstrApp, opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		secret, err := client.AddPasswordToApplication(app.ID, rootcmd.OptDisplayName+" secret", opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("Application ID: %s\n\n\n", rootcmd.OutputJSON(app))
		cmd.Printf("Client ID: %s\n", app.AppID)
		cmd.Printf("Client secret: %s\n\n", *secret.SecretText)

		appPatch := &models.Application{}
		web := &models.WebSection{}
		web.RedirectURIs = []string{OptRedirectURI}

		web.ImplicitGrantSettings = &models.Grant{EnableIDTokenIssuance: true}

		appPatch.Web = web

		err = rootcmd.Client.PatchApplication(app.ID, appPatch, opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		// By default, use grant type: Authorization Code Flow with PKCE.

		// Configure supported account types

		// Configure claims

		// Customize application
		spPatch := &models.ServicePrincipal{}
		sp.Homepage = "https://www.epfl.ch"
		// spPatch.ReplyUrls = []string{OptRedirectURI}
		spPatch.Tags = []string{"WindowsAzureActiveDirectoryIntegratedApp"}

		err = client.PatchServicePrincipal(sp.ID, spPatch, opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationOIDCCmd.AddCommand(applicationOIDCCreateCmd)

	applicationOIDCCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationOIDCCreateCmd.Flags().MarkHidden("batch")
		applicationOIDCCreateCmd.Flags().MarkHidden("id")
		applicationOIDCCreateCmd.Flags().MarkHidden("post")
		applicationOIDCCreateCmd.Flags().MarkHidden("search")
		applicationOIDCCreateCmd.Flags().MarkHidden("select")
		applicationOIDCCreateCmd.Flags().MarkHidden("skip")
		applicationOIDCCreateCmd.Flags().MarkHidden("skiptoken")
		applicationOIDCCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
