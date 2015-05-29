package vitica

import (
	"fmt"
	"net/http"
	"strings"
	"vitica/_vendor/src/github.com/gocraft/web"
)

type Context struct {
	HelloCount int
}

func (c *Context) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *Context) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func StartWebServer() {
	router := web.New(Context{}). // Create your router
					Middleware(LoggerMiddleware).         // Use some included middleware
					Middleware(web.ShowErrorsMiddleware). // ...
					Middleware((*Context).SetHelloCount). // Your own middleware!
					Get("/", (*Context).SayHello)         // Add a route
	http.ListenAndServe("0.0.0.0:8001", router) // Start the server!
}
