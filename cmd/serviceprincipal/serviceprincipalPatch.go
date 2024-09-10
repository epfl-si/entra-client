package serviceprincipalcmd

import (
	"encoding/json"
	rootcmd "epfl-entra/cmd"
	"epfl-entra/pkg/entra-client/models"

	"github.com/spf13/cobra"
)

// serviceprincipalPatchCmd represents the serviceprincipalPatch command
var serviceprincipalPatchCmd = &cobra.Command{
	Use:   "patch",
	Short: "Patch a ServicePrincipal",
	Run: func(cmd *cobra.Command, args []string) {
		var app models.ServicePrincipal
		err := json.Unmarshal([]byte(rootcmd.OptPostData), &app)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		err = rootcmd.Client.PatchServicePrincipal(rootcmd.OptID, &app, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(serviceprincipalPatchCmd)

	serviceprincipalPatchCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		serviceprincipalPatchCmd.Flags().MarkHidden("batch")
		serviceprincipalPatchCmd.Flags().MarkHidden("search")
		serviceprincipalPatchCmd.Flags().MarkHidden("select")
		serviceprincipalPatchCmd.Flags().MarkHidden("skip")
		serviceprincipalPatchCmd.Flags().MarkHidden("skiptoken")
		serviceprincipalPatchCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
