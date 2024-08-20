// Package models provides the models for the application
package models

import (
	"encoding/json"
	"time"
)

// KeyCredential struct used to assign key credentials to an application
type KeyCredential struct {
	CustomKeyIdentifier string `json:"customKeyIdentifier,omitempty"`
	KeyID               string `json:"keyId,omitempty"`
	// EndDateTime         *time.Time `json:"endDateTime,omitempty"`
	// StartDateTime       *time.Time `json:"startDateTime,omitempty"`
	EndDateTime   *CustomTime `json:"endDateTime,omitempty"`
	StartDateTime *CustomTime `json:"startDateTime,omitempty"`
	Type          string      `json:"type,omitempty"`
	Usage         string      `json:"usage,omitempty"`
	// Key                 []byte     `json:"key,omitempty"`
	Key         string `json:"key,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
}

// CustomTime is a wrapper around time.Time
type CustomTime time.Time

// MarshalJSON customizes the JSON encoding of CustomTime
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	t := time.Time(ct)
	formattedTime := t.Format("2006-01-02T15:04:05Z")
	return json.Marshal(formattedTime)
}

// UnmarshalJSON customizes the JSON decoding of CustomTime
func (ct *CustomTime) UnmarshalJSON(b []byte) error {
	var formattedTime string
	if err := json.Unmarshal(b, &formattedTime); err != nil {
		return err
	}
	t, err := time.Parse("2006-01-02T15:04:05Z", formattedTime)
	if err != nil {
		return err
	}
	*ct = CustomTime(t)
	return nil
}
