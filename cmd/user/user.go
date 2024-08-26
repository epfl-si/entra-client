// Package usercmd is used for user commands
package usercmd

import (
	rootCmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// userCmd represents the user command
var userCmd = &cobra.Command{
	Use:   "user",
	Short: "Manage users",
	Long: `This command enables you to
	* Create
	* Get
	* Modify
	* Delete

	user(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("user called")
	},
}

func init() {
	rootCmd.RootCmd.AddCommand(userCmd)
}
