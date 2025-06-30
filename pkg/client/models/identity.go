package models

type OnTokenIssuanceStartListener struct {
	ODataType  string     `json:"@odata.type"`
	Conditions Conditions `json:"conditions"`
	Priority   int        `json:"priority"`
	Handler    Handler    `json:"handler"`
}

type Conditions struct {
	Applications Applications `json:"applications"`
}

type Applications struct {
	IncludeAllApplications bool                  `json:"includeAllApplications"`
	IncludeApplications    []ApplicationIdentity `json:"includeApplications"`
}

type ApplicationIdentity struct {
	AppId string `json:"appId"`
}

type Handler struct {
	ODataType       string          `json:"@odata.type"`
	CustomExtension CustomExtension `json:"customExtension"`
}

type CustomExtension struct {
	Id string `json:"id"`
}
