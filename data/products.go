package data

func GetAllProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Find(&products).Error
	return
}

func GetProductByCode(code string) (err error, product Product) {
	product = Product{}
	db := DB()
	err = db.Where("code = '" + code + "'").Find(&product).Error
	return
}

func GetPopularProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Where("is_top_seller = 1").Find(&products).Error
	return
}

func GetNewProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Where("is_new_release = 1").Find(&products).Error
	return
}

func GetDesignerPicksProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Where("is_designer_pick = 1").Find(&products).Error
	return
}

func GetOnSaleProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Where("is_on_sale = 1").Find(&products).Error
	return
}
