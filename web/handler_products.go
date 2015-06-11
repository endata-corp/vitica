package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/data"
)

func (c *WebContext) HandleProduct(rw web.ResponseWriter, req *web.Request) {
	code := req.PathParams["code"]
	_, product := data.GetProductById(code)
	Render(rw, "catalog/product-detail", struct {
		Title   string
		Product data.Product
	}{
		"My Product",
		product,
	})
}
