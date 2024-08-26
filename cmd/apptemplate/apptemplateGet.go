package cmdapptemplate

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// apptemplateGetCmd represents the apptemplateGet command
var apptemplateGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get an application template by ID",
	Run: func(cmd *cobra.Command, args []string) {
		application, err := rootcmd.Client.GetApplicationTemplate(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("ApplicationTemplate: %s\n", rootcmd.OutputJSON(application))
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateGetCmd)
}
