package dbmodels

// Provinces ..
type Provinces struct {
	DatabaseModel
	Name         string `gorm:"column:name" json:"name" example:"IDM"`
	ProvinceCode string `gorm:"column:province_code" json:"province_code" example:"IDM"`
	//CountryID 		int       `gorm:"column:country_id" json:"country_id"`
}

// ProvincesByCountry ..
type ProvincesByCountry struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *Provinces) TableName() string {
	return "public.provinces"
}

// TableName ..
func (t *ProvincesByCountry) TableName() string {
	return "public.provinces"
}
