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

type Garment struct {
	BaseModel
	Name        string
	Description string
	Brand       string
	Material    string
	StyleId     string
}

type GarmentColor struct {
	BaseModel
	Name string
	Hex  string
}

type GarmentSize struct {
	BaseModel
	Code        string
	Name        string
	Description string
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
