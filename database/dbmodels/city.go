package dbmodels

import "time"

// City ..
type City struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Name       string `json:"name"`
	CityCode   string `json:"city_code"`
	ProvinceID uint   `json:"province_id"`
	Districts  []District
	Branches   []Branch `gorm:"many2many:branches_cities;"`
}

// TableName ..
func (t *City) TableName() string {
	return "cities"
}
