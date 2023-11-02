package domain

import (
	"context"

	"github.com/tanveerprottoy/architectures-go/internal/clean/author/dto"
)

// UseCase defines the business logic that needs to be implemented
type UseCase interface {
	Create(ctx context.Context, dto dto.CreateUserDTO, args ...any) (Author, error)
	ReadMany(ctx context.Context, args ...any) ([]Author, error)
	ReadOne(ctx context.Context, id string, args ...any) (Author, error)
	Delete(ctx context.Context, id string, args ...any) (Author, error)
}
