package vitica

import (
	"net/http"
	"vitica/_vendor/src/github.com/gocraft/web"
	"vitica/_vendor/src/github.com/unrolled/render"
)

func (c *WebContext) GetProduct(rw web.ResponseWriter, req *web.Request) {
	product := Product{}
	product.Code = "The product code"
	r := render.New(render.Options{Layout: "home"})
	r.HTML(rw, http.StatusOK, "sample/example", map[string]string{
		"Message": "json",
	})

	//	r.HTML(rw, http.StatusOK, "sample/example", struct {
	//		Message string
	//	}{
	//		"This is my message here!",
	//	})

	//	orgId := req.PathParams["org_id"]
	//	org := Organization{}
	//	db := DB()
	//	err := db.Preload("Sites").Where("id = ?", orgId).Find(&org).Error
}
