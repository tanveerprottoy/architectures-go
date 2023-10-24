package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	pgxstdlib "github.com/jackc/pgx/v5/stdlib"
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/user/entity"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
)

const tableName = "users"

type Repository[T entity.User] struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository[entity.User] {
	return &Repository[entity.User]{db: db}
}

func (r *Repository[T]) Create(e entity.User, ctx context.Context) (string, error) {
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

func (r *Repository[T]) Update(id string, e entity.User, ctx context.Context) (int64, error) {
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

// createMany Batch inserts contents
func (r *Repository[T]) createMany(entities []entity.User, ctx context.Context) error {
	ctx1 := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx1, 20*time.Second)
	defer cancelFn()
	dbConn, err := r.db.Conn(ctx)
	if err != nil {
		return err
	}
	err = dbConn.Raw(func(driverConn any) error {
		if conn, ok := driverConn.(*pgxstdlib.Conn); ok {
			var rows [][]any
			for _, e := range entities {
				rows = append(rows, []any{e.Name, e.CreatedAt, e.UpdatedAt})
			}
			copyCount, err := conn.Conn().CopyFrom(
				context.Background(),
				pgx.Identifier{tableName},
				[]string{"name", "created_at", "updated_at"},
				pgx.CopyFromRows(rows),
			)
			if err != nil {
				return err
			}
			l := len(entities)
			if int(copyCount) != l {
				return fmt.Errorf("bulk insert failed, insert count: %d param count: %d", copyCount, l)
			}
			return nil
		}
		return errors.New("driver connection is not of expected type")
	})
	if err != nil {
		return err
	}
	return nil
}
