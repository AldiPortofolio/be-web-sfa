package dbmodels

import (
	"github.com/jinzhu/gorm"
)

// TodolistCategory ..
type TodolistCategory struct {
	gorm.Model
	Name string `json:"name"`
	Code string `json:"code"`
}
