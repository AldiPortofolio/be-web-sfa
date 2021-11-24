package dbmodels

import "time"

// Area ..
type Area struct {
	ID        uint
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string `json:"name"`
	Code      string `json:"code"`
	BranchID  uint   `json:"branch_id"`
	SubAreas  []SubArea
	Districts []District `gorm:"many2many:areas_districts;"`
	Branch    *Branch
}

// AreaList ..
type AreaList struct {
	ID        uint     `json:"id"`
	Branch    string   `json:"branch"`
	Name      string   `json:"name"`
	Districts []string `json:"districts"`
}

// AreaListDB ..
type AreaListDB struct {
	ID        uint   `json:"id"`
	Branch    string `json:"branch"`
	Name      string `json:"name"`
	Districts string `json:"districts"`
}

// AreaDetail ..
type AreaDetail struct {
	ID        uint             `json:"id" gorm:"primary_key"`
	Name      string           `json:"name"`
	Code      string           `json:"code"`
	Districts []CustomDistrict `json:"districts" gorm:"many2many:areas_districts;association_jointable_foreignkey:district_id;jointable_foreignkey:area_id;"`
	BranchID  uint             `json:"branch_id"`
	Branch    CustomBranch     `json:"branch"`
	Region    interface{}      `json:"region"`
}

// CustomDistrict ..
type CustomDistrict struct {
	ID   string `json:"id" gorm:"primary_key:true;"`
	Name string `json:"name"`
}

// CustomBranch ..
type CustomBranch struct {
	ID   uint   `json:"id" gorm:"primary_key:true"`
	Name string `json:"name"`
}

// AreaReq ..
type AreaReq struct {
	AreaID      uint   `json:"area_id"`
	BranchID    uint   `json:"branch_id"`
	Name        string `json:"name"`
	Code        string `json:"code"`
	DistrictIds []int  `json:"district_ids"`
}

// BulkDeleteAreaReq ..
type BulkDeleteAreaReq struct {
	AreaIDs []int `json:"area_ids"`
}

// AreasDistrict ..
type AreasDistrict struct {
	AreaID     uint `json:"area_id"`
	DistrictID uint `json:"district_id"`
}

// AreaListByBranch ..
type AreaListByBranch struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// TableName ..
func (t *AreaDetail) TableName() string {
	return "areas"
}

// TableName ..
func (t *CustomDistrict) TableName() string {
	return "districts"
}

// TableName ..
func (t *CustomBranch) TableName() string {
	return "branches"
}

// TableName ..
func (t *AreasDistrict) TableName() string {
	return "areas_districts"
}
