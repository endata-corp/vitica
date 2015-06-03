package vitica

import (
	"net/http"
	"os"
	"vitica/_vendor/src/github.com/fvbock/endless"
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/_vendor/src/github.com/unrolled/render"
)

type WebContext struct {
	HelloCount int
}

func Render(rw http.ResponseWriter, name string, binding interface{}) {
	options := render.Options{
		Layout:        "layout",
		IsDevelopment: (os.Getenv("GO_ENV") != "PRODUCTION"),
	}
	r := render.New(options)
	r.HTML(rw, http.StatusOK, name, binding)
}

func StartWebServer() {
	router := web.New(WebContext{})               // Create your router
	router.Middleware(LoggerMiddleware)           // Use some included middleware
	router.Middleware(web.ShowErrorsMiddleware)   // ...
	router.Middleware((*WebContext).MyMiddleWare) // My own middleware!
	router.NotFound((*WebContext).NotFound)

	createRoutes(router)
	router.Middleware(web.StaticMiddleware("public/images", web.StaticOption{Prefix: "/images"}))
	router.Middleware(web.StaticMiddleware("public/assets", web.StaticOption{Prefix: "/assets"}))

	if os.Getenv("GO_ENV") == "PRODUCTION" {
		endless.ListenAndServe("0.0.0.0:8000", router)
	} else {
		http.ListenAndServe("0.0.0.0:8000", router)
	}
}

func createRoutes(router *web.Router) {
	router.Get("/", (*WebContext).GetHome)
	router.Get("/product", (*WebContext).GetProduct)
}
