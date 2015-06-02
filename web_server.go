package vitica

import (
	"vitica/_vendor/src/github.com/fvbock/endless"
	"vitica/_vendor/src/github.com/gocraft/web"
)

type WebContext struct {
	HelloCount int
}

func StartWebServer() {
	router := web.New(WebContext{})               // Create your router
	router.Middleware(LoggerMiddleware)           // Use some included middleware
	router.Middleware(web.ShowErrorsMiddleware)   // ...
	router.Middleware((*WebContext).MyMiddleWare) // My own middleware!
	router.NotFound((*WebContext).NotFound)
	createRoutes(router)
	router.Middleware(web.StaticMiddleware("public",
		web.StaticOption{Prefix: "/static"}))
	endless.ListenAndServe("0.0.0.0:8000", router) // Start the server!
}

func createRoutes(router *web.Router) {
	router.Get("/", (*WebContext).GetHome)
	router.Get("/product", (*WebContext).GetProduct)
}
