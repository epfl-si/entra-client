package cmdapplication

// @task write a test file applicationOIDCCreate_test.go that test the command applicationOIDCCreate @run

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// OptTypeWeb is true if the application is of type web
var OptTypeWeb = false

// OptTypeSpa is true if the application is of type spa (default)
var OptTypeSpa = false

// applicationOIDCCreateCmd represents the applicationOIDCCreate command
//
// Ref: https://learn.microsoft.com/en-us/entra/identity-platform/v2-protocols-oidc
var applicationOIDCCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an OIDC application",
	Long: `Create an OIDC application

Usage:
  ./ecli application oidc create --displayname "<Application name>" --redirect_uri "<Redirect URI> --spa"

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

		bootstrApp := &models.Application{
			DisplayName: rootcmd.OptDisplayName,
		}

		if OptTypeSpa == true {
			bootstrApp.Spa = &models.SpaApplication{
				RedirectURIs: OptRedirectURI,
			}
		} else if OptTypeWeb == true {
			URIList := []models.URI{}
			for i, uri := range OptRedirectURI {
				URIList = append(URIList, models.URI{URI: uri, Index: i})
			}
			bootstrApp.Web = &models.WebSection{
				RedirectURIs:        OptRedirectURI,
				RedirectURISettings: URIList,
			}
		} else {
			panic("Invalid application type")
		}

		app, _, secret, err := rootcmd.Client.CreateOIDCApplication(bootstrApp)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		// cmd.Printf("Application ID: %s\n\n\n", rootcmd.OutputJSON(app))
		cmd.Printf("Application ID: %s\n\n\n", app.AppID)
		cmd.Printf("Client secret: %s\n\n", secret)
	},
}

func init() {
	applicationOIDCCmd.AddCommand(applicationOIDCCreateCmd)

	applicationOIDCCreateCmd.Flags().BoolVar(&OptTypeWeb, "web", false, "The application type is web")
	applicationOIDCCreateCmd.Flags().BoolVar(&OptTypeSpa, "spa", true, "The application type is spa (default)")

	applicationOIDCCreateCmd.MarkFlagsMutuallyExclusive("web", "spa")
	applicationOIDCCreateCmd.MarkFlagsRequiredTogether("web", "spa")

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
