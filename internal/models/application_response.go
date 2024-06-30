package models

type ApplicationResponse struct {
	Context  string         `json:"@odata.context"`
	NextLink string         `json:"@odata.nextLink"`
	Value    []*Application `json:"value"`
}
