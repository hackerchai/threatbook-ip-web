package schema

type StatusText string

func (t StatusText) String() string {
	return string(t)
}

const (
	OKStatus    StatusText = "success"
	ErrorStatus StatusText = "Error"
	FailStatus  StatusText = "Fail"
)

type ErrorResponse struct {
	Code       int    `json:"code"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

type StatusResult struct {
	Status StatusText `json:"status"`
}

type ErrorResult struct {
	Error ErrorItem `json:"error"`
}

type ErrorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type ListResult struct {
	List       interface{}         `json:"list"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
}

type PaginationResponse struct {
	Total    int `json:"total"`
	Current  int `json:"current"`
	PageSize int `json:"page_size"`
}

type PaginationParam struct {
	Current  int `query:"current,default=1"`
	PageSize int `query:"page_size,default=10"`
}

func (a *PaginationParam) GetCurrent() int {
	current := a.Current
	if a.Current <= 0 {
		current = 1
	}
	return current
}

func (a *PaginationParam) GetPageSize() int {
	pageSize := a.PageSize
	if a.PageSize <= 0 {
		pageSize = 10
	}
	return pageSize
}

func (r *PaginationResponse) NewPaginationResponse(total int, current int, pageSize int) PaginationResponse {
	return PaginationResponse{
		Total:    total,
		Current:  current,
		PageSize: pageSize,
	}
}
