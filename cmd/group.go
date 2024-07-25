// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// groupCmd represents the group command
var groupCmd = &cobra.Command{
	Use:   "group",
	Short: "Manage groups",
	Long: `This command enables you to
	* Create
	* Get
	* Modify
	* Delete

	group(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("group called")
	},
}

func init() {
	rootCmd.AddCommand(groupCmd)
}
