package dbmodels

// Role ..
type Role struct {
	ID   int    `gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *Role) TableName() string {
	return "public.roles"
}
