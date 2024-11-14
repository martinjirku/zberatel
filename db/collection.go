package db

import (
	"jirku.sk/zberatel/model"
)

func (c *Collection) ToModel() model.Collection {
	collection := model.Collection{
		ID:     c.ID,
		Title:  c.Title,
		Type:   c.Type,
		UserID: c.UserID,
	}
	if c.Description.Valid {
		collection.Description = c.Description.String
	}
	return collection
}
