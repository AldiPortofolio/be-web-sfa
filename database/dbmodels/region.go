package dbmodels

import (
	"time"
)

// Region ..
type Region struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	Code      string `json:"code"`
	Branches  []Branch
	Provinces []Province `json:"provinces" gorm:"many2many:provinces_regions"`
}

// Regions ..
type Regions struct {
	Regions interface{} `json:"regions,omitempty"`
	Meta    interface{} `json:"meta,omitempty"`
}

// RegionsList ..
type RegionsList struct {
	ID        uint   `gorm:"id"`
	Name      string `gorm:"name"`
	Branches  string `gorm:"branches"`
	Provinces string `gorm:"provinces"`
}

// RegionsDB ..
type RegionsDB struct {
	ID        uint     `json:"id"`
	Name      string   `json:"name"`
	Provinces []string `json:"provinces"`
	Branches  []string `json:"branches"`
}

// ProvincesRegion ..
type ProvincesRegion struct {
	RegionID   uint `json:"region_id"`
	ProvinceID uint `json:"province_id"`
}

// RegionDetail ..
type RegionDetail struct {
	ID        uint             `json:"id"`
	Name      string           `json:"name"`
	Code      string           `json:"code"`
	Branches  []BranchCustom   `json:"branches"`
	Provinces []ProvinceCustom `json:"provinces"`
}

// BranchCustom ..
type BranchCustom struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// ProvinceCustom ..
type ProvinceCustom struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// RegionReq ...
type RegionReq struct {
	RegionID    uint   `json:"region_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	ProvinceIDs []int  `json:"province_ids"`
}

// SaveRegion ..
type SaveRegion struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	Code      string `json:"code"`
}

// BulkDeleteRegionReq ..
type BulkDeleteRegionReq struct {
	RegionIDs []int `json:"region_ids"`
}

// TableName ..
func (t *ProvincesRegion) TableName() string {
	return "provinces_regions"
}

// TableName ..
func (t *SaveRegion) TableName() string {
	return "regions"
}
