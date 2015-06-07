package vitica

import (
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) GetHome(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "home", map[string]string{
		"Title": "Vitica",
	})
}
