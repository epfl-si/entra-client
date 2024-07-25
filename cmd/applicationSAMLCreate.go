// Package cmd provides the commands for the command line application
package cmd

import (
	httpengine "epfl-entra/internal/client/http_client"
	"epfl-entra/internal/models"
	"fmt"

	"github.com/spf13/cobra"
)

// applicationSAMLCreateCmd represents the applicationSAMLCreate command
var applicationSAMLCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a SAML application",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationSAMLCreate called")
		if OptAppName == "" {
			panic("Name is required (use --name)")
		}
		if OptSAMLID == "" {
			panic("SAML identifier is required (use --identifier)")
		}
		if OptSAMLRedirectURI == "" {
			panic("SAML redirect URI is required (use --redirect_uri)")
		}
		client, err := httpengine.New()
		if err != nil {
			panic(err)
		}

		opts := models.ClientOptions{}
		app, sp, err := client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", OptAppName, opts)
		if err != nil {
			panic(err)
		}

		err = client.WaitServicePrincipal(sp.ID, 60, opts)
		if err != nil {
			panic(err)
		}
		err = client.PatchServicePrincipal(sp.ID, &models.ServicePrincipal{PreferredSingleSignOnMode: "saml"}, opts)
		err = client.WaitApplication(app.ID, 60, opts)
		if err != nil {
			panic(err)
		}
		if OptSAMLID != "" || OptSAMLRedirectURI != "" {
			appPatch := &models.Application{}
			if OptSAMLID != "" {
				appPatch.IdentifierUris = []string{OptSAMLID}
			}
			if OptSAMLRedirectURI != "" {
				appPatch.Web = &models.WebSection{RedirectURIs: []string{OptSAMLRedirectURI}}
			}
			err = client.PatchApplication(app.ID, appPatch, opts)
		}

		claims := &models.ClaimsMappingPolicy{
			Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"true\", \"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, {\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}]}}"},
			DisplayName:           OptDisplayName,
			IsOrganizationDefault: false,
		}

		claimsID, err := client.CreateClaimsMappingPolicy(claims, opts)

		err = client.AssociateClaimsPolicyToServicePrincipal(claimsID, sp.ID)
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLCreateCmd)
}
