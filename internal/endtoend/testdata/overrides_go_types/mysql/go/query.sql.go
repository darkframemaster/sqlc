// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: query.sql

package override

import (
	"context"
	"strings"

	"github.com/darkframemaster/sqlc-testdata/pkg"
)

const testIN = `-- name: TestIN :many
SELECT other, total, retyped FROM foo WHERE retyped IN (/*SLICE:paramname*/?)
`

func (q *Queries) TestIN(ctx context.Context, paramname []pkg.CustomType) ([]Foo, error) {
	query := testIN
	var queryParams []interface{}
	if len(paramname) > 0 {
		for _, v := range paramname {
			queryParams = append(queryParams, v)
		}
		query = strings.Replace(query, "/*SLICE:paramname*/?", strings.Repeat(",?", len(paramname))[1:], 1)
	} else {
		query = strings.Replace(query, "/*SLICE:paramname*/?", "NULL", 1)
	}
	rows, err := q.db.QueryContext(ctx, query, queryParams...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Foo
	for rows.Next() {
		var i Foo
		if err := rows.Scan(&i.Other, &i.Total, &i.Retyped); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
