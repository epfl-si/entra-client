package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"
	"time"

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

	// Initialize Entra client
	client, err := httpengine.New()
	if err != nil {
		log.Fatalf("Failed to initialize Entra client: %v", err)
	}

	fmt.Println("Entra client initialized successfully")

	// Calculate date 30 days from now
	futureDate := time.Now().AddDate(0, 0, int(config.ExpiryThresholdDays))
	dateLimit := futureDate.Format("2006-01-02")

	log.Printf("Checking for secrets expiring before: %s", dateLimit)

	// Query for key credentials expiring within 30 days
	clientOptions := models.ClientOptions{
		Select: "keyCredentials,displayName,appId",
	}

	keyCredentialsMap, err := client.GetKeyCredentials(dateLimit, clientOptions)
	if err != nil {
		log.Fatalf("Failed to get key credentials: %v", err)
	}

	// Query for password credentials expiring within 30 days
	passwordClientOptions := models.ClientOptions{
		Select: "passwordCredentials,displayName,appId",
	}

	passwordCredentialsMap, err := client.GetPasswordCredentials(dateLimit, passwordClientOptions)
	if err != nil {
		log.Fatalf("Failed to get password credentials: %v", err)
	}

	// Filter secrets expiring within 30 days (RemainingDays <= Threshold and >= -1)
	alertData := processCredentials(config.ExpiryThresholdDays, keyCredentialsMap, passwordCredentialsMap)

	if alertData.TotalSecrets == 0 {
		log.Println("No secrets expiring within %d days found", config.ExpiryThresholdDays)
		return
	}

	alertData.ThresholdDays = config.ExpiryThresholdDays
	log.Printf("Found %d applications with %d expiring secrets withing %d days", alertData.TotalApplications, alertData.TotalSecrets, alertData.ThresholdDays)

	err = sendEmailAlert(config, alertData)
	if err != nil {
		log.Fatalf("Failed to send email alert: %v", err)
	}

	log.Println("Email alert sent successfully")
}

func loadConfig() (*Config, error) {
	config := &Config{}

	config.SMTPHost = getEnvOrDefault("SMTP_HOST", "mail.epfl.ch")
	config.SMTPPort = 25 // Fixed port for EPFL mail server
	config.FromEmail = getEnvOrDefault("FROM_EMAIL", "noreply@epfl.ch")
	config.ToEmail = getEnvOrDefault("TO_EMAIL", "arnaud.assad@epfl.ch")

	var err error
	config.ExpiryThresholdDays, err = strconv.ParseInt(getEnvOrDefault("EXPIRY_THRESHOLD_DAYS", "30"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("invalid EXPIRY_THRESHOLD_DAYS: %v", err)
	}

	config.Subject = getEnvOrDefault("EMAIL_SUBJECT", "Entra Secrets Expiring Soon")

	return config, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	fmt.Printf("Checking environment variable: %s\n", key)
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// processCredentials processes both key and password credentials and filters for expiring ones
func processCredentials(expiryThresholdDays int64, keyCredentialsMap map[string][]models.KeyCredentialEPFL, passwordCredentialsMap map[string][]models.PasswordCredentialEPFL) *AlertData {
	alertData := &AlertData{
		Applications: make([]ApplicationAlert, 0),
		GeneratedAt:  time.Now(),
	}

	// Create a map to merge results from both key and password credentials
	allApplications := make(map[string]*ApplicationAlert)

	// Process key credentials
	for appID, credentials := range keyCredentialsMap {
		var expiringSecrets []SecretAlert
		var appDisplayName string

		for _, cred := range credentials {
			// Calculate actual remaining days from today
			actualRemainingDays := int64(time.Time(*cred.EndDateTime).Sub(time.Now()).Hours() / 24)

			// Only include secrets expiring within threshold (including recently expired ones up to 7 days)
			if actualRemainingDays <= expiryThresholdDays && actualRemainingDays >= -7 {
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

	// Process password credentials
	for appID, credentials := range passwordCredentialsMap {
		var expiringSecrets []SecretAlert
		var appDisplayName string

		for _, cred := range credentials {
			// Calculate actual remaining days from today
			actualRemainingDays := int64(cred.EndDateTime.Sub(time.Now()).Hours() / 24)

			// Only include secrets expiring within threshold (including recently expired ones up to 7 days)
			if actualRemainingDays <= expiryThresholdDays && actualRemainingDays >= -7 {
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

	// Convert map to slice and count totals
	for _, appAlert := range allApplications {
		alertData.Applications = append(alertData.Applications, *appAlert)
		alertData.TotalSecrets += len(appAlert.Secrets)
	}

	alertData.TotalApplications = len(alertData.Applications)
	return alertData
}

func sendEmailAlert(config *Config, alertData *AlertData) error {
	// Parse and execute template
	tmpl, err := template.New("email").Parse(emailTemplate)
	if err != nil {
		return fmt.Errorf("failed to parse email template: %v", err)
	}

	// Execute template to get HTML body
	emailBuffer := &EmailBuffer{}
	err = tmpl.Execute(emailBuffer, alertData)
	if err != nil {
		return fmt.Errorf("failed to execute email template: %v", err)
	}

	body := string(emailBuffer.Bytes())
	fmt.Printf("Generated email body:\n%s\n", body) // Debug output
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

	return nil
}

// EmailBuffer is a simple buffer that implements io.Writer
type EmailBuffer struct {
	data []byte
}

func (b *EmailBuffer) Write(p []byte) (n int, err error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func (b *EmailBuffer) Bytes() []byte {
	return b.data
}
