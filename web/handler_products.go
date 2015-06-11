package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) HandleProduct(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/product-detail", map[string]string{
		"Title": "Hello Product",
	})
}
