package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationSAMLClaimGetCmd represents the applicationSAMLClaimGet command
var applicationSAMLClaimGetCmd = &cobra.Command{
	Use:   "applicationSAMLClaimGet",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLClaimGet called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(applicationSAMLClaimGetCmd)
}
