package dbmodels

import (
	"time"
)

// Branch ..
type Branch struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string `json:"name"`
	Code         string `json:"code"`
	BranchOffice string `json:"branch_office"`
	Cities       []City `json:"cities" gorm:"many2many:branches_cities;"`
	RegionID     uint   `json:"region_id"`
	Region       *Region
}

// BranchList ..
type BranchList struct {
	ID     uint     `json:"id"`
	Region string   `json:"region"`
	Name   string   `json:"name"`
	Cities []string `json:"cities"`
}

// BranchListDB ..
type BranchListDB struct {
	ID     uint   `json:"id"`
	Region string `json:"region"`
	Name   string `json:"name"`
	Cities string `json:"cities"`
}

// BranchDetail ..
type BranchDetail struct {
	ID           uint
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string       `json:"name"`
	Code         string       `json:"code"`
	BranchOffice string       `json:"branch_office"`
	Cities       []CustomCity `json:"cities" gorm:"many2many:branches_cities;association_jointable_foreignkey:city_id;jointable_foreignkey:branch_id;"`
	RegionID     uint         `json:"region_id"`
	Region       CustomRegion `json:"region"`
}

// CustomCity ..
type CustomCity struct {
	ID   string `json:"id" gorm:"primary_key:true;"`
	Name string `json:"name"`
}

// CustomRegion ..
type CustomRegion struct {
	ID   uint   `json:"id" gorm:"primary_key:true"`
	Name string `json:"name"`
}

// BranchReq ..
type BranchReq struct {
	BranchID     uint   `json:"branch_id"`
	RegionID     uint   `json:"region_id"`
	Name         string `json:"name"`
	Code         string `json:"code"`
	BranchOffice string `json:"branch_office"`
	CityIds      []int  `json:"city_ids"`
}

// AutoCityList ..
type AutoCityList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// BulkDeleteBranchReq ..
type BulkDeleteBranchReq struct {
	BranchIDs []int `json:"branch_ids"`
}

// BranchListByRegion ..
type BranchListByRegion struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// TableName ..
func (t *BranchDetail) TableName() string {
	return "branches"
}

// TableName ..
func (t *CustomCity) TableName() string {
	return "cities"
}

// TableName ..
func (t *CustomRegion) TableName() string {
	return "regions"
}
