// Package models provides the models for the application
package models

import "time"

// KeyCredential struct used to assign key credentials to an application
type KeyCredential struct {
	CustomKeyIdentifier string     `json:"customKeyIdentifier,omitempty"`
	KeyID               string     `json:"keyId,omitempty"`
	EndDateTime         *time.Time `json:"endDateTime,omitempty"`
	StartDateTime       *time.Time `json:"startDateTime,omitempty"`
	Type                string     `json:"type,omitempty"`
	Usage               string     `json:"usage,omitempty"`
	// Key                 []byte     `json:"key,omitempty"`
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}
