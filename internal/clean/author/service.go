package author

import (
	"context"

	"github.com/tanveerprottoy/architectures-go/internal/clean/author/dto"
)

type Service struct {
	repo Repository
}

func NewService(r Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) Create(ctx context.Context, dto dto.CreateUserDTO, args ...any) (Author, error) {
	var a Author
	return a, nil
}

func (s *Service) ReadMany(ctx context.Context, args ...any) ([]Author, error) {
	var arr []Author
	return arr, nil
}

func (s *Service) ReadOne(ctx context.Context, id string, args ...any) (Author, error) {
	var a Author
	return a, nil
}

func (s *Service) Delete(ctx context.Context, id string, args ...any) (Author, error) {
	var a Author
	return a, nil
}
