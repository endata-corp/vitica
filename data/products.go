package data

func GetAllProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Find(&products).Error
	return
}

func GetProductById(code string) (err error, product Product) {
	product = Product{}
	db := DB()
	err = db.Where("code = '" + code + "'").Find(&product).Error
	return
}
