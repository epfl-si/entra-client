// Package models provides the models for the application
package models

import "time"

// PasswordCredential struct used to assign password credentials to an application
type PasswordCredential struct {
	CustomKeyIdentifier string    `json:"customKeyIdentifier"`
	KeyID               string    `json:"keyId"`
	EndDateTime         time.Time `json:"endDateTime"`
	StartDateTime       time.Time `json:"startDateTime"`
	SecretText          string    `json:"secretText"`
}
