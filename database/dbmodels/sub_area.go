package dbmodels

import "time"

// SubArea ..
type SubArea struct {
	ID                uint
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Name              string             `json:"name"`
	Code              string             `json:"code"`
	AreaID            uint               `json:"area_id"`
	Villages          []Village          `gorm:"many2many:sub_areas_villages;omitempty"`
	SalesAreaChannels []SalesAreaChannel `gorm:"many2many:sales_area_channels_sub_areas;omitempty"`
}

// SubAreaList ..
type SubAreaList struct {
	ID               uint     `json:"id"`
	Area             string   `json:"area"`
	Name             string   `json:"name"`
	Villages         []string `json:"villages"`
	SalesAreaChannel string   `json:"sales_area_channel"`
}

// SubAreaListDB ..
type SubAreaListDB struct {
	ID               uint   `json:"id"`
	Area             string `json:"area"`
	Name             string `json:"name"`
	Villages         string `json:"villages"`
	SalesAreaChannel string `json:"sales_area_channel"`
}

// SubAreaDetail ..
type SubAreaDetail struct {
	ID               uint            `json:"id" gorm:"primary_key"`
	Name             string          `json:"name"`
	Code             string          `json:"code"`
	Villages         []CustomVillage `json:"villages" gorm:"many2many:sub_areas_villages;association_jointable_foreignkey:village_id;jointable_foreignkey:sub_area_id;"`
	AreaID           uint            `json:"area_id"`
	Area             CustomArea      `json:"area"`
	Branch           interface{}     `json:"branch"`
	Region           interface{}     `json:"region"`
	SalesAreaChannel interface{}     `json:"sales_area_channel"`
}

// CustomVillage ..
type CustomVillage struct {
	ID   string `json:"id" gorm:"primary_key:true;"`
	Name string `json:"name"`
}

// CustomArea ..
type CustomArea struct {
	ID   string `json:"id" gorm:"primary_key:true;"`
	Name string `json:"name"`
}

// SubAreaReq ..
type SubAreaReq struct {
	SubAreaID  uint   `json:"sub_area_id"`
	AreaID     uint   `json:"area_id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	VillageIds []int  `json:"village_ids"`
	SacID      uint   `json:"sac_id"`
}

// BulkSubArea ..
type BulkSubArea struct {
	AreaName         string `csv:"Area Name"`
	SubAreaName      string `csv:"Sub Area Name"`
	VillageIDs       string `csv:"Village IDs"`
	SalesAreaChannel string `csv:"Sales Area Channel"`
}

// DataSubAreaErrorByRow ..
type DataSubAreaErrorByRow struct {
	NoRow            int    `csv:"Nomor Row"`
	AreaName         string `csv:"Area Name"`
	SubAreaName      string `csv:"Sub Area Name"`
	VillageIDs       string `csv:"Village IDs"`
	SalesAreaChannel string `csv:"Sales Area Channel"`
	ErrorMessages    string `csv:"Error Messages"`
}

// AutoVillageList ..
type AutoVillageList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

// SubAreaListByArea ..
type SubAreaListByArea struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

// SubAreaTemplate ..
type SubAreaTemplate struct {
	ID           uint   `json:"id"`
	TemplateFile string `json:"template_file"`
}

// BulkDeleteSubAreaReq ..
type BulkDeleteSubAreaReq struct {
	SubAreaIDs []int `json:"sub_area_ids"`
}

// TableName ..
func (t *SubAreaDetail) TableName() string {
	return "sub_areas"
}

// TableName ..
func (t *CustomVillage) TableName() string {
	return "villages"
}

// TableName ..
func (t *CustomArea) TableName() string {
	return "areas"
}
