package models

import "time"

// ServiceProvisioningXMLError represents a service provisioning XML error (in Group description)
type ServiceProvisioningXMLError struct {
	RelayState      string     `json:"relayState,omitempty"`
	CreatedDateTime *time.Time `json:"createdDateTime"`
	ErrorDetail     string     `json:"errorDetail"`
	IsResolved      bool       `json:"isResolved"`
	ServiceInstance string     `json:"serviceInstance"`
}
