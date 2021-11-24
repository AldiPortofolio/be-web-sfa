package dbmodels

import "time"

// Attendance ..
type Attendance struct {
	ID                 uint      `json:"id"`
	SalesID            int       `json:"sales_id"`
	SalesPhone         string    `json:"sales_phone"`
	SalesName          string    `json:"sales_name"`
	Selfie             string    `json:"selfie"`
	ClocktimeServer    time.Time `json:"date"`
	ClocktimeLocal     time.Time `json:"clocktime_local"`
	Location           string    `json:"location"`
	Latitude           string    `json:"latitude"`
	Longitude          string    `json:"longitude"`
	AttendCategory     string    `json:"attend_category"`
	AttendCategoryType string    `json:"attend_category_type"`
	TypeAttendance     string    `json:"type_attendance"`
	Notes              string    `json:"notes"`
	PhotoAccuration    string    `json:"photo_accuration"`
	PhotoProfile       string    `json:"photo_profile"`
	MinAccPercentage   string    `json:"min_accuration_percentage"`
	Status			   int		 `json:"status"`
	UpdatedAt		   time.Time `json:"updated_at"`
}