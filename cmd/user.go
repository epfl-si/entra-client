// Package cmd provides the commands for the command line application
package cmd

import (
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
	rootCmd.AddCommand(userCmd)
}
