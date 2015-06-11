package web

import (
	"net/http"
	"os"
	"vitica/_vendor/src/github.com/fvbock/endless"
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/_vendor/src/github.com/unrolled/render"
)

var DEV_MODE = (os.Getenv("GO_ENV") != "PRODUCTION")
var APP_DIR = os.Getenv("VITICA_DIR")

type WebContext struct {
	HelloCount int
}

// Setup rendering options
func Render(rw http.ResponseWriter, name string, binding interface{}) {
	options := render.Options{
		Layout:        "layout",
		IsDevelopment: DEV_MODE,
		Directory:     APP_DIR + "/templates",
	}
	r := render.New(options)
	r.HTML(rw, http.StatusOK, name, binding)
}

func StartWebServer() {
	router := web.New(WebContext{})     // Create your router
	router.Middleware(LoggerMiddleware) // Use some included middleware
	if DEV_MODE {
		router.Middleware(web.ShowErrorsMiddleware)
	}
	router.Middleware((*WebContext).MyMiddleWare) // My own middleware!
	router.NotFound((*WebContext).NotFound)
	router.Middleware(web.StaticMiddleware(APP_DIR+"/public/images", web.StaticOption{Prefix: "/images"}))
	router.Middleware(web.StaticMiddleware(APP_DIR+"/public/assets", web.StaticOption{Prefix: "/assets"}))

	defineRoutes(router)

	if DEV_MODE {
		http.ListenAndServe("0.0.0.0:8000", router)
	} else {
		endless.ListenAndServe("0.0.0.0:8000", router)
	}
}
