package db

import (
	"context"
	"fmt"
	"time"

	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func GetUpdateCollectionQuery(ctx context.Context, c Collection, fields []string) (string, []any, error) {
	baseQuery := psql.Update(
		um.Table("collections"),
		um.SetCol("updated_at").ToArg(time.Now()),
		um.Where(
			psql.Quote("user_id").EQ(psql.Arg(c.UserID)).And(
				psql.Quote("id").EQ(psql.Arg(c.ID))),
		),
		um.Returning(psql.RawQuery("id, user_id, title, description, type, created_at, updated_at, blueprint_id, is_blueprint")),
	)
	for _, f := range fields {
		col, val, err := GetDbColumnByJsonField(c, f)
		if err != nil {
			return "", nil, fmt.Errorf("building query: %s", err)
		}
		baseQuery.Apply(um.SetCol(col).ToArg(val))
	}
	return baseQuery.Build(ctx)
}

func (q *Queries) UpdateMyCollection(ctx context.Context, c Collection, fields []string) (Collection, error) {
	query, args, err := GetUpdateCollectionQuery(ctx, c, fields)
	if err != nil {
		return Collection{}, fmt.Errorf("creating query: %s", err)
	}
	row := q.db.QueryRow(ctx, query, args...)

	var i Collection
	err = row.Scan(&i)
	if err != nil {
		return i, fmt.Errorf("calling query %q: %s", query, err)
	}
	return i, err
}
