package dbmodels

// District ..
type District struct {
	DatabaseModel
	Name         string `gorm:"column:name" json:"name"`
	CityID       int    `gorm:"column:city_id" json:"city_id"`
	DistrictCode string `gorm:"column:district_code" json:"district_code"`
}

// TableName ..
func (t *District) TableName() string {
	return "public.districts"
}
