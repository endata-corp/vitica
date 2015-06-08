package vitica

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) HandleCategories(rw web.ResponseWriter, req *web.Request) {
	s := []string{"1", "2", "3", "4", "5", "6", "7"}
	data := struct {
		Title    string
		Products []string
	}{
		"MY Categories",
		s,
	}
	Render(rw, "catalog/categories", data)
}

func (c *WebContext) HandleCategory(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/category", map[string]string{
		"Title": "Category",
	})
}
