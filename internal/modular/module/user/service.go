package user

import (
	"context"
	"errors"
	"net/http"

	"github.com/tanveerprottoy/architectures-go/internal/modular/module/user/dto"
	"github.com/tanveerprottoy/architectures-go/internal/modular/module/user/entity"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/constant"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/errorext"
	"github.com/tanveerprottoy/architectures-go/pkg/timeext"
)

type Service struct {
	repository data.Repository[entity.User]
}

// NewService initializes a new ServiceSQL
func NewService(r data.Repository[entity.User]) *Service {
	return &Service{repository: r}
}

func (s *Service) readOneInternal(id string, ctx context.Context) (entity.User, errorext.HTTPError) {
	var e entity.User
	row := s.repository.ReadOne(id, ctx)
	err := row.Err()
	if err != nil {
		return e, errorext.HTTPError{Code: http.StatusInternalServerError, Err: err}
	}
	httpErr := data.ScanRow[entity.User](row, &e, &e.ID, &e.Email, &e.Name, &e.CreatedAt, &e.UpdatedAt)
	return e, httpErr
}

// Create defines the business logic for create post request
func (s *Service) Create(d dto.CreateUpdateUserDTO, ctx context.Context) (entity.User, errorext.HTTPError) {
	// build entity
	n := timeext.NowUnixMilli()
	e := entity.User{
		Name:      d.Name,
		CreatedAt: n,
		UpdatedAt: n,
	}
	l, err := s.repository.Create(e, ctx)
	if err != nil {
		return e, errorext.BuildDBError(err)
	}
	e.ID = l
	return e, errorext.HTTPError{}
}

func (s *Service) ReadMany(limit, page int, ctx context.Context) (map[string]any, errorext.HTTPError) {
	m := make(map[string]any)
	m["items"] = make([]entity.User, 0)
	m["limit"] = limit
	m["page"] = page
	offset := limit * (page - 1)
	rows, err := s.repository.ReadMany(limit, offset, ctx)
	if err != nil {
		return m, errorext.BuildDBError(err)
	}
	defer rows.Close()
	var e entity.User
	d, httpErr := data.ScanRows(rows, &e, &e.ID, &e.Name, &e.CreatedAt, &e.UpdatedAt)
	if httpErr.Err != nil {
		return m, errorext.BuildDBError(err)
	}
	m["items"] = d
	return m, errorext.HTTPError{}
}

func (s *Service) ReadOne(id string, ctx context.Context) (entity.User, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	return b, errorext.HTTPError{}
}

func (s *Service) Update(id string, d *dto.CreateUpdateUserDTO, ctx context.Context) (entity.User, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	b.Name = d.Name
	b.UpdatedAt = timeext.NowUnixMilli()
	rows, err := s.repository.Update(id, b, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errors.New(constant.OperationNotSuccess)}
}

func (s *Service) Delete(id string, ctx context.Context) (entity.User, errorext.HTTPError) {
	b, httpErr := s.readOneInternal(id, ctx)
	if httpErr.Err != nil {
		return b, errorext.BuildDBError(httpErr.Err)
	}
	rows, err := s.repository.Delete(id, ctx)
	if err != nil {
		return b, errorext.BuildDBError(err)
	}
	if rows > 0 {
		return b, errorext.HTTPError{}
	}
	return b, errorext.HTTPError{Code: http.StatusBadRequest, Err: errors.New(constant.OperationNotSuccess)}
}
