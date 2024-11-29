package db

import (
	"context"
	"fmt"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func GetUpdateCollectionQuery(c Collection, fields []string) (string, []any, error) {
	sets := make([]bob.Expression, 0, len(fields))
	for _, f := range fields {
		col, val, err := GetDbColumnByJsonField(c, f)
		if err != nil {
			return "", nil, fmt.Errorf("building query: %s", err)
		}
		sets = append(sets, psql.Quote(col).EQ(psql.Arg(val)))

	}
	insert, args, err := psql.Update(
		um.Table("collections"),
		um.Set(sets...),
		um.Where(
			psql.Quote("user_id").EQ(psql.Arg(c.UserID)).And(
				psql.Quote("id").EQ(psql.Arg(c.ID))),
		),
		um.Returning(psql.RawQuery("*")),
	).Build(context.Background())

	return insert, args, err
}

func (q *Queries) UpdateMyCollection(ctx context.Context, c Collection, fields []string) (Collection, error) {
	query, args, err := GetUpdateCollectionQuery(c, fields)
	if err != nil {
		return Collection{}, fmt.Errorf("creating query: %s", err)
	}
	row := q.db.QueryRow(ctx, query, args)

	var i Collection
	err = row.Scan(
		&i.ID,
		&i.UserID,
		&i.Title,
		&i.Description,
		&i.Type,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.BlueprintID,
		&i.IsBlueprint,
	)
	if err != nil {
		fmt.Printf(">>>>> query:")
		fmt.Printf(query)
		return i, fmt.Errorf("calling query %q: %s", query, err)
	}
	return i, err
}
