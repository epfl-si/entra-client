package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationList.goCmd represents the applicationList.go command
var applicationListCmd = &cobra.Command{
	Use:   "list",
	Short: "List applications",
	Run: func(cmd *cobra.Command, args []string) {
		applications, _, err := rootcmd.Client.GetApplications(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, application := range applications {
			cmd.Printf("%s\n", rootcmd.OutputJSON(application))
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationListCmd)
}
