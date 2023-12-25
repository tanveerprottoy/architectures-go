package author

import "context"

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

	Update(ctx context.Context, id string, a Author, args ...any) (Author, error)

	Delete(ctx context.Context, id string, args ...any) (Author, error)
}

type Reader interface {
	ReadMany(ctx context.Context, args ...any) ([]Author, error)

	ReadOne(ctx context.Context, id string, args ...any) (Author, error)
}

// RepositoryCQRS defines the data persistance logic that needs to be implemented
type RepositoryCQRS interface {
	Writer
	Reader
}
