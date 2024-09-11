package common

type successResponse struct {
	Data   interface{} `json:"data"`
	Paging interface{} `json:"paging,omitempty"`
	Filter interface{} `json:"filter,omitempty"`
}

func NewSuccessRes(data, paging, filter interface{}) *successResponse {
	return &successResponse{Data: data, Paging: paging, Filter: filter}
}

func SimpleSuccessRes(data interface{}) *successResponse {
	return &successResponse{data, nil, nil}
}
