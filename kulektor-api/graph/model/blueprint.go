package model

import "jirku.sk/kulektor/db"

func BlueprintFromDb(c db.Blueprint) Blueprint {
	return Blueprint{
		ID:          c.ID,
		Title:       c.Title,
		Description: c.Description,
		CreatedAt:   c.CreatedAt.Time,
		UpdatedAt:   c.UpdatedAt.Time,
	}
}
