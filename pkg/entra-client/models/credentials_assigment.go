// Package models provides the models for the application
package models

type credentialsAssignment struct {
	KeyCredentials      []KeyCredential      `json:"keyCredentials"`
	PasswordCredentials []PasswordCredential `json:"passwordCredentials"`
}
