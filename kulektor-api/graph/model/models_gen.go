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

type Collection struct {
	ID          ksuid.KSUID       `json:"id"`
	Title       string            `json:"title"`
	Description *string           `json:"description,omitempty"`
	Type        *string           `json:"type,omitempty"`
	Variant     CollectionVariant `json:"variant"`
	CreatedAt   time.Time         `json:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt"`
}

type CollectionInput struct {
	Title       string             `json:"title"`
	Description *string            `json:"description,omitempty"`
	Type        *string            `json:"type,omitempty"`
	Variant     *CollectionVariant `json:"variant,omitempty"`
}

type CollectionsListInput struct {
	Paging *grid.Paging `json:"paging"`
}

type CollectionsListResp struct {
	Items []Collection `json:"items"`
	Meta  *Meta        `json:"meta"`
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

type CollectionField string

const (
	CollectionFieldTitle       CollectionField = "title"
	CollectionFieldDescription CollectionField = "description"
	CollectionFieldType        CollectionField = "type"
	CollectionFieldVariant     CollectionField = "variant"
)

var AllCollectionField = []CollectionField{
	CollectionFieldTitle,
	CollectionFieldDescription,
	CollectionFieldType,
	CollectionFieldVariant,
}

func (e CollectionField) IsValid() bool {
	switch e {
	case CollectionFieldTitle, CollectionFieldDescription, CollectionFieldType, CollectionFieldVariant:
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

type CollectionVariant string

const (
	CollectionVariantNormal    CollectionVariant = "NORMAL"
	CollectionVariantBlueprint CollectionVariant = "BLUEPRINT"
)

var AllCollectionVariant = []CollectionVariant{
	CollectionVariantNormal,
	CollectionVariantBlueprint,
}

func (e CollectionVariant) IsValid() bool {
	switch e {
	case CollectionVariantNormal, CollectionVariantBlueprint:
		return true
	}
	return false
}

func (e CollectionVariant) String() string {
	return string(e)
}

func (e *CollectionVariant) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CollectionVariant(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CollectionVariant", str)
	}
	return nil
}

func (e CollectionVariant) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
