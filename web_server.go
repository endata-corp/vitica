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
	router := web.New(Context{})                // Create your router
	router.Middleware(LoggerMiddleware)         // Use some included middleware
	router.Middleware(web.ShowErrorsMiddleware) // ...
	router.Middleware((*Context).SetHelloCount) // Your own middleware!
	router.Get("/", (*Context).SayHello)        // Add a route
	http.ListenAndServe("0.0.0.0:8000", router) // Start the server!
}
