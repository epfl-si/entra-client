package cmdapplication

import (
	"github.com/spf13/cobra"
)

// OptSAMLID is associated with the --id flag
var OptSAMLID string

// applicationSAMLCmd represents the applicationSAML command
var applicationSAMLCmd = &cobra.Command{
	Use:   "saml",
	Short: "Handle SAML applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAML called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationSAMLCmd)

	applicationSAMLCmd.PersistentFlags().StringVar(&OptSAMLID, "identifier", "", "SP SAML Identifier")
}
