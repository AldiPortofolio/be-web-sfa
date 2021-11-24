package dbmodels

import "time"

// Province ..
type Province struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string `json:"name"`
	ProvinceCode string `json:"province_code"`
	CountryID    uint   `json:"country_id"`
	Cities       []City
	Regions      []Region `gorm:"many2many:provinces_regions;"`
}
