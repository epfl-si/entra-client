// Secret Alerter for Entra
package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/joho/godotenv"

	"github.com/epfl-si/entra-client/pkg/client/httpengine"
	"github.com/epfl-si/entra-client/pkg/client/models"
	mail_sender "github.com/go-mail/mail"
)

// Config holds configuration for the secret alerter
type Config struct {
	SMTPHost            string
	SMTPPort            int
	FromEmail           string
	ToEmail             string
	Subject             string
	ExpiryThresholdDays int64
	Fake                bool
}

// AlertData holds data for the email template
type AlertData struct {
	TotalApplications int
	TotalSecrets      int
	Applications      []ApplicationAlert
	GeneratedAt       time.Time
	ThresholdDays     int64
}

// ApplicationAlert holds data for each application with expiring secrets
type ApplicationAlert struct {
	AppID       string
	DisplayName string
	Secrets     []SecretAlert
}

// SecretAlert holds data for each expiring secret
type SecretAlert struct {
	Type          string
	KeyID         string
	DisplayName   string
	ExpiryDate    time.Time
	RemainingDays int64
}

// Options
var keyClientOptions = models.ClientOptions{
	Select: "keyCredentials,displayName,appId",
}
var passwordClientOptions = models.ClientOptions{
	Select: "passwordCredentials,displayName,appId",
}

const sevenDays = -7

const emailTemplate = `<!DOCTYPE html
<html>
<head>
    <title>Entra Secrets Expiring Soon</title>
    <style>
        body { font-family: Arial, sans-serif; margin: 20px; }
        .header { background-color: #f8f9fa; padding: 20px; border-radius: 5px; margin-bottom: 20px; }
        .summary { background-color: #fff3cd; padding: 15px; border-radius: 5px; margin-bottom: 20px; }
        .application { margin-bottom: 30px; }
        .app-header { background-color: #e9ecef; padding: 10px; border-radius: 5px; margin-bottom: 10px; }
        .secret { margin-left: 20px; padding: 10px; border-left: 3px solid #ffc107; margin-bottom: 10px; }
        .critical { border-left-color: #dc3545; }
        .warning { border-left-color: #fd7e14; }
        .footer { margin-top: 30px; font-size: 12px; color: #6c757d; }
        table { border-collapse: collapse; width: 100%; }
        th, td { border: 1px solid #ddd; padding: 8px; text-align: left; }
        th { background-color: #f2f2f2; }
    </style>
</head>
<body>
    <div class="header">
        <h1>Expiring Entra Secrets</h1>
        <p>Generated on: {{.GeneratedAt.Format "2006-01-02 15:04:05 UTC"}}</p>
    </div>

    <div class="summary">
        <h2>Summary</h2>
        <p><strong>{{.TotalApplications}}</strong> applications have <strong>{{.TotalSecrets}}</strong> secrets expiring within {{.ThresholdDays}} days</p>
    </div>

    {{range .Applications}}
    <div class="application">
        <div class="app-header">
            <h3>{{.DisplayName}}</h3>
            <p><strong>App ID:</strong> {{.AppID}}</p>
        </div>
        
        {{range .Secrets}}
        <div class="secret {{if lt .RemainingDays 7}}critical{{else if lt .RemainingDays 14}}warning{{end}}">
            <h4>{{if .DisplayName}}{{.DisplayName}}{{else}}Secret {{.KeyID}}{{end}}</h4>
			<p><strong>Type:</strong> {{.Type}}</p>
            <p><strong>Key ID:</strong> {{.KeyID}}</p>
            <p><strong>Expires:</strong> {{.ExpiryDate.Format "2006-01-02 15:04:05 UTC"}}</p>
            <p><strong>Days remaining:</strong> 
                {{if lt .RemainingDays 0}}
                    <span style="color: #dc3545; font-weight: bold;">EXPIRED ({{.RemainingDays}} days ago)</span>
                {{else}}
                    <span style="color: #ffc107;">{{.RemainingDays}} days</span>
                {{end}}
            </p>
        </div>
        {{end}}
    </div>
    {{end}}

    <div class="footer">
        <p>This is an automated alert from the Entra Secret Alerter.</p>
        <p>Please review and renew expiring secrets as needed.</p>
    </div>
</body>
</html>
`

func main() {
	// Load configuration from environment variables
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	var fake *bool
	fake = flag.Bool("fake", false, "fake mode")
	flag.Parse()

	config.Fake = *fake

	// Initialize Entra client
	client, err := httpengine.New()
	if err != nil {
		log.Fatalf("Failed to initialize Entra client: %v", err)
	}

	// Calculate date 30 days from now
	futureDate := time.Now().AddDate(0, 0, int(config.ExpiryThresholdDays))
	LimitDate := futureDate.Format("2006-01-02")

	log.Printf("Checking for secrets expiring before: %s", LimitDate)

	keyCredentialsMap, err := client.GetLocalKeyCredentials(LimitDate, keyClientOptions)
	if err != nil {
		log.Fatalf("Failed to get key credentials: %v", err)
	}

	passwordCredentialsMap, err := client.GetLocalPasswordCredentials(LimitDate, passwordClientOptions)
	if err != nil {
		log.Fatalf("Failed to get password credentials: %v", err)
	}

	// Filter secrets expiring within 30 days (RemainingDays <= Threshold and >= -1)
	alertData := processCredentials(config.ExpiryThresholdDays, keyCredentialsMap, passwordCredentialsMap, client, LimitDate)

	if alertData.TotalSecrets == 0 {
		log.Printf("No secrets expiring within %d days found\n", config.ExpiryThresholdDays)
		return
	}

	alertData.ThresholdDays = config.ExpiryThresholdDays
	log.Printf("Found %d applications with %d expiring secrets withing %d days", alertData.TotalApplications, alertData.TotalSecrets, alertData.ThresholdDays)

	err = sendEmailAlert(config, alertData)
	if err != nil {
		log.Fatalf("Failed to send email alert: %v", err)
	}

}

func loadConfig() (*Config, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	config := &Config{}

	config.SMTPHost = getEnvOrDefault("SMTP_HOST", "mail.epfl.ch")
	config.SMTPPort = 25 // Fixed port for EPFL mail server
	config.FromEmail = getEnvOrDefault("FROM_EMAIL", "noreply@epfl.ch")
	config.ToEmail = getEnvOrDefault("TO_EMAIL", "idev-md@groupes.epfl.ch")

	var err error
	config.ExpiryThresholdDays, err = strconv.ParseInt(getEnvOrDefault("EXPIRY_THRESHOLD_DAYS", "30"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid EXPIRY_THRESHOLD_DAYS: %v", err)
	}

	config.Subject = getEnvOrDefault("EMAIL_SUBJECT", "Entra Secrets Expiring Soon")

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// processKeyCredentials extracts expiring key credentials from the provided map
func processKeyCredentials(expiryThresholdDays int64, keyCredentialsMap map[string][]models.KeyCredentialEPFL) map[string]*ApplicationAlert {
	allApplications := make(map[string]*ApplicationAlert)

	for appID, credentials := range keyCredentialsMap {
		var expiringSecrets []SecretAlert
		var appDisplayName string

		for _, cred := range credentials {
			// Calculate actual remaining days from today
			actualRemainingDays := int64(time.Until(time.Time(*cred.EndDateTime)).Hours() / 24)

			// Only include secrets expiring within threshold (including recently expired ones up to 7 days)
			if actualRemainingDays <= expiryThresholdDays && actualRemainingDays >= sevenDays {
				secret := SecretAlert{
					Type:          cred.Type,
					KeyID:         cred.KeyID,
					DisplayName:   cred.DisplayName,
					ExpiryDate:    time.Time(*cred.EndDateTime),
					RemainingDays: actualRemainingDays, // Use actual remaining days
				}
				expiringSecrets = append(expiringSecrets, secret)
				appDisplayName = cred.AppDisplayName // Get display name from any credential
			}
		}

		if len(expiringSecrets) > 0 {
			allApplications[appID] = &ApplicationAlert{
				AppID:       appID,
				DisplayName: appDisplayName,
				Secrets:     expiringSecrets,
			}
		}
	}

	return allApplications
}

// processPasswordCredentials processes password credentials and merges with existing applications
func processPasswordCredentials(expiryThresholdDays int64, passwordCredentialsMap map[string][]models.PasswordCredentialEPFL, allApplications map[string]*ApplicationAlert) {
	for appID, credentials := range passwordCredentialsMap {
		var expiringSecrets []SecretAlert
		var appDisplayName string

		for _, cred := range credentials {
			// Calculate actual remaining days from today
			actualRemainingDays := int64(time.Until(cred.EndDateTime).Hours() / 24)

			// Only include secrets expiring within threshold (including recently expired ones up to 7 days)
			if actualRemainingDays <= expiryThresholdDays && actualRemainingDays >= sevenDays {
				secret := SecretAlert{
					Type:          "secret",
					KeyID:         cred.KeyID,
					DisplayName:   cred.DisplayName,
					ExpiryDate:    cred.EndDateTime,
					RemainingDays: actualRemainingDays, // Use actual remaining days
				}
				expiringSecrets = append(expiringSecrets, secret)
				appDisplayName = cred.AppDisplayName // Get display name from any credential
			}
		}

		if len(expiringSecrets) > 0 {
			if existingApp, exists := allApplications[appID]; exists {
				// Merge with existing key credentials for this app
				existingApp.Secrets = append(existingApp.Secrets, expiringSecrets...)
			} else {
				// Create new entry for this app
				allApplications[appID] = &ApplicationAlert{
					AppID:       appID,
					DisplayName: appDisplayName,
					Secrets:     expiringSecrets,
				}
			}
		}
	}
}

// filterApplicationsWithValidSecrets removes applications that have valid secrets beyond the limit date
func filterApplicationsWithValidSecrets(allApplications map[string]*ApplicationAlert, client *httpengine.HTTPClient, limitDate string) {
	futureDate := time.Date(2099, 12, 31, 23, 59, 59, 0, time.UTC)
	limitDateTime, _ := time.Parse("2006-01-02", limitDate)

	for appID, appAlert := range allApplications {
		hasValidKey := false
		hasValidPassword := false
		alerts := make([]SecretAlert, 0)

		for _, alert := range appAlert.Secrets {
			if alert.Type == "AsymmetricX509Cert" || alert.Type == "Symmetric" {
				allKeys, err := client.GetKeyCredentialsByAppID(futureDate.Format("2006-01-02"), appID, keyClientOptions)
				if err == nil {
					for _, key := range allKeys {
						if time.Time(*key.EndDateTime).After(limitDateTime) {
							hasValidKey = true
							break
						}
					}
					if !hasValidKey {
						alerts = append(alerts, alert)
					}
				}
			}
			if alert.Type == "secret" {
				allPasswords, err := client.GetPasswordCredentialsByAppID(futureDate.Format("2006-01-02"), appID, passwordClientOptions)
				if err == nil {
					for _, pwd := range allPasswords {
						if pwd.EndDateTime.After(limitDateTime) {
							hasValidPassword = true
							break
						}
					}
				}
				if !hasValidPassword {
					alerts = append(alerts, alert)
				}
			}
		}

		if len(alerts) == 0 {
			delete(allApplications, appID) // Remove application if no expiring secrets remain
		} else {
			appAlert.Secrets = alerts // Update with filtered secrets
		}
	}
}

// buildAlertData converts the applications map to AlertData structure with totals
func buildAlertData(allApplications map[string]*ApplicationAlert) *AlertData {
	alertData := &AlertData{
		Applications: make([]ApplicationAlert, 0),
		GeneratedAt:  time.Now(),
	}

	// Convert map to slice and count totals
	for _, appAlert := range allApplications {
		alertData.Applications = append(alertData.Applications, *appAlert)
		alertData.TotalSecrets += len(appAlert.Secrets)
	}

	alertData.TotalApplications = len(alertData.Applications)
	return alertData
}

// processCredentials processes both key and password credentials and filters for expiring ones
func processCredentials(expiryThresholdDays int64, keyCredentialsMap map[string][]models.KeyCredentialEPFL, passwordCredentialsMap map[string][]models.PasswordCredentialEPFL, client *httpengine.HTTPClient, limitDate string) *AlertData {
	// Process key credentials first
	allApplications := processKeyCredentials(expiryThresholdDays, keyCredentialsMap)

	// Process password credentials and merge with key credentials
	processPasswordCredentials(expiryThresholdDays, passwordCredentialsMap, allApplications)

	// Filter applications that have valid secrets beyond the limit date
	filterApplicationsWithValidSecrets(allApplications, client, limitDate)

	// Build and return alert data
	return buildAlertData(allApplications)
}

func sendEmailAlert(config *Config, alertData *AlertData) error {
	// Parse and execute template
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	// Execute template to get HTML body
	var emailBuffer strings.Builder
	err = tmpl.Execute(&emailBuffer, alertData)
	if err != nil {
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	body := emailBuffer.String()

	if config.Fake {
		log.Println("Fake mode enabled, not sending email")
		log.Printf("Generated email body:\n%s\n", body)

		return nil
	}

	// Create message using go-mail
	m := mail_sender.NewMessage()
	m.SetHeader("From", config.FromEmail)
	m.SetHeader("To", config.ToEmail)
	m.SetHeader("Subject", config.Subject)
	m.SetBody("text/html", body)

	// Create dialer (no authentication for EPFL mail server)
	d := mail_sender.NewDialer(config.SMTPHost, config.SMTPPort, "", "")
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Send email
	err = d.DialAndSend(m)
	if err != nil {
		return fmt.Errorf("failed to send email: %v", err)
	}

	log.Println("Email alert sent successfully")

	return nil
}
