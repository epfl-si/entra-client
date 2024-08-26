package models

import "time"

// VerifiedPublisher represents a verified publisher
type VerifiedPublisher struct {
	DisplayName         string     `json:"displayName,omitempty"`
	VerifiedPublisherID string     `json:"verifiedPublisherId,omitempty"`
	AddedDateTime       *time.Time `json:"addedDateTime,omitempty"`
}
