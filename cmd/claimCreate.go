// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"
	"fmt"

	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

var OptDefault bool = false

// claimCreateCmd represents the claimCreate command
var claimCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a claims mapping policy",
	Long: `This command enables you to create a claims mapping policy.

	Example:
		./ecli claim create --data 
		{ "{\"ClaimsMappingPolicy\":
		 		{\"Version\":1,
				\"IncludeBasicClaimSet\":\"true\", 
				\"ClaimsSchema\": [
		 			{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, 
		 			{\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, 
		 			{\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, 
		 			{\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}
					]
		 		}
			}"
    ],
    "displayName": "AWS Claims Policy",
    "isOrganizationDefault": false
}
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("claimCreate called")
		if OptDisplayName == "" {
			panic("DisplayName is required (use --displayName)")
		}

		if OptPostData == "" && !OptDefault {
			panic("Data or default flag is required (use --data or --default)")
		}

		if OptPostData != "" && OptDefault {
			panic("Data OR default flag are mutually exclusive (use --data OR --default)")
		}

		var claim models.ClaimsMappingPolicy

		if OptDefault {
			claim = models.ClaimsMappingPolicy{
				Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"true\", \"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/Role\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/RoleSessionName\"}, {\"Value\":\"900\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/SessionDuration\"}, {\"Source\":\"user\",\"ID\":\"assignedroles\",\"SamlClaimType\": \"appRoles\"}, {\"Source\":\"user\",\"ID\":\"userprincipalname\",\"SamlClaimType\": \"https://aws.amazon.com/SAML/Attributes/nameidentifier\"}]}}"},
				DisplayName:           OptDisplayName,
				IsOrganizationDefault: false,
			}
		}

		if OptPostData != "" {
			err := json.Unmarshal([]byte(OptPostData), &claim)
			if err != nil {
				fmt.Println("Error unmarshalling JSON data")
				panic(err)
			}
		}

		id, err := Client.CreateClaimsMappingPolicy(&claim, clientOptions)
		if err != nil {
			fmt.Println("Error creating claims mapping policy")
			panic(err)
		}
		claim.ID = id

		fmt.Printf("%s\n", OutputJSON(claim))
	},
}

func init() {
	claimCmd.AddCommand(claimCreateCmd)

	claimCreateCmd.Flags().BoolVar(&OptDefault, "default", false, "Create a default claims mapping policy")
}
