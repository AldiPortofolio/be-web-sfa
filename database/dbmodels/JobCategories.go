package dbmodels

// JobCategories ..
type JobCategories struct {
	ID          int64  `gorm:"column:id" json:"id"`
	Name        string `gorm:"column:name" json:"name"`
	Description string `gorm:"column:description" json:"description"`
}

// TableName ..
func (t *JobCategories) TableName() string {
	return "public.job_categories"
}
