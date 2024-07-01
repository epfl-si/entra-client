// Package cmd provides the command line application for the application
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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// groupCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// groupCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
