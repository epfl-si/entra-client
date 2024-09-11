package cmdclaim

import (
	"encoding/json"

	rootcmd "entra-client/cmd"
	"entra-client/pkg/client/models"

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
		if rootcmd.OptDisplayName == "" && !OptDefault {
			rootcmd.PrintErrString("DisplayName is required (use --displayName, or use --default)")
			return
		}

		if rootcmd.OptPostData == "" && !OptDefault {
			rootcmd.PrintErrString("Data or default flag is required (use --data or --default)")
			return
		}

		if rootcmd.OptPostData != "" && OptDefault {
			rootcmd.PrintErrString("Data OR default flag are mutually exclusive (use --data OR --default)")
			return
		}

		var claim models.ClaimsMappingPolicy

		if OptDefault {
			claim = models.ClaimsMappingPolicy{
				Definition:            []string{"{\"ClaimsMappingPolicy\":{\"Version\":1,\"IncludeBasicClaimSet\":\"false\",\"ClaimsSchema\": [{\"Source\":\"user\",\"ID\":\"user.employeeid\",\"JwtClaimType\": \"uniqueid\"}]}}"},
				DisplayName:           "EPFL Default Claims Policy",
				IsOrganizationDefault: false,
			}
		}

		if rootcmd.OptPostData != "" {
			err := json.Unmarshal([]byte(rootcmd.OptPostData), &claim)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}
		}

		id, err := rootcmd.Client.CreateClaimsMappingPolicy(&claim, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		claim.ID = id

		cmd.Println(rootcmd.OutputJSON(claim))
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
