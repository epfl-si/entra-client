package cmdapplication

import (
	"encoding/json"
	"fmt"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/epfl-si/entra-client/pkg/saml"

	"github.com/spf13/cobra"
)

// OptMetadataFile is associated with the --metadata_file flag
var OptMetadataFile string

// OptLogoutURI is associated with the --logout_uri flag
var OptLogoutURI string

// applicationSAMLCreateCmd represents the applicationSAMLCreate command
// Resources:
// - https://github.com/MicrosoftDocs/azure-docs/issues/59275
// - https://learn.microsoft.com/en-us/graph/application-saml-sso-configure-api?tabs=http%2Cpowershell-script
var applicationSAMLCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a SAML application",
	Long: `Create a SAML application
	
	Example:
	  ./ecli application saml create --displayname "AA SAML provisioning 1" --identifier "https://aasamlprovisioning1" --redirect_uri "https://aasamlprovisioning1/redirect"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var SAMLID string
		var RedirectURI []string
		var LogoutURI string
		var m *saml.EntityDescriptor
		var err error

		if rootcmd.OptDisplayName == "" {
			cmd.PrintErr("Name is required (use --displayname)\n")
			return
		}

		bootstrApp := &models.Application{
			DisplayName: rootcmd.OptDisplayName,
			Web:         &models.WebSection{},
		}

		if OptMetadataFile != "" {
			m, err = saml.GetMetadata(OptMetadataFile)
			if err != nil {
				cmd.PrintErr(fmt.Errorf("getting metadata: %w", err))
				return
			}

			if rootcmd.OptDebug {
				jsonMetadata, _ := json.Marshal(m)
				cmd.Printf("Metadata: %s\n\n", jsonMetadata)
			}

			// Create the application
			// app, sp, err := CreateApplication(OptDisplayName)
			SAMLID = m.EntityID
			if m.SPSSODescriptors != nil && m.SPSSODescriptors[0].AssertionConsumerServices != nil {
				RedirectURI = []string{m.SPSSODescriptors[0].AssertionConsumerServices[0].Location}
			}
			if m.SPSSODescriptors != nil && m.SPSSODescriptors[0].SingleLogoutServices != nil {
				LogoutURI = m.SPSSODescriptors[0].SingleLogoutServices[0].Location
			}
		}

		if OptSAMLID == "" && OptMetadataFile == "" {
			cmd.PrintErr("SAML identifier is required (use --identifier or --metadata_file)\n")
			return
		}
		if len(OptRedirectURI) == 0 && OptMetadataFile == "" {
			cmd.PrintErr("SAML redirect URI is required (use --redirect_uri or --metadata_file)\n")
			return
		}

		// Explicit flags overrides metadata file
		if OptSAMLID != "" {
			SAMLID = OptSAMLID
		}
		if len(OptRedirectURI) != 0 {
			RedirectURI = OptRedirectURI
		}
		if OptLogoutURI != "" {
			LogoutURI = OptLogoutURI
		}

		// Should be applied at the flag level (with a kind of transformer/validator)
		for i, uri := range RedirectURI {
			RedirectURI[i] = rootcmd.NormalizeURI(uri)
		}
		LogoutURI = rootcmd.NormalizeURI(LogoutURI)
		SAMLID = rootcmd.NormalizeURI(SAMLID)

		if rootcmd.OptDebug {
			cmd.Printf("EntityID: %s\n", SAMLID)
			cmd.Printf("Redirect URI: %s\n", RedirectURI)
			cmd.Printf("Logout URI: %s\n", LogoutURI)
		}

		URIList := []models.URI{}
		for i, uri := range RedirectURI {
			URIList = append(URIList, models.URI{URI: uri, Index: i})
		}
		if len(RedirectURI) == 0 {
			bootstrApp.Web.RedirectURISettings = URIList
		}

		if LogoutURI != "" {
			bootstrApp.Web.LogoutURL = LogoutURI
		}

		opts := models.ClientOptions{}

		app, sp, err := rootcmd.Client.CreatePortalApplication(bootstrApp, opts)
		if err != nil {
			rootcmd.PrintErr(fmt.Errorf("createApplication: %W", err))
			return
		}

		appPatch := &models.Application{}
		web := &models.WebSection{}
		if SAMLID != "" {
			appPatch.IdentifierUris = []string{SAMLID}
		}
		if len(RedirectURI) != 0 {
			web.RedirectURIs = RedirectURI
			// Can't be modified at the same time
			// web.RedirectURISettings = []models.URI{{URI: RedirectURI, Index: 1}}
		}
		if LogoutURI != "" {
			web.LogoutURL = LogoutURI
		}

		web.ImplicitGrantSettings = &models.Grant{EnableIDTokenIssuance: true}

		appPatch.Web = web

		err = rootcmd.Client.PatchApplication(app.ID, appPatch, opts)
		if err != nil {
			fmt.Printf("appPatch: %+v\n", appPatch)
			rootcmd.PrintErr(fmt.Errorf("patching Application %s: %w", app.ID, err))
			return
		}

		// Custom settings for Service Principal
		tags := sp.Tags
		tags = append(tags, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication")
		spName := sp.ServicePrincipalNames
		spName = append(spName, SAMLID)

		err = rootcmd.Client.PatchServicePrincipal(sp.ID, &models.ServicePrincipal{
			PreferredSingleSignOnMode: "saml",
			ReplyUrls:                 RedirectURI,
			LogoutURL:                 LogoutURI,
			Tags:                      tags,
			ServicePrincipalNames:     spName,
		}, opts)
		if err != nil {
			rootcmd.PrintErr(fmt.Errorf("patching ServicePrincipal %s: %w", sp.ID, err))
			return
		}

		// IdentifierUris has to be set later to avoid some issues with reserved domain names
		appPatch = &models.Application{}
		appPatch.IdentifierUris = []string{SAMLID}
		err = rootcmd.Client.PatchApplication(app.ID, appPatch, opts)
		if err != nil {
			rootcmd.PrintErr(fmt.Errorf("patching Application %s: %w", app.ID, err))
			return
		}

		claims := &models.ClaimsMappingPolicy{
			Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\", \"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, {\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}]}}"},
			DisplayName:           rootcmd.OptDisplayName + " Policy",
			IsOrganizationDefault: false,
		}

		claimsID, err := rootcmd.Client.CreateClaimsMappingPolicy(claims, opts)
		if err != nil {
			rootcmd.PrintErr(fmt.Errorf("creating ClaimsPolicy %s: %w", claimsID, err))
			return
		}

		err = rootcmd.Client.AssignClaimsPolicyToServicePrincipal(claimsID, sp.ID)
		if err != nil {
			rootcmd.PrintErr(fmt.Errorf("assign ClaimsPolicy %s to ServicePrincipal %s: %w", claimsID, sp.ID, err))
			return
		}

		// use := map[string]string{"signing": "Verify", "encryption": "Sign"}
		for _, desc := range m.SPSSODescriptors {
			for _, crt := range desc.KeyDescriptors {

				fmt.Print("\n\n    Adding " + crt.Use + " certificate\n")

				// Original from metadata
				err = rootcmd.Client.AddCertificateToServicePrincipal(sp.ID, crt.KeyInfo.X509Data.X509Certificates[0].Data, opts)
				if err != nil {
					rootcmd.PrintErr(fmt.Errorf("could'nt add certificate: %W", err))
					return
				}

				fmt.Print("\n\n    Adding key\n")
				err = rootcmd.Client.AddKeyToServicePrincipal(sp.ID, crt, opts)
				if err != nil {
					rootcmd.PrintErr(fmt.Errorf("could'nt add key: %W", err))
					return
				}

				// New way Create both Sign/Verify certificates
				// err = addCertificate(app.ID, sp.ID, "Verify", "AsymmetricX509Cert", crt.KeyInfo.X509Data.X509Certificates[0].Data, opts)
				// if err != nil {
				// 	PrintErr(fmt.Errorf("could'nt add SIGN certificate: %W", err))
				// 	return
				// }
				// err = addCertificate(app.ID, sp.ID, "Sign", "AsymmetricX509Cert", crt.KeyInfo.X509Data.X509Certificates[0].Data, opts)
				// if err != nil {
				// 	PrintErr(fmt.Errorf("could'nt add VERIFY certificate: %W", err))
				// 	return
				// }

			}
		}
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLCreateCmd)

	applicationSAMLCreateCmd.PersistentFlags().StringVar(&OptMetadataFile, "metadata_file", "", "The metadata file name")
	applicationSAMLCreateCmd.PersistentFlags().StringVar(&OptLogoutURI, "logout_uri", "", "The SAML logout URI")

	applicationSAMLCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationSAMLCreateCmd.Flags().MarkHidden("batch")
		applicationSAMLCreateCmd.Flags().MarkHidden("search")
		applicationSAMLCreateCmd.Flags().MarkHidden("select")
		applicationSAMLCreateCmd.Flags().MarkHidden("skip")
		applicationSAMLCreateCmd.Flags().MarkHidden("skiptoken")
		applicationSAMLCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
