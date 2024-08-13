// Package cmd provides the commands for the command line application
package cmd

import (
	"github.com/spf13/cobra"
)

// apptemplateCmd represents the apptemplate command
var apptemplateCmd = &cobra.Command{
	Use:   "apptemplate",
	Short: "Manage application templates",
	Long: `This command enables you to
	* Get

	application template(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("apptemplate called")
	},
}

func init() {
	rootCmd.AddCommand(apptemplateCmd)
}
