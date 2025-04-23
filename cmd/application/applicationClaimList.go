package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationClaimListCmd represents the applicationClaimList command
var applicationClaimListCmd = &cobra.Command{
	Use:   "list",
	Short: "List claims mapping for an application",
	Long: `List claims mapping for an application

Example:

	  ./ecli application claim list --id 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			cmd.PrintErr("Service Principal ID is required (use --id)\n")
			return
		}

		application, err := rootcmd.Client.GetApplication(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(application.OptionalClaims))

	},
}

func init() {
	applicationClaimCmd.AddCommand(applicationClaimListCmd)
}
