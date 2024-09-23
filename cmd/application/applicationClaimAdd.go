package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

var ClaimType string
var ClaimName string
var ClaimSource string

// applicationClaimAddCmd represents the applicationClaimAdd command
var applicationClaimAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add claims mapping for an application",
	Long: `Add claims mapping for an application

Example:

	  ./ecli application claim add --id 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --type "idtoken" --name "uniqueid" --source "user.employeeid"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ID is required (use --id)")
			return
		}
		if ClaimName == "" {
			rootcmd.PrintErrString("Claim name is required (use --name)")
			return
		}
		if ClaimType == "" {
			rootcmd.PrintErrString("Claim type is required (use --type [id/access/saml2])")
			return
		}

		application, err := rootcmd.Client.GetApplication(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		switch ClaimType {
		case "id":
			claims := application.OptionalClaims.IDToken

			claims = append(claims, models.OptionalClaim{
				Name:   ClaimName,
				Source: ClaimSource,
			})
			application.OptionalClaims.IDToken = claims

		case "access":
			claims := application.OptionalClaims.AccessToken

			claims = append(claims, models.OptionalClaim{
				Name: ClaimName,

				Source: ClaimSource,
			})
			application.OptionalClaims.AccessToken = claims

		case "saml2":
			claims := application.OptionalClaims.SAML2Token

			claims = append(claims, models.OptionalClaim{
				Name:   ClaimName,
				Source: ClaimSource,
			})
			application.OptionalClaims.SAML2Token = claims
		}

		err = rootcmd.Client.PatchApplication(application.ID, application, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(application.OptionalClaims))

	},
}

func init() {
	applicationClaimCmd.AddCommand(applicationClaimAddCmd)

	applicationClaimAddCmd.Flags().StringVar(&ClaimType, "type", "", "The type of claim (id/access/saml2)")
	applicationClaimAddCmd.Flags().StringVar(&ClaimName, "name", "", "The name of claim")
	applicationClaimAddCmd.Flags().StringVar(&ClaimSource, "source", "", "The source of claim")
}
