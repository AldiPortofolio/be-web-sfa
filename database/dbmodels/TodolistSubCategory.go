package dbmodels

import (
	"github.com/jinzhu/gorm"
)

// TodolistSubCategory ..
type TodolistSubCategory struct {
	gorm.Model
	Name       string `json:"name"`
	Code       string `json:"code"`
	CategoryID int    `json:"category_id"`
}
