// Package models provides the models for the application
package models

// ClientOptions represents the options passed from command line to the client
type ClientOptions struct {
	Batch     string
	Debug     bool
	Default   bool
	Filter    string
	Paging    bool
	Search    string
	Select    string
	Skip      string
	SkipToken string
	Top       string
}
