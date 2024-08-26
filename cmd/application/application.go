// Package cmdapplication is used for application commands
package cmdapplication

import (
	rootcmd "epfl-entra/cmd"

	"github.com/spf13/cobra"
)

// OptRedirectURI is associated with the --redirect_uri flag
var OptRedirectURI string

// OptHomeURI is associated with the --home_uri flag
var OptHomeURI string

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "Manage applications",
	Long: `This command enables you to
	* Create
	* Get
	* Modify
	* Delete

	application(s).
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Use cmd.Println() instead of fmt.Println() to be able to capture the output (in tests)
		cmd.Println("application called")
	},
}

func init() {
	rootcmd.RootCmd.AddCommand(applicationCmd)
	applicationCmd.PersistentFlags().StringVar(&OptRedirectURI, "redirect_uri", "", "Redirect URI")
	applicationCmd.PersistentFlags().StringVar(&OptHomeURI, "home_uri", "", "Homepage URI")
}
