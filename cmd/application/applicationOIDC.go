package cmdapplication

import (
	"github.com/spf13/cobra"
)

var OptType string

// applicationOIDCCmd represents the applicationOIDC command
var applicationOIDCCmd = &cobra.Command{
	Use:   "oidc",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationOIDC called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationOIDCCmd)
	applicationOIDCCmd.PersistentFlags().StringVar(&OptType, "type", "", "Type of OIDC application ('web', 'spa')")
}
