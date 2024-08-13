// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// applicationOIDCCmd represents the applicationOIDC command
var applicationOIDCCmd = &cobra.Command{
	Use:   "oidc",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationOIDC called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationOIDCCmd)
}
