package dbmodels

// Countries ..
type Countries struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *Countries) TableName() string {
	return "public.countries"
}
