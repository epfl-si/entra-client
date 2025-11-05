package cmdapprole

import (
	"encoding/json"

	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// approleCreateCmd represents the approleCreate command
var approleCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create an AppRole",
	Long: `This command enables you to create an AppRole.

	Example:
		./ecli approle create --appid "your-app-id" --data '{"displayName":"Default Access","description":"Default access (used for M2M)","isEnabled":true,"value":"DefaultAccess","allowedMemberTypes":["Application"]}'
	    ./ecli approle create --appid "your-app-id" --default # to create a default AppRole named "Default Access" for M2M scenarios
`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptAppID == "" {
			rootcmd.PrintErr("AppID is required (use --appid)")
			return
		}
		if rootcmd.OptPostData == "" && !rootcmd.OptDefault {
			cmd.PrintErr("Data or default flag is required (use --data or --default)")
			return
		}

		if rootcmd.OptPostData != "" && rootcmd.OptDefault {
			cmd.PrintErr("Data OR default flag are mutually exclusive (use --data OR --default)")
			return
		}

		var appRole models.AppRole

		if rootcmd.OptDefault {
			rootcmd.ClientOptions.Default = true
		}

		if rootcmd.OptPostData != "" {
			err := json.Unmarshal([]byte(rootcmd.OptPostData), &appRole)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}
		}

		id, err := rootcmd.Client.CreateAppRoleByAppID(rootcmd.OptAppID, &appRole, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}
		appRole.ID = id

		cmd.Println(rootcmd.OutputJSON(appRole))
	},
}

func init() {
	approleCmd.AddCommand(approleCreateCmd)

	approleCreateCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		approleCreateCmd.Flags().MarkHidden("batch")
		approleCreateCmd.Flags().MarkHidden("search")
		approleCreateCmd.Flags().MarkHidden("select")
		approleCreateCmd.Flags().MarkHidden("skip")
		approleCreateCmd.Flags().MarkHidden("skiptoken")
		approleCreateCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
