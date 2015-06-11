package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/data"
)

func (c *WebContext) HandleCategories(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetAllProducts()
	Render(rw, "catalog/categories", struct {
		Title    string
		Products []data.Product
	}{
		"All Categories",
		products,
	})
}

func (c *WebContext) HandleCategory(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/category", map[string]string{
		"Title": "Category",
	})
}
