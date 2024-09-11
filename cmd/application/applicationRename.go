package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// applicationRenameCmd represents the applicationRename command
var applicationRenameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Rename an application",
	Long: `Rename an application by providing the appID and the new name:

Example:

  ./ecli application rename --id 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --displayname "New name" 
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptID == "" {
			rootcmd.PrintErrString("Service Principal ID is required (use --id)")
			return
		}

		if rootcmd.OptDisplayName == "" {
			rootcmd.PrintErrString("New display name is required (use --displayname)")
			return
		}

		rootcmd.ClientOptions.Filter = "appId%20eq%20'" + rootcmd.OptID + "'"
		apps, _, err := rootcmd.Client.GetApplications(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		if len(apps) == 0 {
			rootcmd.PrintErrString("Application not found")
			return
		}
		if len(apps) != 1 {
			rootcmd.PrintErrString("Ambigouous application ID")
			return
		}

		newApp := models.Application{
			DisplayName: rootcmd.OptDisplayName,
		}

		err = rootcmd.Client.PatchApplication(apps[0].ID, &newApp, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
	},
}

func init() {
	applicationCmd.AddCommand(applicationRenameCmd)

	// hideInCommand(applicationRenameCmd, "top', 'skip', 'skiptoken', 'select', 'search")
	applicationRenameCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		applicationRenameCmd.Flags().MarkHidden("batch")
		applicationRenameCmd.Flags().MarkHidden("search")
		applicationRenameCmd.Flags().MarkHidden("select")
		applicationRenameCmd.Flags().MarkHidden("skip")
		applicationRenameCmd.Flags().MarkHidden("skiptoken")
		applicationRenameCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
