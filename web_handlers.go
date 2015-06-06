package vitica

import (
	"fmt"
	"net/http"
	"vitica/_vendor/src/github.com/gocraft/web"
)

func (c *WebContext) NotFound(rw web.ResponseWriter, r *web.Request) {
	rw.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(rw, "My Not Found")
}

func (c *WebContext) GetHome(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "sample/blank", map[string]string{
		"Title": "My Home",
	})
}

func (c *WebContext) GetProductDetail(rw web.ResponseWriter, req *web.Request) {
	Render(rw, "catalog/product-detail", map[string]string{
		"Title": "Hello Product",
	})
}

func (c *WebContext) GetProduct(rw web.ResponseWriter, req *web.Request) {

	product := Product{}
	product.Code = "The product code"
	Render(rw, "sample/example", map[string]string{
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
