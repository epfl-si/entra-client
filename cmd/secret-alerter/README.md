# Entra Secret Alerter

A monitoring tool that scans Microsoft Entra (Azure AD) applications for expiring secrets and certificates, sending automated email alerts to administrators.

## Purpose

The Secret Alerter helps organizations maintain security by:
- Monitoring application secrets and certificates in Microsoft Entra
- Identifying secrets approaching expiration within a configurable threshold
- Sending HTML-formatted email alerts with detailed expiration information
- Preventing service disruptions caused by expired credentials

## Features

- **Configurable Threshold**: Set custom expiration warning periods (default: 30 days)
- **Rich Email Alerts**: HTML-formatted emails with color-coded urgency levels
- **Multiple Applications**: Monitors all applications accessible to the configured service principal
- **Expired Secret Detection**: Also alerts on recently expired secrets (up to 7 days past expiration)
- **Detailed Reporting**: Includes application names, secret IDs, expiration dates, and remaining days

## Alert Categories

The tool categorizes secrets based on remaining days:
- **EXPIRED**: Already expired (red, critical)
- **CRITICAL**: Less than 7 days remaining (red)
- **WARNING**: 7-14 days remaining (orange)
- **NORMAL**: More than 14 days remaining (yellow)

## Environment Variables

### Required Variables

| Variable | Description | Example |
|----------|-------------|---------|
| `ENTRA_TENANT` | Azure tenant ID | `f6c2556a-c4fb-4ab1-a2c7-9e220df11c43` |
| `ENTRA_CLIENTID` | Application/Client ID with required permissions | `593ef0c8b286ddfc` |
| `ENTRA_SECRET` | Client secret for authentication | `j5A8QWNKEXzXmTKScUA` |

### Optional Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `TO_EMAIL` | Email address to receive alerts | `arnaud.assad@epfl.ch` |
| `FROM_EMAIL` | Sender email address | `noreply@epfl.ch` |
| `SMTP_HOST` | SMTP server hostname | `mail.epfl.ch` |
| `EMAIL_SUBJECT` | Email subject line | `"Entra Secrets Expiring Soon"` |
| `EXPIRY_THRESHOLD_DAYS` | Days before expiration to trigger alerts | `30` |

## Prerequisites

### Required Entra Permissions

The configured application must have the following Microsoft Graph API permissions:
- `Application.Read.All` - To read application registrations
- `Directory.Read.All` - To read directory objects

### SMTP Configuration

The tool is configured for EPFL's internal mail server by default:
- Uses port 25 (no authentication required)
- TLS certificate verification is disabled for internal servers

## Usage

1. **Set up environment variables**:
   ```bash
   cp env.sample .env
   # Edit .env with your configuration
   ```

2. **Run the alerter**:
   ```bash
   ./secret-alerter
   ```

3. **Example output**:
   ```
   Entra client initialized successfully
   Checking for secrets expiring before: 2024-08-23
   Found 2 applications with 3 expiring secrets within 30 days
   Email alert sent successfully
   ```

## Email Format

The generated email includes:
- **Summary**: Total applications and secrets requiring attention
- **Application Details**: For each affected application:
  - Application display name and ID
  - List of expiring secrets with:
    - Secret display name and key ID
    - Exact expiration date
    - Days remaining (color-coded by urgency)

## Security Considerations

- Store sensitive environment variables securely
- Use dedicated service accounts with minimal required permissions
- Regularly rotate the client secret used for authentication
- Monitor email delivery to ensure alerts are received

## Build and Deployment

```bash
# Build the binary
go build -o secret-alerter main.go

# Run directly
go run main.go
```

## Dependencies

- `github.com/epfl-si/entra-client/pkg/client/httpengine` - Entra client library
- `github.com/go-mail/mail` - SMTP email sending