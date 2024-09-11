package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"time"

	"github.com/epfl-si/entra-client/pkg/client/models"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// CreateApplication create an application
func CreateApplication(app *models.Application, clientOptions models.ClientOptions) (*models.Application, *models.ServicePrincipal, error) {

	newApp, err := Client.CreateApplication(app, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("CreateApplication: %w", err)
	}

	err = Client.WaitApplication(newApp.ID, 60, clientOptions)
	if err != nil {
		return nil, nil, fmt.Errorf("WaitApplication: %w", err)
	}

	sp, err := Client.CreateServicePrincipal(&models.ServicePrincipal{
		AppID:                newApp.AppID,
		ServicePrincipalType: "Application"}, clientOptions)

	if err != nil {
		return nil, nil, fmt.Errorf("CreateServicePrincipal: %W", err)
	}

	return newApp, sp, nil
}

// NormalizeURI performs some modification (required by Microsoft QPI) on URI
//   - removes the trailing slash from a string
//   - replace http with api
func NormalizeURI(s string) string {
	var n string
	if len(s) > 0 && s[len(s)-1] == '/' {
		n = s[:len(s)-1]
		s = n
	}

	if len(s) > 5 && s[:5] == "http:" {
		n = "https:" + s[5:]
		s = n
	}

	return s
}

// OutputJSON returns a JSON representation of the data
func OutputJSON(data interface{}) string {
	jdata, err := json.Marshal(data)
	if err != nil {
		PrintErr(err)
		os.Exit(0)
	}
	result := string(jdata)
	if OptPrettyJSON {
		var out bytes.Buffer
		err := json.Indent(&out, []byte(result), "", "  ")
		if err != nil {
			PrintErr(err)
			os.Exit(0)
		}
		result = out.String()
	}

	return result
}

// AddCertificate adds a certificate to a Service Principal
//   - spID: the Service Principal ID
//   - certUsage: the certificate usage (e.g. 'Verify'/'Sign')
//   - certType: the certificate type	(e.g. 'AsymmetricX509Cert')
//   - certBase64: the certificate in base64 format
//
// Resources:
// - https://github.com/MicrosoftDocs/azure-docs/issues/58484
// (Why Graph API is really misleading)
func AddCertificate(spID string, certUsage, certType, certBase64 string, clientOptions models.ClientOptions) error {

	if certUsage != "Verify" && certUsage != "Sign" {
		return fmt.Errorf("Invalid certificate usage: %s", certUsage)
	}

	sp, err := Client.GetServicePrincipal(spID, clientOptions)
	if err != nil {
		return fmt.Errorf("could'nt get sp: %w", err)
	}

	keyCredentials := []models.KeyCredential{}

	startDateTime := time.Now()
	startDateTime, _ = time.Parse(time.RFC3339, startDateTime.String())
	endDateTime := startDateTime.AddDate(0, 0, 364)
	endDateTime, _ = time.Parse(time.RFC3339, endDateTime.String())
	// Format date to this format: "2024-01-11T15:31:26Z
	// Weird bug due to Timezone, can make end date off of few hours

	// Build new KeyCredential
	keyID := NormalizeThumbprint(uuid.Must(uuid.NewRandom()).String())
	newKeyCredential := models.KeyCredential{
		CustomKeyIdentifier: keyID,
		// EndDateTime: endDateTime,
		// KeyId:         keyID,
		// StartDateTime: startDateTime,
		DisplayName: sp.DisplayName + " " + certUsage + "ing certificate",
		Usage:       certUsage,
		Type:        certType,
		// Key:         "base64" + certBase64,
		// Key: []byte(certBase64),
		Key: certBase64,
	}

	keyCredentials = append(keyCredentials, newKeyCredential)

	// Build new PasswordCredential
	// newPasswordCredential := models.PasswordCredential{
	// 	CustomKeyIdentifier: keyID,
	// 	// EndDateTime: endDateTime,
	// 	KeyID: keyID,
	// 	// StartDateTime: startDateTime,
	// 	DisplayName: sp.DisplayName + " " + certUsage + "ing certificate",
	// 	// Secret text is null for signing certificates
	// }

	// if sp.KeyCredentials != nil {
	// 	keyCredentials = sp.KeyCredentials

	// } else {
	// 	keyCredentials = []models.KeyCredential{}
	// }

	spPatch := models.ServicePrincipal{
		KeyCredentials: keyCredentials,
		// PasswordCredentials: sp.PasswordCredentials,
	}

	// if certUsage == "Verify" {
	// 	sp.PasswordCredentials = append(sp.PasswordCredentials, newPasswordCredential)
	// 	spPatch.PasswordCredentials = sp.PasswordCredentials
	// }

	err = Client.PatchServicePrincipal(spID, &spPatch, clientOptions)
	if err != nil {
		return fmt.Errorf("could'nt patch sp for KeyCredentials: %w", err)
	}

	// Get the updated Service Principal
	sp, err = Client.GetServicePrincipal(spID, clientOptions)
	if err != nil {
		return fmt.Errorf("could'nt get updated sp: %w", err)
	}

	// Activate the certificate by its keyId
	// err = Client.PatchServicePrincipal(spID, &models.ServicePrincipal{PreferredTokenSigningKeyThumbprint: normalizeThumbprint(sp.KeyCredentials[0].CustomKeyIdentifier)}, clientOptions)
	// if err != nil {
	// 	return fmt.Errorf("could'nt patch sp to activate certificate: %w", err)
	// }

	return nil
}

// HideInCommand hides flags in a command
func HideInCommand(c *cobra.Command, flags ...string) {
	c.SetHelpFunc(func(command *cobra.Command, strings []string) {
		// Hide flags for this command
		for _, flag := range flags {
			fmt.Println("Hiding flag: ", flag)
			c.Flags().MarkHidden(flag)
		}
		// Call parent help func
		c.Parent().HelpFunc()(command, strings)
	})

}

// PrintErr prints an error to stderr
func PrintErr(err error) {
	fmt.Fprintln(os.Stderr, err)
}

// PrintErrString prints an error string to stderr
func PrintErrString(str string) {
	fmt.Fprintln(os.Stderr, str)
}

// NormalizeThumbprint removes spaces and dashes from a thumbprint
func NormalizeThumbprint(thumbprint string) string {
	re, _ := regexp.Compile(`[\s\-]`)
	thumbprint = re.ReplaceAllString(thumbprint, "")

	return thumbprint
}

// CaptureOutput redirect stdout/stderr to pipes and keep the old values
// rout out reader, wout out writer, oldout old out writer
// rerr err reader, werr err writer, olderr old err writer
func CaptureOutput() (rout, wout, oldout, rerr, werr, olderr *os.File) {
	oldout = os.Stdout
	rout, wout, _ = os.Pipe()
	os.Stdout = wout

	olderr = os.Stderr
	rerr, werr, _ = os.Pipe()
	os.Stderr = werr

	return rout, wout, oldout, rerr, werr, olderr
}

// ReleaseOutput read the ghoutput from the pipes and restore the old values
func ReleaseOutput(rout, wout, oldout, rerr, werr, olderr *os.File) (out, err []byte) {
	// read output
	wout.Close()
	out, _ = io.ReadAll(rout)
	rout.Close()
	os.Stdout = oldout

	// read err
	werr.Close()
	err, _ = io.ReadAll(rerr)
	rerr.Close()
	os.Stderr = olderr

	return out, err
}
