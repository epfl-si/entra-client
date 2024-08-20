// Package models provides the models for the application
package models

import "time"

// PasswordCredential struct used to assign password credentials to an application
type PasswordCredential struct {
	CustomKeyIdentifier string    `json:"customKeyIdentifier,ommitEmpty"`
	KeyID               string    `json:"keyId,ommitEmpty"`
	DisplayName         string    `json:"displayName,ommitEmpty"`
	Hint                string    `json:"hint,ommitEmpty"`
	EndDateTime         time.Time `json:"endDateTime,ommitEmpty"`
	StartDateTime       time.Time `json:"startDateTime,ommitEmpty"`
	SecretText          *string   `json:"secretText,ommitEmpty"`
}
