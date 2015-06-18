package web

import (
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/data"
)

func (c *WebContext) HandleCategories(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	path := req.RoutePath()
	_, products := data.GetProducts(data.CategoryOptions{
		Path:   path,
		Themes: req.Form["theme"],
		Low:    req.FormValue("low"),
		High:   req.FormValue("high"),
	})
	RenderCategories(rw, products, getTitleFromPath(path))
}

func getTitleFromPath(path string) string {
	switch path {
	default:
		return ""
	case "/":
		return "All Products"
	case "/popular":
		return "Popular"
	case "/new":
		return "New Releases"
	case "/picks":
		return "Designer Picks"
	case "/sale":
		return "On Sale"
	}
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
