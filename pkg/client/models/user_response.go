// Package models provides the models for the application
package models

type UserResponse struct {
	Context  string  `json:"@odata.context"`
	NextLink string  `json:"@odata.nextLink"`
	Value    []*User `json:"value"`
}
