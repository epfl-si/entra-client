package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/epfl-si/entra-client/pkg/client/models"
	"github.com/joho/godotenv"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

// NormalizeName normalizes a name by adding prefixing with EPFL and suffixing with environment when it's not PROD, if it's not already the case
// Note: " - " is replaced by "-" in the name.
//   - name: the name to normalize
//   - env: the environment (1: DEV, 2: TEST, 3: PROD)
//
// Returns the normalized name
func NormalizeName(name string, env int) (string, error) {
	// First remove any existing prefix/suffix
	re := regexp.MustCompile(`^(?:EPFL - )?(.*?)(?:\s*-\s*(?i:TEST|DEV))?$`)
	matches := re.FindStringSubmatch(name)

	// Use the captured group if there's a match, otherwise use the original name
	n := name
	if len(matches) > 1 {
		n = matches[1]
	}

	// Ensure bare name contains no " - " => replace with "-"
	re = regexp.MustCompile(` - `)
	n = re.ReplaceAllString(n, "-")

	switch env {
	case 1:
		return fmt.Sprintf("EPFL - %s - DEV", n), nil
	case 2:
		return fmt.Sprintf("EPFL - %s - TEST", n), nil
	case 3:
		return fmt.Sprintf("EPFL - %s", n), nil
	default:
		return "", fmt.Errorf("invalid environment: %d, must be 1, 2 or 3", env)
	}
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
		return fmt.Errorf("invalid certificate usage: %s", certUsage)
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

	spPatch := models.ServicePrincipal{
		KeyCredentials: keyCredentials,
	}

	err = Client.PatchServicePrincipal(spID, &spPatch, clientOptions)
	if err != nil {
		return fmt.Errorf("could'nt patch sp for KeyCredentials: %w", err)
	}

	// Get the updated Service Principal
	sp, err = Client.GetServicePrincipal(spID, clientOptions)
	if err != nil {
		return fmt.Errorf("could'nt get updated sp: %w", err)
	}

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

// PrintErr prints an error or string to stderr
func PrintErr(v interface{}) {
	switch val := v.(type) {
	case error:
		fmt.Fprintln(os.Stderr, val)
	case string:
		if !strings.HasSuffix(val, "\n") {
			val += "\n"
		}
		fmt.Fprint(os.Stderr, val)
	default:
		fmt.Fprintln(os.Stderr, val)
	}
}

// NormalizeThumbprint removes spaces and dashes from a thumbprint
func NormalizeThumbprint(thumbprint string) string {
	re, _ := regexp.Compile(`[\s\-]`)
	thumbprint = re.ReplaceAllString(thumbprint, "")

	return thumbprint
}

func findGoModDir() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("go.mod not found in any parent directory")
		}
		dir = parent
	}
}

// LoadEnv loads environment variables from a .env file located in the same directory as the go.mod file
func LoadEnv(envFile string) error {
	modDir, err := findGoModDir()
	if err != nil {
		fmt.Printf("Error finding go.mod directory: %s\n", err)
		return err
	}

	// Load environment variables from .env file in the go.mod directory
	err = godotenv.Load(filepath.Join(modDir, envFile))
	if err != nil {
		fmt.Printf("Error loading %s file: %s", envFile, err)
		return err
	}
	return nil
}

// CaptureStdOutputs create 2 bytes buffers to capture the output and error of a command
func CaptureStdOutputs(cmd *cobra.Command) (newOut, newErr *bytes.Buffer) {

	// Capture output
	stdOut := new(bytes.Buffer)
	stdErr := new(bytes.Buffer)
	cmd.SetOut(stdOut)
	cmd.SetErr(stdErr)

	return stdOut, stdErr
}
