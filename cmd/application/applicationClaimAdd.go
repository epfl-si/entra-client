package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

var ClaimLocation string
var ClaimName string
var ClaimSource string
var ClaimPresetBasics bool

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

		if !ClaimPresetBasics {

			if ClaimName == "" {
				rootcmd.PrintErrString("Claim name is required (use --name)")
				return
			}
			if ClaimLocation == "" {
				rootcmd.PrintErrString("Claim type is required (use --type [id/access/saml2])")
				return
			}

			if ClaimSource == "" {
				rootcmd.PrintErrString("Claim source is required (use --source)")
				return
			}
		}

		err := rootcmd.Client.AddClaimToApplication(rootcmd.OptID, ClaimName, ClaimSource, ClaimLocation, ClaimPresetBasics, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationClaimCmd.AddCommand(applicationClaimAddCmd)

	applicationClaimAddCmd.Flags().StringVar(&ClaimLocation, "location", "", "The type of claim (id/access/saml2)")
	applicationClaimAddCmd.Flags().StringVar(&ClaimName, "name", "", "The name of claim")
	applicationClaimAddCmd.Flags().StringVar(&ClaimSource, "source", "", "The source of claim")
	applicationClaimAddCmd.Flags().BoolVar(&ClaimPresetBasics, "basics", false, "Add a predefinite set of basic claims")
}
