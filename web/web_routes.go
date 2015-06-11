package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func defineRoutes(router *web.Router) {

	router.Get("/", (*WebContext).HandleCategories)
	router.Get("/categories", (*WebContext).HandleCategories)
	router.Get("/categories/:code", (*WebContext).HandleCategory)

	// "/product/:id:\\d.*" regex for int type
	router.Get("/products", (*WebContext).HandleCategories)
	router.Get("/products/:code", (*WebContext).HandleProduct)
}
