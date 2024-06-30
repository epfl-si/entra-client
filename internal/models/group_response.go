package models

type GroupResponse struct {
	Context  string   `json:"@odata.context"`
	NextLink string   `json:"@odata.nextLink"`
	Value    []*Group `json:"value"`
}
