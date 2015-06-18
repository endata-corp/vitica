package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func defineRoutes(router *web.Router) {

	// Category routes
	router.Get("/", (*WebContext).HandleCategories)
	router.Get("/popular", (*WebContext).HandleCategories)
	router.Get("/new", (*WebContext).HandleCategories)
	router.Get("/picks", (*WebContext).HandleCategories)
	router.Get("/sale", (*WebContext).HandleCategories)

	// "/product/:id:\\d.*" regex for int type
	router.Get("/products/:code", (*WebContext).HandleProduct)
}
