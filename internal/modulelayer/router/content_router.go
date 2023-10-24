package router

import (
	"github.com/tanveerprottoy/architectures-go/internal/modulelayer/module/content"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/constant"
	"github.com/tanveerprottoy/architectures-go/internal/pkg/router"

	"github.com/go-chi/chi"
)

func RegisterContentRoutes(router *router.Router, versions []string, module *content.Module) {
	router.Mux.Route(
		constant.ApiPattern+versions[0]+constant.ContentsPattern,
		func(r chi.Router) {
			// public routes
			r.Get(constant.RootPattern+"public", module.Handler.Public)
			r.Group(func(r chi.Router) {
				// protected routes
				// r.Use(rbacMiddleWare.AuthRole)
				// r.Use(authMiddleWare.AuthUser)
				r.Post(constant.RootPattern, module.Handler.Create)
				r.Get(constant.RootPattern, module.Handler.ReadMany)
				r.Get(constant.RootPattern+"{id}", module.Handler.ReadOne)
				r.Patch(constant.RootPattern+"{id}", module.Handler.Update)
				r.Delete(constant.RootPattern+"{id}", module.Handler.Delete)
			})
		},
	)
}
