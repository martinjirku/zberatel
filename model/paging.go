package model

type PagingRequest struct {
	Start int64 `json:"page" validate:"required,min=0"`
	Take  int64 `json:"perPage" validate:"required,min=0,max=100"`
}

type PagingResponse[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}
