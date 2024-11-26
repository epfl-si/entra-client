package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationOIDCCmd represents the applicationOIDC command
var applicationPermissionCmd = &cobra.Command{
	Use:   "permission",
	Short: "Handles permission for applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("permission called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationPermissionCmd)
}
