package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/spf13/cobra"
)

var OptScope []string

// applicationOIDCCmd represents the applicationOIDC command
var applicationConsentGiveCmd = &cobra.Command{
	Use:   "give",
	Short: "Give consent to an application's permissions",
	Long: `Give consent to an application's permissions

	By default with consent to Microsoft Graph API permissions.

	Example:
		ecli application consent give --id <service_principal_object_id> --scope openid --scope email --scope profile --scope offline_access --scope User.Read
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: add a ressource_id flag to consent to non Microsoft Graph API permisisons
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ObjectID is required (use --id)")
			return
		}
		if len(OptScope) == 0 {
			rootcmd.PrintErrString("Scopes are required (use --scope)")
			return
		}

		err := rootcmd.Client.GiveConsentToApplication(rootcmd.OptID, OptScope, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		cmd.Println("Consent given to application")
	},
}

func init() {
	applicationConsentCmd.AddCommand(applicationConsentGiveCmd)
	applicationConsentGiveCmd.PersistentFlags().StringArrayVar(&OptScope, "scope", []string{}, "Scopes")
}