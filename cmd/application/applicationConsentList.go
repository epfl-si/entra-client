package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/spf13/cobra"
)

// applicationOIDCCmd represents the applicationOIDC command
var applicationConsentListCmd = &cobra.Command{
	Use:   "list",
	Short: "List an application consent",
	Long: `This command enables you to list an application consent.
	Example:
	       ./ecli application consent list --filter "clientId eq 'c060a310-19d2-4e48-bad9-8eb38f078f1a'"

		   ClientId is the application Service Principal ObjectID
		   `,
	Run: func(cmd *cobra.Command, args []string) {

		body, err := rootcmd.Client.GetApplicationConsents(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		// cmd.Printf("%s\n", rootcmd.OutputJSON(application))
		cmd.Println("body: ", body)
	},
}

func init() {
	applicationConsentCmd.AddCommand(applicationConsentListCmd)
}
