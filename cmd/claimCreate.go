// Package cmd provides the commands for the command line application
package cmd

import (
	"encoding/json"

	"epfl-entra/internal/models"

	"github.com/spf13/cobra"
)

// OptDefault is associated with the --default flag
var OptDefault = false

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
		if OptDisplayName == "" {
			printErrString("DisplayName is required (use --displayName)")
			return
		}

		if OptPostData == "" && !OptDefault {
			printErrString("Data or default flag is required (use --data or --default)")
			return
		}

		if OptPostData != "" && OptDefault {
			printErrString("Data OR default flag are mutually exclusive (use --data OR --default)")
			return
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
				printErr(err)
				return
			}
		}

		id, err := Client.CreateClaimsMappingPolicy(&claim, clientOptions)
		if err != nil {
			printErr(err)
			return
		}
		claim.ID = id

		cmd.Println(OutputJSON(claim))
	},
}

func init() {
	claimCmd.AddCommand(claimCreateCmd)

	claimCreateCmd.Flags().BoolVar(&OptDefault, "default", false, "Create a default claims mapping policy")

	claimCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimCreateCmd.Flags().MarkHidden("batch")
		claimCreateCmd.Flags().MarkHidden("search")
		claimCreateCmd.Flags().MarkHidden("select")
		claimCreateCmd.Flags().MarkHidden("skip")
		claimCreateCmd.Flags().MarkHidden("skiptoken")
		claimCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
