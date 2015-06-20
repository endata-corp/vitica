package data

import (
	"strings"
)

type CategoryOptions struct {
	Path   string
	Themes []string
	Price  string
}

func GetProducts(options CategoryOptions) (err error, products []Product) {
	products = []Product{}
	db := DB()
	db = db.Preload("ProductTag")
	db = db.Where(getFieldNameFromPath(options.Path) + " = 1")
	if options.Price != "" && options.Price != "any" {
		prices := strings.Split(options.Price, "-")
		db = db.Where("price >= " + prices[0])
		db = db.Where("price <= " + prices[1])
	}
	err = db.Find(&products).Error
	products = filterProductsByTheme(products, options.Themes)
	return
}

func filterProductsByTheme(products []Product, themes []string) []Product {
	if len(themes) == 0 || themes[0] == "all" {
		return products
	}
	filtered := []Product{}
	for _, product := range products {
		tags := product.ProductTag
		for _, tag := range tags {
			for _, theme := range themes {
				if tag.TagCode == theme {
					filtered = append(filtered, product)
				}
			}
		}
	}
	return filtered
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
