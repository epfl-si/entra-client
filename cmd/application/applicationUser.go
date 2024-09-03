package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationSAMLUser.goCmd represents the applicationSAMLUser.go command
var applicationUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Handle users for an application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationUser called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationUserCmd)
}
