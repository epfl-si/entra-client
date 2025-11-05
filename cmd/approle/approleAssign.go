package cmdapprole

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// approleAssignCmd represents the approleAssign command
var approleAssignCmd = &cobra.Command{
	Use:   "assign",
	Short: "Assign an AppRole to an application",
	Long: `This command enables you to assign an AppRole to an application.

	Example:
		./ecli approle assign --approleid "your-approle-id" --appid "your-app-id"
`,
	Run: func(cmd *cobra.Command, args []string) {
		if OptAppRoleID == "" {
			rootcmd.PrintErr("AppRoleID is required (use --approleid)")
			return
		}
		if rootcmd.OptAppID == "" {
			rootcmd.PrintErr("AppID is required (use --appid)")
			return
		}

		// Get the service principal for the application
		sp, err := rootcmd.Client.GetServicePrincipalByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		// Create the app role assignment
		assignment := &models.AppRoleAssignment{
			PrincipalID:   sp.ID,
			PrincipalType: "ServicePrincipal",
			AppRoleID:     OptAppRoleID,
			ResourceID:    sp.ID,
		}

		err = rootcmd.Client.AssignAppRoleToServicePrincipal(assignment, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Println("AppRole assigned successfully")
		cmd.Println(rootcmd.OutputJSON(assignment))
	},
}

func init() {
	approleCmd.AddCommand(approleAssignCmd)

	approleAssignCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		approleAssignCmd.Flags().MarkHidden("batch")
		approleAssignCmd.Flags().MarkHidden("search")
		approleAssignCmd.Flags().MarkHidden("select")
		approleAssignCmd.Flags().MarkHidden("skip")
		approleAssignCmd.Flags().MarkHidden("skiptoken")
		approleAssignCmd.Flags().MarkHidden("top")
		approleAssignCmd.Flags().MarkHidden("post")
		approleAssignCmd.Flags().MarkHidden("default")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
