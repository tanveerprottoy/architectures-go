package content

import (
	"database/sql"

	"github.com/go-playground/validator/v10"
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/content/entity"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/data"
)

type Module struct {
	Handler    *Handler
	Service    *Service
	Repository data.Repository[entity.Content]
}

func NewModule(db *sql.DB, v *validator.Validate) *Module {
	// init order is reversed of the field decleration
	// as the dependency is served this way
	r := NewRepository(db)
	s := NewService(r)
	h := NewHandler(s, v)
	return &Module{Handler: h, Service: s, Repository: r}
}
