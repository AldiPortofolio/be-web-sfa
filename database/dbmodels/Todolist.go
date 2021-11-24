package dbmodels

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Todolist models
type Todolist struct {
	// ID                   int64     `json:"id"`
	gorm.Model
	ActionDate               time.Time `json:"action_date"`
	MerchantName             string    `json:"merchant_name"`
	Mid                      string    `json:"mid"`
	SalesPhone               string    `json:"sales_phone"`
	Status                   string    `json:"status"`
	TaskDate                 time.Time `json:"task_date"`
	TodolistCategoryID       uint      `json:"todolist_category_id"`
	VillageID                int64     `json:"village_id"`
	Longitude                string    `json:"longitude"`
	Latitude                 string    `json:"latitude"`
	Notes                    string    `json:"notes"`
	Tasks                    []Task
	TodolistHistories        []TodolistHistory
	MerchantNewRecruitmentID uint   `json:"merchant_new_recruitment_id"`
	MerchantPhone            string `json:"merchant_phone"`
	Address                  string `json:"address"`
	AddressBenchmark         string `json:"address_benchmark"`
	OwnerName                string `json:"owner_name"`
	MerchantID               int    `json:"merchant_id"`
	SalesTypeID              int    `json:"sales_type_id"`

	// CustomerCode       string `json:"customer_code"`
}
