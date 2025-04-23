package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/spf13/cobra"
)

// OptSecretDuration is associated with the --duration flag
var OptSecretDuration string

// OptSecretName is associated with the --name flag
var OptSecretName string

// applicationOIDCCmd represents the applicationOIDC command
var applicationSecretAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a secret to an OIDC application",
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			cmd.PrintErr("Service Principal ID is required (use --id)\n")
			return
		}

		// if OptSecretDuration == "" {
		// 	rootcmd.PrintErrString("Duration is required (use --duration)")
		// 	return
		// }

		if OptSecretName == "" {
			cmd.PrintErr("Name is required (use --name)\n")
			return
		}

		// password, err := rootcmd.Client.AddPasswordToApplication(rootcmd.OptID, rootcmd.OptDisplayName, OptSecretDuration, rootcmd.ClientOptions)
		password, err := rootcmd.Client.AddPasswordToApplication(rootcmd.OptID, OptSecretName, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("Client secret: %s\n\n", *password.SecretText)
	},
}

func init() {
	applicationSecretCmd.AddCommand(applicationSecretAddCmd)
	applicationSecretAddCmd.PersistentFlags().StringVar(&OptSecretDuration, "duration", "", "Duration of the secret('1m', '1y', '2y')")
	applicationSecretAddCmd.PersistentFlags().StringVar(&OptSecretName, "name", "", "Name of the secret")
}
