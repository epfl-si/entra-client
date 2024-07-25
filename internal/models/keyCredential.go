// Package models provides the models for the application
package models

import "time"

// KeyCredential struct used to assign key credentials to an application
type KeyCredential struct {
	KeyIdentifier string    `json:"customKeyIdentifier"`
	KeyId         string    `json:"keyId"`
	EndDateTime   time.Time `json:"endDateTime"`
	StartDateTime time.Time `json:"startDateTime"`
	Type          string    `json:"type"`
	Usage         string    `json:"usage"`
	Key           string    `json:"key"`
	DisplayName   string    `json:"displayName"`
}
