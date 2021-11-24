package dbmodels

import (
	"github.com/jinzhu/gorm"
)

// Task ..
type LabelTask struct {
	gorm.Model
	Name          string `json:"name"`
	LabelType     string `json:"label_type"`
	SubCategoryID uint   `json:"sub_category_id"`
}
