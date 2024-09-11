package cmdapplication

import (
	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

// applicationSAMLCertificateCmd represents the applicationSAMLCertificate command
var applicationSAMLCertificateCmd = &cobra.Command{
	Use:   "certificate",
	Short: "Manage certificates",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Println("applicationSAMLCertificate called")
	},
}

func init() {
	applicationSAMLCmd.AddCommand(applicationSAMLCertificateCmd)

	rootcmd.HideInCommand(applicationSAMLCertificateCmd, "top", "skip", "skiptoken", "select", "search")
}
