package cmdapplication

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// applicationGetCmd represents the applicationGet command
var applicationGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application by ID",
	Run: func(cmd *cobra.Command, args []string) {
		application, err := rootcmd.Client.GetApplication(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(application))
	},
}

func init() {
	applicationCmd.AddCommand(applicationGetCmd)
}
