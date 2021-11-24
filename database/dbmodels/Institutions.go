package dbmodels

// Institutions ..
type Institutions struct {
	DatabaseModel
	Code string `gorm:"code" json:"code" example:"IDM"`
	Name string `gorm:"name" json:"name" example:"Indomarco"`
}

// TableName ..
func (t *Institutions) TableName() string {
	return "institutions"
}
