package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationOIDCCmd represents the applicationOIDC command
var applicationConsentCmd = &cobra.Command{
	Use:   "consent",
	Short: "Handles consent for applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("consent called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationConsentCmd)
}
