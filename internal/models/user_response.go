package models

type UserResponse struct {
	Context  string  `json:"@odata.context"`
	NextLink string  `json:"@odata.nextLink"`
	Value    []*User `json:"value"`
}
