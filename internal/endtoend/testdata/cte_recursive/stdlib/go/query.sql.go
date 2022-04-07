// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package querytest

import (
	"context"
	"database/sql"
)

const cTERecursive = `-- name: CTERecursive :many
WITH RECURSIVE cte AS (
        SELECT b.id, b.parent_id FROM bar AS b
        WHERE b.id = $1
    UNION ALL
        SELECT b.id, b.parent_id
        FROM bar AS b, cte AS c
        WHERE b.parent_id = c.id
) SELECT id, parent_id FROM cte
`

type CTERecursiveRow struct {
	ID       int32
	ParentID sql.NullInt32
}

func (q *Queries) CTERecursive(ctx context.Context, id int32) ([]CTERecursiveRow, error) {
	rows, err := q.db.QueryContext(ctx, cTERecursive, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CTERecursiveRow
	for rows.Next() {
		var i CTERecursiveRow
		if err := rows.Scan(&i.ID, &i.ParentID); err != nil {
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
