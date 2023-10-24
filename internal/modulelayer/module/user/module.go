package user

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/user/entity"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository data.Repository[entity.User]
}

func NewModule(db *sql.DB, validate *validator.Validate) *Module {
	m := new(Module)
	// init order is reversed of the field decleration
	// as the dependency is served this way
	m.Repository = NewRepository(db)
	m.Service = NewService(m.Repository)
	m.Handler = NewHandler(m.Service, validate)
	return m
}
