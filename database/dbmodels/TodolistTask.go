package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Task ..
type Task struct {
	gorm.Model
	ActionDate            time.Time `json:"action_date"`
	ActionBy              string    `json:"action_by"`
	SupplierName          string    `json:"supplier_name"`
	TodolistID            uint      `json:"todolist_id"`
	TodolistSubCategoryID uint      `json:"todolist_sub_category_id"`
	FileEdukasi           string    `json:"file_edukasi"`
}
