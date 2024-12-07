package db

import (
	"context"
	"fmt"
	"time"

	"github.com/stephenafamo/bob"
	"github.com/stephenafamo/bob/dialect/psql"
	"github.com/stephenafamo/bob/dialect/psql/dialect"
	"github.com/stephenafamo/bob/dialect/psql/sm"
	"github.com/stephenafamo/bob/dialect/psql/um"
)

func GetUpdateBlueprintQuery(ctx context.Context, c Blueprint, fields []string) (string, []any, error) {
	baseQuery := psql.Update(
		um.Table("blueprints"),
		um.SetCol("updated_at").ToArg(time.Now()),
		um.Where(psql.Quote("id").EQ(psql.Arg(c.ID))),
		um.Returning(psql.RawQuery("id, user_id, title, description, type, created_at, updated_at")),
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

func (q *Queries) UpdateBlueprint(ctx context.Context, c Blueprint, fields []string) (Blueprint, error) {
	query, args, err := GetUpdateBlueprintQuery(ctx, c, fields)
	if err != nil {
		return Blueprint{}, fmt.Errorf("creating query: %s", err)
	}
	row := q.db.QueryRow(ctx, query, args...)

	i := Blueprint{}
	err = row.Scan(&i)
	if err != nil {
		return i, fmt.Errorf("calling query %q: %s", query, err)
	}
	return i, err
}

func FilterToExpression(filters []Filter) dialect.Expression {
	conditions := make([]bob.Expression, 0)
	for _, f := range filters {
		if f.Type == FilterTypeEQ {
			conditions = append(conditions, psql.Quote(f.Column).EQ(psql.Arg(f.Value)))
		}
	}
	return psql.And(conditions...)
}

func GetListBlueprintCountQuery(ctx context.Context, filters []Filter) (string, []any, error) {
	baseQuery := psql.Select(
		sm.From("blueprints"),
		sm.Columns(psql.Raw("count(*)").As("Counts")),
	)
	if len(filters) > 0 {
		filterExpression := FilterToExpression(filters)
		baseQuery.Apply(sm.Where(filterExpression))
	}
	return baseQuery.Build(ctx)
}

func GetListBlueprintQuery(ctx context.Context, filters []Filter, offset, limit int) (string, []any, error) {
	baseQuery := psql.Select(
		sm.From("blueprints"),
		sm.Columns(psql.RawQuery("id, user_id, title, description, created_at, updated_at")),
		sm.Offset(psql.Arg(offset)),
		sm.Limit(psql.Arg(limit)),
	)
	if len(filters) > 0 {
		filterExpression := FilterToExpression(filters)
		baseQuery.Apply(sm.Where(filterExpression))
	}
	return baseQuery.Build(ctx)
}

func (q *Queries) ListBlueprints(ctx context.Context, filters []Filter, offset, limit int) ([]Blueprint, error) {
	query, args, err := GetListBlueprintQuery(ctx, filters, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("getting blueprints list count: %s", err)
	}
	rows, err := q.db.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query blueprints: %s", err)
	}
	results := make([]Blueprint, 0)
	for rows.Next() {
		b := Blueprint{}
		err := rows.Scan(&b)
		if err != nil {
			return nil, fmt.Errorf("scanning blueprint: %s", err)
		}
		results = append(results, b)
	}
	return results, nil
}

func (q *Queries) ListBlueprintsCount(ctx context.Context, filters []Filter) (int64, error) {
	query, args, err := GetListBlueprintCountQuery(ctx, filters)
	if err != nil {
		return 0, fmt.Errorf("getting blueprints list count: %s", err)
	}
	row := q.db.QueryRow(ctx, query, args...)
	var i int64
	errScan := row.Scan(&i)
	if errScan != nil {
		return i, fmt.Errorf("scanning counts: %s", errScan)
	}
	return i, nil

}
