package models

type PreAuthorizedApplication struct {
	AppID                 string   `json:"appId,omitempty"`
	DelegatePermissionIDs []string `json:"delegatePermissionIds,omitempty"`
}
