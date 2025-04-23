package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			cmd.PrintErr("ID is required (use --id)\n")
			return
		}
		application, err := rootcmd.Client.GetApplication(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(application))
	},
}

func init() {
	applicationCmd.AddCommand(applicationGetCmd)
}
