package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func defineRoutes(router *web.Router) {

	// Category routes
	router.Get("/", (*WebContext).HandleAllCategories)
	router.Get("/popular", (*WebContext).HandlePopularCategory)
	router.Get("/new", (*WebContext).HandleNewCategory)
	router.Get("/picks", (*WebContext).HandleDesignerPicksCategory)
	router.Get("/sale", (*WebContext).HandleOnSaleCategory)

	// "/product/:id:\\d.*" regex for int type
	router.Get("/products/:code", (*WebContext).HandleProduct)
}
