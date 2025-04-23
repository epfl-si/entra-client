package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// ClaimType is the type of claim to be added (--type)
var ClaimType string

// ClaimName is the name of claim to be added (--name)
var ClaimName string

// ClaimSource is the source of claim to be added (--source)
var ClaimSource string

// ClaimPresetBasics is a flag to add a predefinite set of basic claims (--basics)
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
			cmd.PrintErr("Service Principal ID is required (use --id)\n")
			return
		}

		if !ClaimPresetBasics {

			if ClaimName == "" {
				cmd.PrintErr("Claim name is required (use --name)\n")
				return
			}
			if ClaimType == "" {
				cmd.PrintErr("Claim type is required (use --type [id/access/saml2])\n")
				return
			}

			if ClaimSource == "" {
				cmd.PrintErr("Claim source is required (use --source)\n")
				return
			}
		}

		err := rootcmd.Client.AddClaimToApplication(rootcmd.OptID, ClaimName, ClaimSource, ClaimType, ClaimPresetBasics, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr("No service principal found for this appID: " + err.Error() + "\n")
			return
		}
	},
}

func init() {
	applicationClaimCmd.AddCommand(applicationClaimAddCmd)

	applicationClaimAddCmd.Flags().StringVar(&ClaimType, "type", "", "The type of claim (id/access/saml2)")
	applicationClaimAddCmd.Flags().StringVar(&ClaimName, "name", "", "The name of claim")
	applicationClaimAddCmd.Flags().StringVar(&ClaimSource, "source", "", "The source of claim")
	applicationClaimAddCmd.Flags().BoolVar(&ClaimPresetBasics, "basics", false, "Add a predefinite set of basic claims")
}
