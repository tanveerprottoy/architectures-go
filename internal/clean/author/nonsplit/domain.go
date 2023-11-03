package nonsplit

import (
	"context"
	"time"

	"github.com/tanveerprottoy/architectures-go/internal/clean/author/dto"
)

type Author struct {
	Name string    `db:"name" json:"name"`
	DOB  time.Time `db:"dob" json:"dob"`
}

// Repository defines the data persistance logic that needs to be implemented
type Repository interface {
	Create(ctx context.Context, a Author, args ...any) (string, error)
	ReadMany(ctx context.Context, args ...any) ([]Author, error)
	ReadOne(ctx context.Context, id string, args ...any) (Author, error)
	Update(ctx context.Context, id string, a Author, args ...any) (Author, error)
	Delete(ctx context.Context, id string, args ...any) (Author, error)
}

type Writer interface {
	Create(ctx context.Context, a Author, args ...any) (string, error)
}

// RepositoryCQRS defines the data persistance logic that needs to be implemented
type RepositoryCQRS interface {
	Create(ctx context.Context, a Author, args ...any) (string, error)
	ReadMany(ctx context.Context, args ...any) ([]Author, error)
	ReadOne(ctx context.Context, id string, args ...any) (Author, error)
	Delete(ctx context.Context, id string, args ...any) (Author, error)
}

// UseCase defines the business logic that needs to be implemented
type UseCase interface {
	Create(ctx context.Context, dto dto.CreateUserDTO, args ...any) (Author, error)
	ReadMany(ctx context.Context, args ...any) ([]Author, error)
	ReadOne(ctx context.Context, id string, args ...any) (Author, error)
	Update(ctx context.Context, id string, dto dto.UpdateUserDTO, args ...any) (Author, error)
	Delete(ctx context.Context, id string, args ...any) (Author, error)
}
