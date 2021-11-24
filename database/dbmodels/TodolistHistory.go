package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// TodolistHistory ..
type TodolistHistory struct {
	gorm.Model
	// ID                     uint `json:"id"`
	Description  string    `json:"description"`
	FotoLocation string    `json:"foto_location"`
	NewTaskDate  time.Time `json:"new_task_date"`
	OldTaskDate  time.Time `json:"old_task_date"`
	Status       string    `json:"status"`
	TodolistID   uint      `json:"todolist_id"`
	Longitude    string    `json:"longitude"`
	Latitude     string    `json:"latitude"`
}
