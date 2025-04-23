package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationPermissionCmd represents the applicationPermission command
var applicationPermissionCmd = &cobra.Command{
	Use:   "permission",
	Short: "Handles permission for applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationPermission called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationPermissionCmd)
}
