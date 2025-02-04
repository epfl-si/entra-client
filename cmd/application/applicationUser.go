package cmdapplication

import (
	"github.com/spf13/cobra"
)

// OptUserID is associated with the --userID flag
var OptUserID string

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
	applicationUserCmd.PersistentFlags().StringVar(&OptUserID, "userid", "", "ID of and user/group to be (un)associated to the application")
}
