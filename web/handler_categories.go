package web

import (
	"strings"
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/data"
)

func (c *WebContext) HandleCategories(rw web.ResponseWriter, req *web.Request) {
	req.ParseForm()
	path := req.RoutePath()
	themes := req.Form["theme"]
	options := data.CategoryOptions{
		Path:   path,
		Themes: themes,
		Price:  req.FormValue("price"),
	}
	_, products := data.GetProducts(options)
	RenderCategories(rw, products, getTitleFromPath(path), options)
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

func RenderCategories(rw web.ResponseWriter, products []data.Product, title string, options data.CategoryOptions) {
	Render(rw, "catalog/categories", struct {
		Title    string
		Name     string
		Path     string
		Price    string
		Products []data.Product
		Themes   string
	}{
		"VITICA | " + title,
		title,
		options.Path,
		options.Price,
		products,
		strings.Join(options.Themes, ","),
	})
}
