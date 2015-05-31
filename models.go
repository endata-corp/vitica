package vitica

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
	BaseModelKey
	Code        string
	Name        string
	Description string `sql:"size:5000"`
	BaseModelTimestamps
}

type Product struct {
	BaseModelKey
	Code               string
	Name               string
	Description        string `sql:"size:5000"`
	DesignId           string
	GarmentIdM         string
	GarmentM           Garment
	GarmentIdW         string
	GarmentW           Garment
	GarmentColor       string
	GarmentHex         string
	GarmentDescription string `sql:"size:5000"`
	AvailableSizesM    string
	AvailableSizesW    string
	PriceSuggested     float32
	PriceCost          float32
	Price              float32
	ImageUrl           string
	ImageUrlM          string
	ImageUrlW          string
	ImageSideUrl       string
	ImageBackUrl       string
	ImageDesignUrl     string
	IsActive           bool `sql:"DEFAULT:1"`
	IsApproved         bool `sql:"DEFAULT:0"`
	BaseModelTimestamps
}
