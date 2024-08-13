// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"epfl-entra/internal/models"
	"epfl-entra/pkg/saml"

	"github.com/spf13/cobra"
)

var OptMetadataFile string
var OptLogoutURI string

// applicationSAMLCreateCmd represents the applicationSAMLCreate command
var applicationSAMLCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a SAML application",
	Long: `Create a SAML application
	
	Example:
	  ./ecli application saml create --displayname "AA SAML provisioning 1" --identifier "https://aasamlprovisioning1" --redirect_uri "https://aasamlprovisioning1/redirect"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		var SAMLID string
		var RedirectURI string
		var LogoutURI string
		var m *saml.EntityDescriptor
		var err error

		if OptDisplayName == "" {
			panic("Name is required (use --displayname)")
		}

		if OptMetadataFile != "" {
			m, err = saml.GetMetadata(OptMetadataFile)
			if err != nil {
				panic(err)
			}

			if OptDebug {
				jsonMetadata, _ := json.Marshal(m)
				cmd.Printf("Metadata: %s\n\n", jsonMetadata)
			}

			// Create the application
			// app, sp, err := CreateApplication(OptDisplayName)
			SAMLID = m.EntityID
			if m.SPSSODescriptors != nil && m.SPSSODescriptors[0].AssertionConsumerServices != nil {
				RedirectURI = m.SPSSODescriptors[0].AssertionConsumerServices[0].Location
			}
			if m.SPSSODescriptors != nil && m.SPSSODescriptors[0].SingleLogoutServices != nil {
				LogoutURI = m.SPSSODescriptors[0].SingleLogoutServices[0].Location
			}
		}

		if OptSAMLID == "" && OptMetadataFile == "" {
			panic("SAML identifier is required (use --identifier or --metadata_file)")
		}
		if OptRedirectURI == "" && OptMetadataFile == "" {
			panic("SAML redirect URI is required (use --redirect_uri or --metadata_file)")
		}

		// Explicit flags overrides metadata file
		if OptSAMLID != "" {
			SAMLID = OptSAMLID
		}
		if OptRedirectURI != "" {
			RedirectURI = OptRedirectURI
		}
		if OptLogoutURI != "" {
			LogoutURI = OptLogoutURI
		}

		// Should be applied at the flag level (with a kind of transformer/validator)
		RedirectURI = NormalizeURI(RedirectURI)
		LogoutURI = NormalizeURI(LogoutURI)
		SAMLID = NormalizeURI(SAMLID)

		if OptDebug {
			cmd.Printf("EntityID: %s\n", SAMLID)
			cmd.Printf("Redirect URI: %s\n", RedirectURI)
			cmd.Printf("Logout URI: %s\n", LogoutURI)
		}

		opts := models.ClientOptions{}
		/*
			// Create the applicatio from template
			app, sp, err := Client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", OptDisplayName, opts)
			if err != nil {
				panic(err)
			}

			err = Client.WaitServicePrincipal(sp.ID, 60, opts)
			if err != nil {
				panic(err)
			}
		*/

		app, sp, err := createApplication(OptDisplayName, opts)

		appPatch := &models.Application{}
		web := &models.WebSection{}
		if SAMLID != "" {
			appPatch.IdentifierUris = []string{SAMLID}
		}
		if RedirectURI != "" {
			web.RedirectURIs = []string{RedirectURI}
			// Can't be modified at the same time
			// web.RedirectURISettings = []models.URI{{URI: RedirectURI, Index: 1}}
		}
		if LogoutURI != "" {
			web.LogoutURL = LogoutURI
		}

		web.ImplicitGrantSettings = &models.Grant{EnableIDTokenIssuance: true}

		appPatch.Web = web

		err = Client.PatchApplication(app.ID, appPatch, opts)
		if err != nil {
			panic(err)
		}

		// Custom settings for Service Principal
		tags := sp.Tags
		tags = append(tags, "WindowsAzureActiveDirectoryCustomSingleSignOnApplication")
		spName := sp.ServicePrincipalNames
		spName = append(spName, SAMLID)

		err = Client.PatchServicePrincipal(sp.ID, &models.ServicePrincipal{
			PreferredSingleSignOnMode: "saml",
			ReplyUrls:                 []string{RedirectURI},
			LogoutURL:                 LogoutURI,
			Tags:                      tags,
			ServicePrincipalNames:     spName,
		}, opts)
		if err != nil {
			panic(err)
		}

		// IdentifierUris has to be set later to avoid some issues with reserved domain names
		appPatch = &models.Application{}
		appPatch.IdentifierUris = []string{SAMLID}
		err = Client.PatchApplication(app.ID, appPatch, opts)
		if err != nil {
			panic(err)
		}

		claims := &models.ClaimsMappingPolicy{
			Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\", \"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, {\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}]}}"},
			DisplayName:           OptDisplayName + " Policy",
			IsOrganizationDefault: false,
		}

		claimsID, err := Client.CreateClaimsMappingPolicy(claims, opts)
		if err != nil {
			panic(err)
		}

		err = Client.AssignClaimsPolicyToServicePrincipal(claimsID, sp.ID)
		if err != nil {
			panic(err)
		}

		for _, crt := range m.SPSSODescriptors[0].KeyDescriptors {
			use := map[string]string{"signing": "Verify", "encryption": "Encrypt"}
			err = addCertificate(sp.ID, use[crt.Use], "AsymmetricX509Cert", crt.KeyInfo.X509Data.X509Certificates[0].Data, opts)
			if err != nil {
				panic(err)
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
