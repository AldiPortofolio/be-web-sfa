package dbmodels

// SalesType ..
type SalesType struct {
	Id   int 	`gorm:"column:id" json:"id"`
	Name string `gorm:"column:name" json:"name"`
	Desc string `gorm:"column:description" json:"description"`
}

// TableName ..
func (t *SalesType) TableName() string {
	return "public.sales_types"
}
