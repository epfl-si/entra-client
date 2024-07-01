// Package cmd provides the command line application for the application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage applications",
	Long: `This command enables you to
	* Create
	* Get
	* Modify
	* Delete

	application(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("application called")
	},
}

func init() {
	rootCmd.AddCommand(applicationCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// applicationCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// applicationCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
