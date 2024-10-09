package extensioncmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// extensionListCmd represents the extensionList command
var extensionListCmd = &cobra.Command{
	Use:   "list",
	Short: "List extensions",
	Run: func(cmd *cobra.Command, args []string) {
		extensions, err := rootcmd.Client.GetExtensions(rootcmd.ClientOptions)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		for _, extension := range extensions {
			cmd.Println(rootcmd.OutputJSON(extension))
		}
	},
}

func init() {
	extensionCmd.AddCommand(extensionListCmd)
}
