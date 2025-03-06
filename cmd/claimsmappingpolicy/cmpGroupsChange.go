package cmdclaimsmappingpolicy

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"
	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/spf13/cobra"
)

// OptGroupAll is a flag to use all groups
var OptGroupAll = false

// OptGroupNone is a flag to disable groups
var OptGroupNone = false

// OptGroupSecurity is a flag to use security groups
var OptGroupSecurity = false

// OptGroupApplication is a flag to use application groups
var OptGroupApplication = false

// claimGroupsAddCmd represents the claims mapping policy usage command
var claimGroupsAddCmd = &cobra.Command{
	Use:   "groups change",
	Short: "Change the groups of a claims mapping policy",
	Long: `This command enables you to change the groups part in a claims mapping policy.

	Example:
		./ecli claim groups change --appid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --none
		./ecli claim groups change --appid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --all
		./ecli claim groups change --appid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --security
		./ecli claim groups change --appid 52c47ba8-f2d2-4c9b-9395-3654fc7d2b51 --application
	`,

	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptAppID == "" {
			rootcmd.PrintErrString("ID is required (use --appid)")
			return
		}

		var applicationModification models.Application

		if OptGroupAll {
			applicationModification.GroupMembershipClaims = "All"
		} else if OptGroupNone {
			applicationModification.GroupMembershipClaims = "None"
		} else if OptGroupSecurity {
			applicationModification.GroupMembershipClaims = "SecurityGroup"
		} else if OptGroupApplication {
			applicationModification.GroupMembershipClaims = "ApplicationGroup"
		} else {
			rootcmd.PrintErrString("One of the following flags is required: --all, --null, --security, --application")
			return
		}

		var ID string
		app, err := rootcmd.Client.GetApplicationByAppID(rootcmd.OptAppID, rootcmd.ClientOptions)
		if err != nil {
			if err.Error() == "application not found" {
				ID = rootcmd.OptAppID
			} else {
				rootcmd.PrintErr(err)
				return
			}
		} else {
			ID = app.ID
		}

		err = rootcmd.Client.PatchApplication(ID, &applicationModification, rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErrString("Application not found by Object ID: " + app.ID)
			rootcmd.PrintErr(err)
			return
		}

	},
}

func init() {
	claimCmd.AddCommand(claimGroupsAddCmd)

	claimGroupsAddCmd.Flags().BoolVar(&OptGroupAll, "all", false, "Use all groups")
	claimGroupsAddCmd.Flags().BoolVar(&OptGroupNone, "none", false, "Disable groups")
	claimGroupsAddCmd.Flags().BoolVar(&OptGroupSecurity, "security", false, "Use security groups")
	claimGroupsAddCmd.Flags().BoolVar(&OptGroupApplication, "application", false, "Use application groups")

	claimGroupsAddCmd.MarkFlagsMutuallyExclusive("all", "none", "security", "application")

	claimGroupsAddCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		claimGroupsAddCmd.Flags().MarkHidden("batch")
		claimGroupsAddCmd.Flags().MarkHidden("search")
		claimGroupsAddCmd.Flags().MarkHidden("select")
		claimGroupsAddCmd.Flags().MarkHidden("skip")
		claimGroupsAddCmd.Flags().MarkHidden("skiptoken")
		claimGroupsAddCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
