package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationOIDCCmd represents the applicationOIDC command
var applicationSecretCmd = &cobra.Command{
	Use:   "secret",
	Short: "Handles secrets for applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("secret called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationSecretCmd)
}
