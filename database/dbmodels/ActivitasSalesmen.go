package dbmodels

import "time"

// ActivitasSalesmenList ..
type ActivitasSalesmenList struct {
	ID                        int    `gorm:"column:id" json:"id"`
	Name                      string `gorm:"column:name" json:"name"`
	PhoneNumber               string `gorm:"column:phone_number" json:"phone_number"`
	SalesType                 string `gorm:"column:sales_type" json:"sales_type"`
	SalesID                   string `gorm:"column:sales_id" json:"sales_id"`
	Photo                     string `gorm:"column:photo" json:"photo"`
	Akusisi                   int    `gorm:"column:akusisi" json:"akusisi"`
	NOO                       int    `gorm:"column:noo" json:"noo"`
	TodoListCount             int    `gorm:"column:todolist_count" json:"todolist_count"`
	CallPlanDate              string `gorm:"column:call_plan_date" json:"call_plan_date"`
	ActionDate                string `gorm:"column:action_date" json:"action_date"`
	SuccessCallPlanCount      int    `gorm:"column:success_callplan_count" json:"success_callplan_count"`
	TotalCallPlanCount        int    `gorm:"column:total_callplan_count" json:"total_callplan_count"`
	SuccessCallPlanPercentage string `gorm:"column:success_call" json:"success_call"`
}

// ActivitasSalesmenDetailTodolist ..
type ActivitasSalesmenDetailTodolist struct {
	ID               int    `gorm:"column:id" json:"id"`
	Date             string `gorm:"column:date" json:"date"`
	MerchantName     string `gorm:"column:merchant_name" json:"merchant_name"`
	MID              string `gorm:"column:mid" json:"mid"`
	MerchantTypeName string `gorm:"column:merchant_type_name" json:"merchant_type_name"`
	Category         string `gorm:"column:category" json:"category"`
	Status           string `gorm:"column:status" json:"status"`
}

// ActivitasSalesmenDetailCallplan ..
type ActivitasSalesmenDetailCallplan struct {
	ID               int    `gorm:"column:id" json:"id"`
	Date             string `gorm:"column:call_plan_date" json:"call_plan_date"`
	ActionDate       string `gorm:"column:action_date" json:"action_date"`
	SalesID          string `gorm:"column:sales_id" json:"sales_id"`
	MerchantName     string `gorm:"column:merchant_name" json:"merchant_name"`
	MID              string `gorm:"column:mid" json:"mid"`
	MerchantTypeName string `gorm:"column:merchant_type_name" json:"merchant_type_name"`
	MerchantStatus   string `gorm:"column:merchant_status" json:"merchant_status"`
	Kelurahan        string `gorm:"column:kelurahan" json:"kelurahan"`
	Status           string `gorm:"column:status" json:"status"`
}

// ActivitasSalesmenDetail ..
type ActivitasSalesmenDetail struct {
	ID                        int     `gorm:"column:id" json:"id"`
	Name                      string  `gorm:"column:name" json:"name"`
	Photo                     string  `gorm:"column:photo" json:"photo"`
	PhoneNumber               string  `gorm:"column:phone_number" json:"phone_number"`
	SalesType                 string  `gorm:"column:sales_type" json:"sales_type"`
	SubArea                   string  `gorm:"column:sub_area" json:"sub_area"`
	Akusisi                   int     `gorm:"column:akusisi" json:"akusisi"`
	NOO                       int     `gorm:"column:noo" json:"noo"`
	TodoListCount             int     `gorm:"column:todolist_count" json:"todolist_count"`
	SuccessCallPlanCount      int     `gorm:"column:success_callplan_count" json:"success_callplan_count"`
	TotalCallPlanCount        int     `gorm:"column:total_callplan_count" json:"total_callplan_count"`
	SuccessCallPlanPercentage string  `gorm:"column:success_callplan_percentage" json:"success_callplan_percentage"`
	Amount                    float64 `gorm:"column:amount" json:"amount"`
	CallPlanDate              string  `gorm:"column:call_plan_date" json:"call_plan_date"`
	Villages                  string  `gorm:"column:villages" json:"villages"`
	SAC                       string  `gorm:"column:sac" json:"sac"`
}

// ListActivitasSalesmenReq ..
type ListActivitasSalesmenReq struct {
	Keyword     string `json:"keyword" example:"SO TIMUR Indomarco"`
	SalesTypeID string `json:"sales_type_id" example:"1"`
	PeriodFrom  string `json:"period_from" example:"2021-02-05"`
	PeriodTo    string `json:"period_to" example:"2021-02-10"`
	Page        int64  `json:"page" example:"1"`
}

// DetailActivitasSalesmenReq ..
type DetailActivitasSalesmenReq struct {
	ID   string `json:"id" example:"9957"`
	Date string `json:"date" example:"2020-11-20"`
}

// DetailListActivitasTodolistReq ..
type DetailListActivitasTodolistReq struct {
	ID         string `json:"id" example:"9957"`
	Date       string `json:"date" example:"2020-11-20"`
	Keyword    string `json:"keyword" example:"OP1B00027785"`
	CategoryID int    `json:"category" example:"2"`
	Status     string `json:"status" example:"Not Exist"`
	Page       int64  `json:"page" example:"1"`
}

//CallPlanMerchant
type CallPlanMerchant struct {
	ID               int              `gorm:"column:id" json:"id"`
	MerchantName     string           `gorm:"column:merchant_name" json:"merchant_name"`
	MID              string           `gorm:"column:mid" json:"mid"`
	OwnerName        string           `gorm:"column:owner_name" json:"owner_name"`
	MerchantTypeName string           `gorm:"column:merchant_type_name" json:"merchant_type_name"`
	Address          string           `gorm:"column:address" json:"address"`
	Longitude        string           `gorm:"column:longitude" json:"longitude"`
	Latitude         string           `gorm:"column:latitude" json:"latitude"`
	MerchantStatus   string           `gorm:"column:merchant_status" json:"merchant_status"`
	CallPlanDate     string           `gorm:"column:call_plan_date" json:"call_plan_date"`
	ActionDate       string           `gorm:"column:action_date" json:"action_date"`
	ClockTime        string           `gorm:"column:clock_time" json:"clock_time"`
	Status           string           `gorm:"column:status" json:"status"`
	Notes            string           `gorm:"column:notes" json:"notes"`
	PhotoLocation    string           `gorm:"column:photo_location" json:"photo_location"`
	CallPlanActions  []CallPlanAction `gorm:"has_many:call_plan_actions;foreignKey:CallPlanMerchantID" json:"call_plan_actions"`
}

//CallPlanAction
type CallPlanAction struct {
	ID                 int     `gorm:"column:id" json:"id"`
	CallPlanMerchantID int     `gorm:"column:call_plan_merchant_id" json:"call_plan_merchant_id"`
	Action             string  `gorm:"column:action" json:"action"`
	Product            string  `gorm:"column:product" json:"product"`
	Result             bool    `gorm:"column:result" json:"result"`
	MerchantAction     string  `gorm:"column:merchant_action" json:"merchant_action"`
	Amount             float64 `gorm:"column:amount" json:"amount"`
	Note               string  `gorm:"column:note" json:"note"`
}

//TodoListDetail
type TodoListDetail struct {
	ID                      int              `gorm:"column:id" json:"id"`
	MerchantNewRecruitmenID int              `gorm:"column:merchant_new_recruitment_id" json:"merchant_new_recruitment_id"`
	MerchantName            string           `gorm:"column:merchant_name" json:"merchant_name"`
	MID                     string           `gorm:"column:mid" json:"mid"`
	OwnerName               string           `gorm:"column:owner_name" json:"owner_name"`
	Address                 string           `gorm:"column:address" json:"address"`
	ActionBy                 string           `gorm:"column:action_by" json:"action_by"`
	Notes                   string           `gorm:"column:notes" json:"notes"`
	MerchantTypeName        string           `gorm:"column:merchant_type_name" json:"merchant_type_name"`
	TaskDate                string           `gorm:"column:task_date" json:"task_date"`
	ActionDate              string           `gorm:"column:action_date" json:"action_date"`
	Longitude               string           `gorm:"column:longitude" json:"longitude"`
	Latitude                string           `gorm:"column:latitude" json:"latitude"`
	Status                  string           `gorm:"column:status" json:"status"`
	CreatedAt               string           `gorm:"column:created_at" json:"created_at"`
	Tasks                   []TodolistTask   `gorm:"has_many:tasks;foreignKey:TodolistID" json:"tasks"`
	TodoListHistories       []TodolistHistories   `gorm:"has_many:todolist_histories;foreignKey:TodolistID" json:"todolist_histories"`
	TodolistCategoryID      int              `gorm:"column:todolist_category_id" json:"todolist_category_id"`
	TodolistCategory        TodolistCategory `gorm:"references:TodolistCategoryID" json:"todolist_category"`
}

// TodolistTask ..
type TodolistTask struct {
	ID                    int                 `gorm:"column:id" json:"id"`
	ActionDate            time.Time           `gorm:"column:action_date" json:"action_date"`
	ActionBy              string              `gorm:"column:action_by" json:"action_by"`
	ActionByName          Sales               `gorm:"foreignKey:action_by;references:PhoneNumber" json:"action_by_name"`
	SupplierName          string              `gorm:"column:supplier_name" json:"supplier_name"`
	TodolistID            uint                `gorm:"column:todolist_id" json:"todolist_id"`
	TodolistSubCategoryID uint                `gorm:"column:todolist_sub_category_id" json:"todolist_sub_category_id"`
	TodolistSubCategory   TodolistSubCategory `gorm:"references:TodolistSubCategoryID" json:"todolist_sub_category"`
	FileEdukasi           string              `gorm:"column:id" json:"file_edukasi"`
	FollowUps             []FollowUp          `gorm:"has_many:follow_ups;foreignKey:TaskID" json:"follow_ups"`
	LabelTask             TodolistLabelTask   `gorm:"foreignKey:todolist_sub_category_id;references:SubCategoryID" json:"label_tasks"`
}

// TodolistLabelTask ..
type TodolistLabelTask struct {
	ID            int    `gorm:"column:id" json:"id"`
	Name          string `json:"name"`
	LabelType     string `json:"label_type"`
	SubCategoryID int    `gorm:"primary_key;column:sub_category_id" json:"sub_category_id"`
}

//TodolistHistories
type TodolistHistories struct {
	ID            int    `gorm:"column:id" json:"id"`
	TodolistID    uint   `gorm:"column:todolist_id" json:"todolist_id"`
	Description   string `json:"description"`
	Status     	  string `json:"status"`
	FotoLocation  string `json:"foto_location"`
}

// Sales ..
type Sales struct {
	Id          int    `gorm:"column:id" json:"id"`
	PhoneNumber string `gorm:"primary_key;column:phone_number" json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
}

// TableName ..
func (t *TodolistTask) TableName() string {
	return "tasks"
}

// TableName ..
func (t *Sales) TableName() string {
	return "salesmen"
}

// TableName ..
func (t *TodolistLabelTask) TableName() string {
	return "label_tasks"
}

// TableName ..
func (t *TodolistHistories) TableName() string {
	return "todolist_histories"
}
