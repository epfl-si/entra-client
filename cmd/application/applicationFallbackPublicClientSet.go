package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// ClaimType is the type of claim to be added (--type)
var True bool

// ClaimName is the name of claim to be added (--name)
var False bool

// applicationFallbackPublicClientSetCmd represents the applicationFallbackPublicClientSet command
var applicationFallbackPublicClientSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Set the FallbackPublicClient value for an application",
	Long: `Set the FallbackPublicClient value for an application

Example:

	  ./ecli application fallbackpublicclient set --appid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --true
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptAppID == "" {
			cmd.PrintErr("Application ID is required (use --appid)\n")
			return
		}

		if !True && !False {
			cmd.PrintErr("Either --true or --false value is required\n")
			return
		}

		var propertyValue bool
		if True {
			propertyValue = true
		} else if False {
			propertyValue = false
		}

		err := rootcmd.Client.SetFallbackPublicClient(rootcmd.OptAppID, propertyValue, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr("Error while setting FallbackPublicClient for appID: " + err.Error() + "\n")
			return
		}
	},
}

func init() {
	applicationFallbackPublicClientCmd.AddCommand(applicationFallbackPublicClientSetCmd)

	applicationFallbackPublicClientSetCmd.Flags().BoolVar(&True, "true", false, "true value")
	applicationFallbackPublicClientSetCmd.Flags().BoolVar(&False, "false", false, "false value")
	applicationFallbackPublicClientSetCmd.MarkFlagsMutuallyExclusive("true", "false")
}
