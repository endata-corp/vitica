package vitica

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) HandleCategories(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/categories", map[string]string{
		"Title": "Categories",
	})
}

func (c *WebContext) HandleCategory(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/category", map[string]string{
		"Title": "Category",
	})
}
