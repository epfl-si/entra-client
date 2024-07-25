// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// OptSAMLID is associated with the --id flag
var OptSAMLID string

// OptSAMLRedirectURI is associated with the --redirect_uri flag
var OptSAMLRedirectURI string

// applicationSAMLCmd represents the applicationSAML command
var applicationSAMLCmd = &cobra.Command{
	Use:   "saml",
	Short: "Handle SAML applications",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("applicationSAML called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationSAMLCmd)

	applicationSAMLCmd.PersistentFlags().StringVar(&OptSAMLID, "identifier", "", "SP SAML Identifier")
	applicationSAMLCmd.PersistentFlags().StringVar(&OptSAMLRedirectURI, "redirect_uri", "", "Redirect URI")
}
