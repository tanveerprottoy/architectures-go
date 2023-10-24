package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/middlewareext"
)

// Router struct
type Router struct {
	Mux *chi.Mux
}

func NewRouter() *Router {
	r := &Router{}
	r.Mux = chi.NewRouter()
	r.registerGlobalMiddlewares()
	return r
}

func (r *Router) registerGlobalMiddlewares() {
	r.Mux.Use(
		middleware.Logger,
		// middleware.Recoverer,
		middlewareext.JSONContentTypeMiddleWare,
		middlewareext.CORSEnableMiddleWare,
	)
}
