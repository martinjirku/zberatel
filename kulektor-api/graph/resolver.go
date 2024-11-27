package graph

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"jirku.sk/kulektor/db"
)

type Resolver struct {
	Queries *db.Queries
	Pool    *pgxpool.Pool
}
