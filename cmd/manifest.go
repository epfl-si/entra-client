// Package cmd provides the commands for the command line application
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// manifestCmd represents the manifest command
var manifestCmd = &cobra.Command{
	Use:   "manifest",
	Short: "Manages application manifests",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("manifest called")
	},
}

func init() {
	rootCmd.AddCommand(manifestCmd)

	manifestCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		manifestCmd.Flags().MarkHidden("batch")
		manifestCmd.Flags().MarkHidden("displayname")
		manifestCmd.Flags().MarkHidden("post")
		manifestCmd.Flags().MarkHidden("search")
		manifestCmd.Flags().MarkHidden("select")
		manifestCmd.Flags().MarkHidden("skip")
		manifestCmd.Flags().MarkHidden("skiptoken")
		manifestCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
