package model

import (
	"net/url"
	"strconv"
)

type PagingRequest struct {
	Start int64 `json:"page" validate:"required,min=0"`
	Take  int64 `json:"perPage" validate:"required,min=0,max=100"`
}

func (p *PagingRequest) FromQuery(u url.Values) error {
	if u.Has("start") {
		start := u.Get("start")
		parsedStart, err := strconv.Atoi(start)
		if err != nil {
			return err
		}
		p.Start = int64(parsedStart)
	}

	if u.Has("take") {
		take := u.Get("take")
		parsedTake, err := strconv.Atoi(take)
		if err != nil {
			return err
		}
		p.Take = int64(parsedTake)
	}
	return nil
}

type PagingResponse[T any] struct {
	Total int64 `json:"total"`
	Items []T   `json:"items"`
}
