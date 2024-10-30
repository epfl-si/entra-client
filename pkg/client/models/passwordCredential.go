// Package models provides the models for the application
package models

import "time"

// PasswordCredential struct used to assign password credentials to an application
type PasswordCredential struct {
	CustomKeyIdentifier string    `json:"customKeyIdentifier,omitempty"`
	KeyID               string    `json:"keyId,omitempty"`
	DisplayName         string    `json:"displayName,omitempty"`
	Hint                string    `json:"hint,omitempty"`
	EndDateTime         time.Time `json:"endDateTime,omitempty"`
	StartDateTime       time.Time `json:"startDateTime,omitempty"`
	SecretText          *string   `json:"secretText,omitempty"`
}
