package cmdapplication

// @task write a test file applicationOIDCCreate_test.go that test the command applicationOIDCCreate @run

import (
	rootcmd "entra-client/cmd"
	"entra-client/pkg/client/models"

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
		if len(OptRedirectURI) == 0 {
			rootcmd.PrintErrString("Callback URL is required (use --redirect_uri)")
			return
		}

		URIList := []models.URI{}
		for i, uri := range OptRedirectURI {
			URIList = append(URIList, models.URI{URI: uri, Index: i})
		}
		bootstrApp := &models.Application{
			DisplayName: rootcmd.OptDisplayName,
			Web: &models.WebSection{
				RedirectURISettings: URIList,
			},
		}

		opts := models.ClientOptions{}

		app, sp, err := rootcmd.CreateApplication(bootstrApp, opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		secret, err := rootcmd.Client.AddPasswordToApplication(app.ID, rootcmd.OptDisplayName+" secret", opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		// cmd.Printf("Application ID: %s\n\n\n", rootcmd.OutputJSON(app))
		cmd.Printf("Application ID: %s\n\n\n", rootcmd.OutputJSON(sp))
		cmd.Printf("Client ID: %s\n", app.AppID)
		cmd.Printf("Client secret: %s\n\n", *secret.SecretText)

		appPatch := &models.Application{}
		appPatch.Web = &models.WebSection{
			ImplicitGrantSettings: &models.Grant{
				EnableIDTokenIssuance:     true,
				EnableAccessTokenIssuance: true,
			},
		}
		// Default to SPA
		if OptType == "" {
			OptType = "spa"
		}
		switch OptType {
		case "web":
			appPatch.Web.RedirectURIs = OptRedirectURI
			appPatch.Web.RedirectURISettings = URIList
		case "spa":
			appPatch.Spa = &models.SpaApplication{
				RedirectURIs: OptRedirectURI,
			}
		}

		version := 2
		t := true
		// appPatch.Tags = []string{"HideApp"}
		appPatch.API = &models.ApiApplication{
			AcceptMappedClaims:          &t,
			RequestedAccessTokenVersion: &version,
		}

		// Causes error:
		// appPatch.AllowPublicClient = true
		appPatch.IsFallbackPublicClient = &t // For PKCE

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
		// spPatch.Tags = []string{"HideApp"}
		spPatch.AppRoleAssignmentRequired = true

		// Causes error:

		err = rootcmd.Client.PatchServicePrincipal(sp.ID, spPatch, opts)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, groupID := range []string{
			"AAD_All Hosts Users",
			"AAD_All Outside EPFL Users",
			"AAD_All Staff Users",
			"AAD_All Student Users",
		} {

			err = rootcmd.Client.AddGroupToServicePrincipal(sp.ID, groupID, opts)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}
		}

		// Works but can't be edited by portal
		// err = rootcmd.Client.AssignClaimsPolicyToServicePrincipal("b0a98d4a-221f-4d76-b6fb-7f6f0089175f", sp.ID)
		// if err != nil {
		// 	rootcmd.PrintErr(fmt.Errorf("Assign ClaimsPolicy %s to ServicePrincipal %s: %w", "b0a98d4a-221f-4d76-b6fb-7f6f0089175", sp.ID, err))
		// 	return
		// }

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
