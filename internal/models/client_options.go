// Package models provides the models for the application
package models

type ClientOptions struct {
	Batch     string
	Paging    bool
	Top       string
	Search    string
	Select    string
	Skip      string
	SkipToken string
}
