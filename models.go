package vitica

import (
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Product struct {
	BaseModel
	Code         string
	Name         string
	Description  string
	Website      string
	LogoUrl      string
	Industry     string
	Headquarters string
	StockSymbol  string
	Employees    string
	Revenue      string
	Reviews      int
	JobCount     int
	IsActive     bool `sql:"DEFAULT:1"`
	IsApproved   bool `sql:"DEFAULT:0"`
}

type Design struct {
	BaseModel
	Code         string
	Name         string
	Description  string
	Website      string
	LogoUrl      string
	Industry     string
	Headquarters string
	StockSymbol  string
	Employees    string
	Revenue      string
	Reviews      int
	JobCount     int
	IsActive     bool `sql:"DEFAULT:1"`
	IsApproved   bool `sql:"DEFAULT:0"`
}
