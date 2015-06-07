package vitica

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func defineRoutes(router *web.Router) {
	router.Get("/", (*WebContext).GetHome)
	//router.Get("/product/:id:\\d.*", (*WebContext).GetProductDetail)
	router.Get("/product/:id", (*WebContext).GetProductDetail)
}
