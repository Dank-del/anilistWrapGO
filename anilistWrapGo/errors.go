package anilistWrapGo

type Locations struct {
	Line   int `json:"line"`
	Column int `json:"column"`
}
type Errors struct {
	Message   string      `json:"message"`
	Status    int         `json:"status"`
	Locations []Locations `json:"locations"`
}
