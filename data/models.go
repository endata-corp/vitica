package data

import (
	"time"
)

type BaseModelKey struct {
	ID uint `gorm:"primary_key"`
}

type BaseModelTimestamps struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}

type Garment struct {
	Code        string `sql:"type:varchar(200);unique_index"`
	Name        string
	Description string `sql:"size:5000"`
}

type Product struct {
	BaseModelKey
	Code               string `sql:"type:varchar(200);unique_index"`
	Slug               string `sql:"type:varchar(200);unique_index"`
	Name               string
	Description        string `sql:"size:5000"`
	DesignId           string
	GarmentIdM         uint
	GarmentM           Garment
	GarmentIdW         uint
	GarmentW           Garment
	GarmentColor       string
	GarmentHex         string
	GarmentDescription string `sql:"size:5000"`
	AvailableSizesM    string
	AvailableSizesW    string
	PriceSuggested     float64
	PriceCost          float64
	Price              float64
	IsNewRelease       bool `sql:"DEFAULT:0"`
	IsOnSale           bool `sql:"DEFAULT:0"`
	IsTopSeller        bool `sql:"DEFAULT:0"`
	IsDesignerPick     bool `sql:"DEFAULT:0"`
	ImageUrl           string
	ImageUrlM          string
	ImageUrlW          string
	ImageSideUrl       string
	ImageBackUrl       string
	ImageDesignUrl     string
	IsActive           bool         `sql:"DEFAULT:1"`
	IsApproved         bool         `sql:"DEFAULT:0"`
	ProductTag         []ProductTag `gorm:"one2many:product_tags;"`
	BaseModelTimestamps
}

type ProductTag struct {
	BaseModelKey
	ProductID int `sql:"index"`
	TagCode   string
}

type Tag struct {
	BaseModelKey
	Code string `gorm:"primary_key`
	Name string `sql:"type:varchar(200);unique_index"`
}
