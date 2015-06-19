package data

import (
	"strings"
	"vitica/log"
)

type CategoryOptions struct {
	Path   string
	Themes []string
	Price  string
}

func GetProducts(options CategoryOptions) (err error, products []Product) {
	log.Info("%+v ", options)
	products = []Product{}
	db := DB()
	db = db.Where(getFieldNameFromPath(options.Path) + " = 1")
	if options.Price != "" && options.Price != "any" {
		prices := strings.Split(options.Price, "-")
		db = db.Where("price >= " + prices[0])
		db = db.Where("price <= " + prices[1])
	}
	err = db.Find(&products).Error
	return
}

func getFieldNameFromPath(path string) string {
	switch path {
	default:
		return "1"
	case "/":
		return "1"
	case "/popular":
		return "is_top_seller"
	case "/new":
		return "is_new_release"
	case "/picks":
		return "is_designer_pick"
	case "/sale":
		return "is_on_sale"
	}
}

func GetProductByCode(code string) (err error, product Product) {
	product = Product{}
	db := DB()
	err = db.Where("code = '" + code + "'").Find(&product).Error
	return
}
