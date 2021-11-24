package models

import (
	"ottosfa-api-web/database/dbmodels"
	"time"

	"github.com/jinzhu/gorm"
)

// CreateTodolist models
type CreateTodolist struct {
	Mid                string       `json:"merchant_id"`
	TaskDate           string       `json:"task_date"`
	SalesPhone         string       `json:"sales_phone"`
	TodolistCategoryID uint         `json:"todolist_category_id"`
	VillageID          int64        `json:"village_id"`
	Notes              string       `json:"notes"`
	CustomerCode       string       `json:"customer_code"`
	MerchantPhone      string       `json:"merchant_phone"`
	MerchantName       string       `json:"merchant_name"`
	SalesTypeID        string       `json:"sales_type_id"`
	TasksAttributes    []TaskParams `json:"task_attributes"`
}

// CreateTodolistV2 models
type CreateTodolistV2 struct {
	Mid                string       `json:"mid" example:"OP1B00000083"`
	TaskDate           string       `json:"task_date" example:"2021-24-04"`
	SalesPhone         string       `json:"sales_phone" example:"082113776997"`
	TodolistCategoryID uint         `json:"todolist_category_id" example:"6"`
	VillageID          int64        `json:"village_id" example:"3171100008"`
	Notes              string       `json:"notes" example:""`
	CustomerCode       string       `json:"customer_code" example:"5200522100082316.0"`
	MerchantPhone      string       `json:"merchant_phone" example:"082222230118"`
	MerchantName       string       `json:"merchant_name" example:"Merchant Test"`
	TasksAttributes    []TaskParams `json:"task_attributes"`
	Address            string       `json:"address" example:"AIA Central Lt. 27, Jl. Jend. Sudirman No.Kav 48A, RT.5/RW.4, Karet Semanggi, Setia Budi, RT.5/RW.4, Karet Semanggi, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12930, Indonesia"`
	AddressBenchmark   string       `json:"address_benchmark" example:"sebelah rumah sakit jiwa"`
	OwnerName          string       `json:"owner_name" example:"Wawan"`
	MerchantID         int          `json:"merchant_id" example:"53229"`
	SalesTypeID        int          `json:"sales_type_id" example:"2"`
}

// TaskParams ..
type TaskParams struct {
	SubCategoryID string `json:"sub_category_id"`
	SupplierName  string `json:"supplier_name,omitempty"`
	FileEdukasi   string `json:"file_edukasi"`
}

// TodolistDetail ..
type TodolistDetail struct {
	ID                   uint      `json:"id"`
	ActionDate           string    `json:"action_date"`
	CustomerCode         string    `json:"customer_code"`
	MerchantPhone        string    `json:"merchant_phone"`
	MerchantName         string    `json:"merchant_name"`
	Mid                  string    `json:"mid"`
	SalesPhone           string    `json:"sales_phone"`
	Status               string    `json:"status"`
	TaskDate             string    `json:"task_date"`
	TodolistCategoryID   int64     `json:"todolist_category_id"`
	TodolistCategoryName string    `json:"category_name"`
	VillageID            int64     `json:"village_id"`
	CreatedAt            time.Time `json:"created_at"`
	Notes                string    `json:"notes"`
	Tasks                []TaskDetail
}

// TaskDetail ..
type TaskDetail struct {
	gorm.Model
	ActionDate            time.Time           `json:"action_date"`
	ActionBy              string              `json:"action_by"`
	SupplierName          string              `json:"supplier_name"`
	TodolistID            uint                `json:"todolist_id"`
	TodolistSubCategoryID uint                `json:"todolist_sub_category_id"`
	SubCategoryName       string              `json:"sub_category_name"`
	FileEdukasi           string              `json:"file_edukasi"`
	FollowUps             []dbmodels.FollowUp `json:"follow_ups"`
}

// NewMerchantDetail ..
type NewMerchantDetail struct {
	ID            uint   `json:"id"`
	MerchantID    string `json:"merchant_id"`
	MerchantName  string `json:"merchant_name"`
	MerchantPhone string `json:"merchant_phone"`
	MerchantImage string `json:"merchant_image"`
	OwnerName     string `json:"owner_name"`
	SubArea       string `json:"sub_area"`
	Address       string `json:"address"`
	CustomerCode  string `json:"customer_code"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
	VillageID     int64  `json:"village_id"`
}

// NewMerchantDetailReq ..
type NewMerchantDetailReq struct {
	CustomerCode string `json:"customer_code"`
	PhoneNumber  string `json:"phone_number"`
}

// ShowTodolist ..
type ShowTodolist struct {
	ID                 uint         `json:"id"`
	ActionDate         string       `json:"action_date"`
	ActionBy           string       `json:"action_by"`
	CreatedAt          time.Time    `json:"created_at"`
	Mid                string       `json:"mid"`
	CustomerCode       string       `json:"customer_code"`
	MerchantPhone      string       `json:"merchant_phone"`
	Status             string       `json:"status"`
	TaskDate           string       `json:"task_date"`
	TodolistCategoryID int64        `json:"todolist_category_id"`
	CategoryName       string       `json:"category_name"`
	PossibleSales      []SalesList  `json:"possible_sales"`
	MerchantDetail     interface{}  `json:"merchant_detail"`
	Longitude          string       `json:"longitude"`
	Latitude           string       `json:"latitude"`
	Notes              string       `json:"notes"`
	SalesPhone         string       `json:"sales_phone"`
	Tasks              []TaskDetail `json:"tasks"`
	Histories          []History    `json:"histories"`
}

// SalesList ..
type SalesList struct {
	ID         uint   `json:"id"`
	LabelSales string `json:"label_sales"`
}

// History ..
type History struct {
	ID           uint         `json:"id"`
	Description  string       `json:"description"`
	FotoLocation string       `json:"foto_location"`
	NewTaskDate  time.Time    `json:"new_task_date"`
	OldTaskDate  time.Time    `json:"old_task_date"`
	Status       string       `json:"status"`
	TodolistID   uint         `json:"todolist_id"`
	Longitude    string       `json:"longitude"`
	Latitude     string       `json:"latitude"`
	Tasks        []TaskDetail `json:"tasks"`
}

// Todolist ..
type Todolist struct {
	ID           uint `json:"id"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	ActionDate   time.Time `json:"action_date"`
	MerchantName string    `json:"merchant_name"`
	Mid          string    `json:"mid"`
	// CustomerCode             string    `json:"customer_code"`
	MerchantPhone            string    `json:"merchant_phone"`
	SalesPhone               string    `json:"sales_phone"`
	Status                   string    `json:"status"`
	TaskDate                 time.Time `json:"task_date"`
	TodolistCategoryID       uint      `json:"todolist_category_id"`
	VillageID                int64     `json:"village_id"`
	Longitude                string    `json:"longitude"`
	Latitude                 string    `json:"latitude"`
	Notes                    string    `json:"notes"`
	MerchantNewRecruitmentID uint      `json:"merchant_new_recruitment_id"`
	Tasks                    []dbmodels.Task
	TodolistHistories        []dbmodels.TodolistHistory
	Address                  string `json:"address"`
	AddressBenchmark         string `json:"address_benchmark"`
	OwnerName                string `json:"owner_name"`
	MerchantID               int    `json:"merchant_id"`
	SalesTypeID              int    `json:"sales_type_id"`
}

// MerchantDetail ..
type MerchantDetail struct {
	ID            uint   `json:"id"`
	MerchantName  string `json:"merchant_name"`
	MerchantPhone string `json:"merchant_phone"`
	// MerchantImage string `json:"merchant_image"`
	MerchantID string `json:"merchant_id"`
	OwnerName  string `json:"owner_name"`
	SubArea    string `json:"sub_area"`
	Address    string `json:"address"`
	Note       string `json:"note"`
}

// NewMerchantListReq ..
type NewMerchantListReq struct {
	Keyword string `json:"keyword"`
}

// NewMerchantList ..
type NewMerchantList struct {
	CustomerCode string `json:"customer_code"`
	PhoneNumber  string `json:"phone_number"`
	Name         string `json:"name"`
}

// UpdateTodolist ..
type UpdateTodolist struct {
	ID              string       `json:"id"`
	MerchantID      string       `json:"merchant_id"` // ID of merchant or merchant_new_recruitment
	CategoryID      string       `json:"category_id"`
	TaskDate        string       `json:"task_date"`
	Notes           string       `json:"notes"`
	SalesPhone      string       `json:"sales_phone"`
	TasksAttributes []TaskParams `json:"task_attributes"`
}

// UpdateTodolistV2 ..
type UpdateTodolistV2 struct {
	ID               string       `json:"id" example:"136129"`
	MerchantID       string       `json:"merchant_id" example:"2"` // ID of merchant or merchant_new_recruitment
	CategoryID       string       `json:"category_id" example:"6"`
	TaskDate         string       `json:"task_date" example:"2021-08-24"`
	Notes            string       `json:"notes"`
	SalesPhone       string       `json:"sales_phone" example:"082113776997"`
	TasksAttributes  []TaskParams `json:"task_attributes"`
	MerchantPhone    string       `json:"merchant_phone" example:"082222230118"`
	Address          string       `json:"address" example:"AIA Central Lt. 27, Jl. Jend. Sudirman No.Kav 48A, RT.5/RW.4, Karet Semanggi, Setia Budi, RT.5/RW.4, Karet Semanggi, Kota Jakarta Selatan, Daerah Khusus Ibukota Jakarta 12930, Indonesia"`
	AddressBenchmark string       `json:"address_benchmark" example:"sebelah rumah sakit jiwa"`
	OwnerName        string       `json:"owner_name" example:"Wawan"`
	SalesTypeID      int          `json:"sales_type_id"`
}
