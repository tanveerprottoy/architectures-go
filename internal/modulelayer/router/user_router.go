package router

import (
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/user"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/constant"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/router"

	"github.com/go-chi/chi"
)

func RegisterUserRoutes(router *router.Router, versions []string, module *user.Module) {
	router.Mux.Route(
		constant.ApiPattern+versions[0]+constant.UsersPattern,
		func(r chi.Router) {
			// public routes
			r.Get(constant.RootPattern+"public", module.Handler.Public)
			r.Group(func(r chi.Router) {
				r.Get(constant.RootPattern, module.Handler.ReadMany)
				r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
				r.Post(constant.RootPattern, module.Handler.Create)
				r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
				r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
			})
		},
	)
}
