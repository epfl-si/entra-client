package cmdapplication

import (
	"github.com/spf13/cobra"
)

// applicationClaimCmd represents the applicationClaim command
var applicationClaimCmd = &cobra.Command{
	Use: "claim",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationClaim called")
	},
}

func init() {
	applicationCmd.AddCommand(applicationClaimCmd)
}
