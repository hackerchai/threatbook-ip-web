package schema

import (
	"github.com/hackerchai/threatbook-ip-web/ent"
)

type ThreatsPage struct {
	List       []*ent.Threat       `json:"list"`
	Pagination *PaginationResponse `json:"pagnition"`
}

type ThreatsResponse struct {
	Code       int         `json:"code"`
	Message    string      `json:"message"`
	StatusCode int         `json:"status_code"`
	Data       ThreatsPage `json:"data"`
}
