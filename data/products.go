package data

func GetAllProducts() (err error, products []Product) {
	products = []Product{}
	db := DB()
	err = db.Find(&products).Error
	return
}
