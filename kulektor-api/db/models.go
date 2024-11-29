// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
	"jirku.sk/kulektor/ksuid"
)

type Collection struct {
	ID          ksuid.KSUID      `db:"id" json:"id"`
	UserID      string           `db:"user_id" json:"userId"`
	Title       string           `db:"title" json:"title"`
	Description *string          `db:"description" json:"description"`
	Type        *string          `db:"type" json:"type"`
	CreatedAt   pgtype.Timestamp `db:"created_at" json:"createdAt"`
	UpdatedAt   pgtype.Timestamp `db:"updated_at" json:"updatedAt"`
	BlueprintID interface{}      `db:"blueprint_id" json:"blueprintId"`
	IsBlueprint bool             `db:"is_blueprint" json:"isBlueprint"`
}

type User struct {
	ID        string           `db:"id" json:"id"`
	Username  string           `db:"username" json:"username"`
	Email     string           `db:"email" json:"email"`
	CreatedAt pgtype.Timestamp `db:"created_at" json:"createdAt"`
	UpdatedAt pgtype.Timestamp `db:"updated_at" json:"updatedAt"`
}
