package cmdcertificate

import (
	"encoding/base64"
	"io"
	"os"
	"time"

	rootcmd "github.com/epfl-si/entra-client/cmd"

	"github.com/spf13/cobra"
)

var (
	// OptCertFile is associated with the --cert-file flag
	OptCertFile string
	// OptCertBase64 is associated with the --cert-base64 flag
	OptCertBase64 string
	// OptDisplayName is associated with the --display-name flag
	OptDisplayName string
	// OptStartDate is associated with the --start-date flag
	OptStartDate string
	// OptEndDate is associated with the --end-date flag
	OptEndDate string
)

// certificateAddCmd represents the certificateAdd command
var certificateAddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a certificate to an application",
	Long: `Add a certificate (key credential) to an application

Example:
  ecli certificate add --appid <app-id> --display-name "My Cert" --cert-file /path/to/cert.pem
  ecli certificate add --appid <app-id> --display-name "My Cert" --cert-base64 <base64-encoded-cert>
  ecli certificate add --appid <app-id> --display-name "My Cert" --cert-file /path/to/cert.pem --start-date 2024-01-01 --end-date 2025-01-01`,
	Run: func(cmd *cobra.Command, args []string) {
		if rootcmd.OptAppID == "" {
			cmd.PrintErr("Application ID is required (use --appid)\n")
			return
		}
		if OptDisplayName == "" {
			cmd.PrintErr("Display name is required (use --display-name)\n")
			return
		}

		// Get certificate content
		var certContent string
		if OptCertBase64 != "" {
			certContent = OptCertBase64
		} else if OptCertFile != "" {
			file, err := os.Open(OptCertFile)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}
			defer file.Close()

			certBytes, err := io.ReadAll(file)
			if err != nil {
				rootcmd.PrintErr(err)
				return
			}

			certContent = base64.StdEncoding.EncodeToString(certBytes)
		} else {
			cmd.PrintErr("Certificate is required (use --cert-file or --cert-base64)\n")
			return
		}

		// Set default dates if not provided
		startDate := OptStartDate
		if startDate == "" {
			startDate = time.Now().UTC().Format("2006-01-02T15:04:05Z")
		}

		endDate := OptEndDate
		if endDate == "" {
			// Default to 1 year from start date
			start, err := time.Parse("2006-01-02T15:04:05Z", startDate)
			if err != nil {
				start = time.Now().UTC()
			}
			endDate = start.AddDate(1, 0, 0).Format("2006-01-02T15:04:05Z")
		}

		err := rootcmd.Client.AddKeyCredentialToApplication(
			rootcmd.OptAppID,
			OptDisplayName,
			startDate,
			endDate,
			certContent,
			rootcmd.ClientOptions,
		)
		if err != nil {
			rootcmd.PrintErr(err)
			return
		}

		cmd.Printf("Certificate '%s' added successfully to application %s\n", OptDisplayName, rootcmd.OptAppID)
	},
}

func init() {
	certificateCmd.AddCommand(certificateAddCmd)

	certificateAddCmd.Flags().StringVar(&OptCertFile, "cert-file", "", "Path to certificate file (PEM or DER format)")
	certificateAddCmd.Flags().StringVar(&OptCertBase64, "cert-base64", "", "Base64-encoded certificate")
	certificateAddCmd.Flags().StringVar(&OptDisplayName, "display-name", "", "Display name for the certificate (required)")
	certificateAddCmd.Flags().StringVar(&OptStartDate, "start-date", "", "Start date (ISO 8601 format, e.g., 2024-01-01T00:00:00Z)")
	certificateAddCmd.Flags().StringVar(&OptEndDate, "end-date", "", "End date (ISO 8601 format, e.g., 2025-01-01T00:00:00Z)")

	certificateAddCmd.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		certificateAddCmd.Flags().MarkHidden("batch")
		certificateAddCmd.Flags().MarkHidden("id")
		certificateAddCmd.Flags().MarkHidden("post")
		certificateAddCmd.Flags().MarkHidden("search")
		certificateAddCmd.Flags().MarkHidden("select")
		certificateAddCmd.Flags().MarkHidden("skip")
		certificateAddCmd.Flags().MarkHidden("skiptoken")
		certificateAddCmd.Flags().MarkHidden("top")
		// Call parent help func
		command.Parent().HelpFunc()(command, strings)
	})
}
