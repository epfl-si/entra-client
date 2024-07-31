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
	Long: `Create a SAML application
	
	Example:
	  ./ecli application saml create --displayname "AA SAML provisioning 1" --identifier "https://aasamlprovisioning1" --redirect_uri "https://aasamlprovisioning1/redirect"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptDisplayName == "" {
			panic("Name is required (use --displayname)")
		}
		if OptSAMLID == "" {
			panic("SAML identifier is required (use --identifier)")
		}
		if OptRedirectURI == "" {
			panic("SAML redirect URI is required (use --redirect_uri)")
		}
		client, err := httpengine.New()
		if err != nil {
			panic(err)
		}

		opts := models.ClientOptions{}
		app, sp, err := client.InstantiateApplicationTemplate("229946b9-a9fb-45b8-9531-efa47453ac9e", OptDisplayName, opts)
		if err != nil {
			panic(err)
		}

		err = client.WaitServicePrincipal(sp.ID, 60, opts)
		if err != nil {
			panic(err)
		}

		// Custom settings
		err = client.PatchServicePrincipal(sp.ID, &models.ServicePrincipal{
			PreferredSingleSignOnMode: "saml",
			Homepage:                  "https://www.epfl.ch",
		}, opts)
		err = client.WaitApplication(app.ID, 60, opts)
		if err != nil {
			panic(err)
		}
		if OptSAMLID != "" || OptRedirectURI != "" {
			appPatch := &models.Application{}
			if OptSAMLID != "" {
				appPatch.IdentifierUris = []string{OptSAMLID}
			}
			if OptRedirectURI != "" {
				appPatch.Web = &models.WebSection{RedirectURIs: []string{OptRedirectURI}}
			}
			err = client.PatchApplication(app.ID, appPatch, opts)
		}

		// Get template claims mapping policies
		templateClaims, _, err := Client.GetClaimsMappingPoliciesForServicePrincipal(sp.ID, opts)
		fmt.Printf("Template claims: %+v\n", templateClaims)
		if err != nil {
			panic(err)
		}

		// Unassign all claims mapping policies (from template)
		for _, claim := range templateClaims {
			fmt.Println("Unassigning claim", claim.ID)
			err = Client.UnassignClaimsPolicyFromServicePrincipal(claim.ID, sp.ID, opts)
			if err != nil {
				panic(err)
			}
		}

		claims := &models.ClaimsMappingPolicy{
			Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\", \"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, {\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}]}}"},
			DisplayName:           OptDisplayName + " Policy",
			IsOrganizationDefault: false,
		}

		claimsID, err := client.CreateClaimsMappingPolicy(claims, opts)
		if err != nil {
			panic(err)
		}

		err = client.AssignClaimsPolicyToServicePrincipal(claimsID, sp.ID)
		if err != nil {
			panic(err)
		}
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLCreateCmd)
}
