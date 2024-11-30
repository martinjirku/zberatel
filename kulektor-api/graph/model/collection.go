package model

import "jirku.sk/kulektor/db"

func CollectionFromDb(c db.Collection) Collection {
	return Collection{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		Variant:     CollectionVariantNormal,
		Type:        c.Type,
		CreatedAt:   c.CreatedAt.Time,
		UpdatedAt:   c.UpdatedAt.Time,
	}
}
