package model

import "github.com/segmentio/ksuid"

type CollectionInput struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Type        string `json:"type" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"min=3,max=100"`
}

type Collection struct {
	ID          ksuid.KSUID
	Title       string
	Type        string
	Description string
}
