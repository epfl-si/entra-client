// Package models provides the models for the application
package models

import "time"

// KeyCredential struct used to assign key credentials to an application
type KeyCredential struct {
	KeyIdentifier string     `json:"customKeyIdentifier,omitempty"`
	KeyId         string     `json:"keyId,omitempty"`
	EndDateTime   *time.Time `json:"endDateTime,omitempty"`
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	Type          string     `json:"type,omitempty"`
	Usage         string     `json:"usage,omitempty"`
	Key           string     `json:"key,omitempty"`
	DisplayName   string     `json:"displayName,omitempty"`
}
