package dbmodels

// Cities ..
type Cities struct {
	DatabaseModel
	Name       string `gorm:"column:name" json:"name"`
	ProvinceID int    `gorm:"column:province_id" json:"province_id"`
	CityCode   string `gorm:"column:city_code" json:"city_code"`
}

// TableName ..
func (t *Cities) TableName() string {
	return "public.cities"
}
