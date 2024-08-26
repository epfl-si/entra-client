package groupcmd

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// groupGetCmd represents the groupGet command
var groupGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get a group by ID",
	Run: func(cmd *cobra.Command, args []string) {
		group, err := rootcmd.Client.GetGroup(rootcmd.OptID, rootcmd.ClientOptions)
		if err != nil {
			cmd.PrintErr(err)
			return
		}

		cmd.Printf("Group: %s\n", rootcmd.OutputJSON(group))
	},
}

func init() {
	groupCmd.AddCommand(groupGetCmd)

	groupGetCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		groupGetCmd.Flags().MarkHidden("batch")
		groupGetCmd.Flags().MarkHidden("search")
		groupGetCmd.Flags().MarkHidden("select")
		groupGetCmd.Flags().MarkHidden("skip")
		groupGetCmd.Flags().MarkHidden("skiptoken")
		groupGetCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
