package cmdapptemplate

import (
	rootcmd "entra-client/cmd"

	"github.com/spf13/cobra"
)

// apptemplateListCmd represents the apptemplateList command
var apptemplateListCmd = &cobra.Command{
	Use:   "list",
	Short: "List apptemplate templates",
	Long: `List application templates
	
	Example:
	  ./ecli apptemplate list --select id,displayname
	`,
	Run: func(cmd *cobra.Command, args []string) {
		apptemplates, _, err := rootcmd.Client.GetApplicationTemplates(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, apptemplate := range apptemplates {
			cmd.Printf("%s\n", rootcmd.OutputJSON(apptemplate))
		}
	},
}

func init() {
	apptemplateCmd.AddCommand(apptemplateListCmd)
}
