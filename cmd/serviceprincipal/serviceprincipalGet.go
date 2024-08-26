package serviceprincipalcmd

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// serviceprincipalGetCmd represents the serviceprincipalGet command
var serviceprincipalGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a service principal by ID",
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ID is required (use --id)")
			return
		}
		sp, err := rootcmd.Client.GetServicePrincipal(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println(rootcmd.OutputJSON(sp))
	},
}

func init() {
	serviceprincipalCmd.AddCommand(serviceprincipalGetCmd)

	serviceprincipalGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalGetCmd.Flags().MarkHidden("batch")
		serviceprincipalGetCmd.Flags().MarkHidden("search")
		serviceprincipalGetCmd.Flags().MarkHidden("select")
		serviceprincipalGetCmd.Flags().MarkHidden("skip")
		serviceprincipalGetCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
