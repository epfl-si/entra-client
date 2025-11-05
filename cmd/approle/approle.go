// Package cmdapprole is used for AppRole commands
package cmdapprole

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// OptAppRoleID is associated with the --approleid flag
var OptAppRoleID string

// approleCmd represents the AppRole command
var approleCmd = &cobra.Command{
	Use:     "approle",
	Aliases: []string{"ar"},
	Short:   "Manage AppRoles",
	Long: `This command enables you to
	* Create
	* Get
	* Modify
	* Delete

	AppRole(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use cmd.Println() instead of fmt.Println() to be able to capture the output (in tests)
		cmd.Println("approle called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(approleCmd)

	approleCmd.PersistentFlags().StringVar(&OptAppRoleID, "approleid", "", "AppRole ID to use")
}
