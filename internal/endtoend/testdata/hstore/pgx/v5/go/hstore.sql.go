// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.17.0
// source: hstore.sql

package hstore

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const listBar = `-- name: ListBar :many
SELECT bar FROM foo
`

func (q *Queries) ListBar(ctx context.Context) ([]pgtype.Hstore, error) {
	rows, err := q.db.Query(ctx, listBar)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Hstore
	for rows.Next() {
		var bar pgtype.Hstore
		if err := rows.Scan(&bar); err != nil {
			return nil, err
		}
		items = append(items, bar)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listBaz = `-- name: ListBaz :many
SELECT baz FROM foo
`

func (q *Queries) ListBaz(ctx context.Context) ([]pgtype.Hstore, error) {
	rows, err := q.db.Query(ctx, listBaz)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []pgtype.Hstore
	for rows.Next() {
		var baz pgtype.Hstore
		if err := rows.Scan(&baz); err != nil {
			return nil, err
		}
		items = append(items, baz)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
