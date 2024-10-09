// Package extensioncmd is used for extension commands
package extensioncmd

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// extensionCmd represents the extension command
var extensionCmd = &cobra.Command{
	Use:   "extension",
	Short: "Manage extensions",
	Long: `This command enables you to
	* Create
	* List
	* Modify
	* Delete

	extension(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use cmd.Println() instead of fmt.Println() to be able to capture the output (in tests)
		cmd.Println("extension called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(extensionCmd)
}
