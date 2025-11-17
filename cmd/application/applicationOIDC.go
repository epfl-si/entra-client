package cmdapplication

import (
	"github.com/spf13/cobra"
)

// OptType is associated with the --type flag
var OptType string

// applicationOIDCCmd represents the applicationOIDC command
var applicationOIDCCmd = &cobra.Command{
	Use:   "oidc",
	Short: "Handles OIDC applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationOIDC called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationOIDCCmd)
	applicationOIDCCmd.PersistentFlags().StringVar(&OptType, "type", "", "Type of OIDC application ('web', 'spa')")
}
