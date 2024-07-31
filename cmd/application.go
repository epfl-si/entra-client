// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OptSAMLRedirectURI is associated with the --redirect_uri flag
var OptRedirectURI string

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
	applicationCmd.PersistentFlags().StringVar(&OptRedirectURI, "redirect_uri", "", "Redirect URI")
}
