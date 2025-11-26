package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationFallbackPublicClientCmd represents the applicationFallbackPublicClient command
var applicationFallbackPublicClientCmd = &cobra.Command{
	Use:   "fallbackpublicclient",
	Short: "Handles isFallbackPublicClient boolean attribute for applications",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationFallbackPublicClient called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationFallbackPublicClientCmd)
}
