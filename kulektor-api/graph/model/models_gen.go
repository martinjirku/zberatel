// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"

	"jirku.sk/kulektor/auth"
	"jirku.sk/kulektor/grid"
	"jirku.sk/kulektor/ksuid"
)

type Blueprint struct {
	ID          ksuid.KSUID `json:"id"`
	Title       string      `json:"title"`
	UserID      *string     `json:"userId,omitempty"`
	Description *string     `json:"description,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type BlueprintInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}

type BlueprintsListInput struct {
	Paging *grid.Paging `json:"paging"`
}

type BlueprintsListResp struct {
	Items []Blueprint `json:"items"`
	Meta  *Meta       `json:"meta"`
}

type Collection struct {
	ID          ksuid.KSUID `json:"id"`
	Title       string      `json:"title"`
	Description *string     `json:"description,omitempty"`
	Type        *string     `json:"type,omitempty"`
	CreatedAt   time.Time   `json:"createdAt"`
	UpdatedAt   time.Time   `json:"updatedAt"`
}

type CollectionInput struct {
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
	Type        *string `json:"type,omitempty"`
}

type CollectionsListInput struct {
	Paging *grid.Paging `json:"paging"`
}

type CollectionsListResp struct {
	Items []Collection `json:"items"`
	Meta  *Meta        `json:"meta"`
}

type CreateBlueprintResp struct {
	Success bool       `json:"success"`
	Data    *Blueprint `json:"data,omitempty"`
}

type CreateCollectionResp struct {
	Success bool        `json:"success"`
	Data    *Collection `json:"data,omitempty"`
}

type DeleteMyCollectionResp struct {
	Success bool `json:"success"`
}

type Meta struct {
	Total       int          `json:"total"`
	PrevPage    *grid.Paging `json:"prevPage,omitempty"`
	CurrentPage *grid.Paging `json:"currentPage"`
	NextPage    *grid.Paging `json:"nextPage,omitempty"`
}

type Mutation struct {
}

type Query struct {
}

type UpdateBlueprintInput struct {
	ID             ksuid.KSUID      `json:"id"`
	Blueprint      *BlueprintInput  `json:"blueprint"`
	FieldsToUpdate []BlueprintField `json:"fieldsToUpdate"`
}

type UpdateBlueprintResp struct {
	Success bool       `json:"success"`
	Data    *Blueprint `json:"data,omitempty"`
}

type UpdateCollectionInput struct {
	ID             ksuid.KSUID       `json:"id"`
	Collection     *CollectionInput  `json:"collection"`
	FieldsToUpdate []CollectionField `json:"fieldsToUpdate"`
}

type UpdateCollectionResp struct {
	Success bool        `json:"success"`
	Data    *Collection `json:"data,omitempty"`
}

type User struct {
	UID   string    `json:"uid"`
	Email string    `json:"email"`
	Role  auth.Role `json:"role"`
}

type BlueprintField string

const (
	BlueprintFieldTitle       BlueprintField = "title"
	BlueprintFieldDescription BlueprintField = "description"
)

var AllBlueprintField = []BlueprintField{
	BlueprintFieldTitle,
	BlueprintFieldDescription,
}

func (e BlueprintField) IsValid() bool {
	switch e {
	case BlueprintFieldTitle, BlueprintFieldDescription:
		return true
	}
	return false
}

func (e BlueprintField) String() string {
	return string(e)
}

func (e *BlueprintField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = BlueprintField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid BlueprintField", str)
	}
	return nil
}

func (e BlueprintField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type CollectionField string

const (
	CollectionFieldTitle       CollectionField = "title"
	CollectionFieldDescription CollectionField = "description"
	CollectionFieldType        CollectionField = "type"
)

var AllCollectionField = []CollectionField{
	CollectionFieldTitle,
	CollectionFieldDescription,
	CollectionFieldType,
}

func (e CollectionField) IsValid() bool {
	switch e {
	case CollectionFieldTitle, CollectionFieldDescription, CollectionFieldType:
		return true
	}
	return false
}

func (e CollectionField) String() string {
	return string(e)
}

func (e *CollectionField) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollectionField(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollectionField", str)
	}
	return nil
}

func (e CollectionField) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
