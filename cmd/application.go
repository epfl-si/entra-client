// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OptAppName is associated with the --name flag
var OptAppName string

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
	applicationCmd.PersistentFlags().StringVar(&OptAppName, "name", "", "Name of the application")
}
