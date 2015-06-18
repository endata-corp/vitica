package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/data"
)

func (c *WebContext) HandleAllCategories(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetAllProducts()
	RenderCategories(rw, products, "All Products")
}

func (c *WebContext) HandlePopularCategory(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetPopularProducts()
	RenderCategories(rw, products, "Popular")
}

func (c *WebContext) HandleNewCategory(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetNewProducts()
	RenderCategories(rw, products, "New Releases")
}

func (c *WebContext) HandleDesignerPicksCategory(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetDesignerPicksProducts()
	RenderCategories(rw, products, "Designer Picks")
}

func (c *WebContext) HandleOnSaleCategory(rw web.ResponseWriter, req *web.Request) {
	_, products := data.GetOnSaleProducts()
	RenderCategories(rw, products, "On Sale")
}

func RenderCategories(rw web.ResponseWriter, products []data.Product, title string) {
	Render(rw, "catalog/categories", struct {
		Title    string
		Products []data.Product
	}{
		"VITICA | " + title,
		products,
	})
}
