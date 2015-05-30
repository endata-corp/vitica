package vitica

import (
	"net/http"
	"vitica/_vendor/src/github.com/gocraft/web"
)

type WebContext struct {
	HelloCount int
}

func StartWebServer() {
	router := web.New(WebContext{})                   // Create your router
	router.Middleware(LoggerMiddleware)               // Use some included middleware
	router.Middleware(web.ShowErrorsMiddleware)       // ...
	router.Middleware(web.StaticMiddleware("public")) //
	router.Middleware((*WebContext).MyMiddleWare)     // My own middleware!
	createRoutes(router)
	http.ListenAndServe("0.0.0.0:8000", router) // Start the server!
}

func createRoutes(router *web.Router) {
	router.Get("/product", (*WebContext).GetProduct)
}
