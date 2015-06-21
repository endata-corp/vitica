package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) HandleCart(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/cart", struct {
		Title string
	}{
		"VITICA | Shopping Cart",
	})
}
