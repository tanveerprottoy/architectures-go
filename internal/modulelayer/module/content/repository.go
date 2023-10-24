package content

import (
	"context"
	"database/sql"

	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/content/entity"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
)

const tableName = "contents"

type Repository[T entity.Content] struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository[entity.Content] {
	return &Repository[entity.Content]{db: db}
}

func (r *Repository[T]) Create(e entity.Content, ctx context.Context) (string, error) {
	var lastID string
	q := data.BuildInsertQuery(tableName, []string{"name", "created_at", "updated_at"}, "RETURNING id")
	err := r.db.QueryRowContext(ctx, q, e.Name, e.CreatedAt, e.UpdatedAt).Scan(&lastID)
	if err != nil {
		return lastID, err
	}
	return lastID, nil
}

func (r *Repository[T]) ReadMany(limit, offset int, ctx context.Context) (*sql.Rows, error) {
	q := data.BuildSelectQuery(tableName, []string{}, []string{"is_deleted"}, "LIMIT $2 OFFSET $3")
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (r *Repository[T]) ReadOne(id string, ctx context.Context) *sql.Row {
	q := data.BuildSelectQuery(tableName, []string{}, []string{"id"}, "LIMIT $2")
	return r.db.QueryRow(q, id, 1)
}

func (r *Repository[T]) Update(id string, e entity.Content, ctx context.Context) (int64, error) {
	q := data.BuildUpdateQuery(tableName, []string{"name", "updated_at"}, []string{"id"}, "")
	res, err := r.db.Exec(q, e.Name, e.UpdatedAt, id)
	if err != nil {
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}

func (r *Repository[T]) Delete(id string, ctx context.Context) (int64, error) {
	q := data.BuildDeleteQuery(tableName, []string{"id"}, "")
	res, err := r.db.Exec(q, id)
	if err != nil {
		return -1, err
	}
	return data.GetRowsAffected(res), nil
}

func (r *Repository[T]) DB() *sql.DB {
	return r.db
}
