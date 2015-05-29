package vitica

import (
	_ "vitica/_vendor/src/github.com/go-sql-driver/mysql"
	"vitica/_vendor/src/github.com/jinzhu/gorm"
)

var dbHandle *gorm.DB

// Returns db handle and connects if needed
func DB() *gorm.DB {
	if dbHandle != nil {
		return dbHandle
	}
	return dbConnect()
}

// Connect to the db
func dbConnect() *gorm.DB {
	err := ReadConfig()
	if err != nil {
		panic(err.Error())
	}
	handle, err := gorm.Open("mysql", GetDBConfig())
	if err != nil {
		Error("Unable to connect to DB %s", err.Error())
	}
	dbHandle = &handle

	err = dbSync()
	if err != nil {
		Error("Error synching db: ", err.Error())
	}
	return dbHandle
}

// Synchronizes the db schema if necessary
func dbSync() error {
	dbHandle = dbHandle.AutoMigrate(
		&Product{},
		&Design{},
	)
	if dbHandle.Error != nil {
		return dbHandle.Error
	}
	return nil
}

// Saves a model
func SaveModel(model interface{}) (s *gorm.DB, err error) {
	handle := DB().Create(model)
	return handle, handle.Error
}

//db.Model(&models.Site{}).Related(&models.Organization{})
//	var (
//		org models.Organization
//		sites []models.Site
//	)
//	db.Where("id = ?", 2).Find(&org)
//	db.Where("organization_id = ?", org.ID).Find(&sites)
//
//	org.Sites = sites
//	out, _ := json.MarshalIndent(org, "", "  ")
//
//	logger.Info(string(out))
//logger.Info("%+v\n", sites)

//	org := models.Organization{}
//	org.Name = "Whole Foods"
//	db.Create(&org)
//
//	site := models.Site{}
//	site.BaseUrl = "http://whatever.com/"
//	site.OrganizationID = 1
//	db.Create(&site)
