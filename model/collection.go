package model

type CollectionInput struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Type        string `json:"type" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"min=3,max=100"`
	UserID      KSUID  `json:"userID" validate:"required"`
}

type Collection struct {
	ID          KSUID
	Title       string
	Type        string
	Description string
	UserID      KSUID
}
