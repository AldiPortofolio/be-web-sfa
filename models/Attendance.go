package models

import (
	"ottosfa-api-web/database/dbmodels"
	"time"
)

// AttendanceReq ..
type AttendanceReq struct {
	ID         string `json:"id"`
	SalesPhone string `json:"sales_phone"`
	SalesName  string `json:"sales_name"`
	Category   string `json:"category"`
	Type       string `json:"type"`
	DateFrom   string `json:"date_from"`
	DateTo     string `json:"date_to"`
	Notes      string `json:"notes"`
	Status     string `json:"status"`
	Page       int64  `json:"page"`
	Limit      int    `json:"limit"`
	Keyword    string `json:"keyword"`
}

// AttendanceRes ..
type AttendanceRes struct {
	ID                 uint      `json:"id"`
	SalesID            int       `json:"sales_id"`
	SalesPhone         string    `json:"sales_phone"`
	SalesName          string    `json:"sales_name"`
	ClocktimeServer    time.Time `json:"date"`
	AttendCategory     string    `json:"attend_category"`
	AttendCategoryType string    `json:"attend_category_type"`
	TypeAttendance     string    `json:"type_attendance"`
	Notes              string    `json:"notes"`
	Status             int       `json:"status"`
	StatusName         string    `json:"status_name"`
}

// SalesmanResponse ..
type SalesmanResponse struct {
	ID          uint                 `json:"id"`
	FirstName   string               `json:"first_name"`
	LastName    string               `json:"last_name"`
	Email       string               `json:"email"`
	IDNumber    string               `json:"id_number"`
	Dob         string               `json:"dob"`
	PhoneNumber string               `json:"phone"`
	Gender      string               `json:"gender"`
	CompanyCode string               `json:"company_code"`
	PhoneArea   string               `json:"phone_area"`
	Address     string               `json:"address"`
	Postcode    string               `json:"postcode"`
	BirthPlace  string               `json:"birth_place"`
	Status      string               `json:"status"`
	IDCardPic   string               `json:"id_card_pic"`
	Photo       string               `json:"photo"`
	SalesID     string               `json:"sales_id"`
	SfaID       string               `json:"sfa_id"`
	Occupation  string               `json:"occupation"`
	WorkDate    string               `json:"work_date"`
	SalesType   string               `json:"sales_type"`
	SalesLevel  string               `json:"sales_level"`
	SalesTypes  []dbmodels.SalesType `json:"sales_types"`
	CreatedAt   time.Time            `json:"created_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Positions   []PositionsResponse  `json:"positions"`
}

// PositionsResponse ..
type PositionsResponse struct {
	ID           uint   `json:"id"`
	RoleID       uint   `json:"role_id"`
	Role         string `json:"role"`
	RegionID     uint   `json:"region_id"`
	Region       string `json:"region"`
	BranchID     uint   `json:"branch_id"`
	Branch       string `json:"branch"`
	BranchOffice string `json:"branch_office"`
	AreaID       uint   `json:"area_id"`
	Area         string `json:"area"`
	SubAreaID    uint   `json:"sub_area_id"`
	SubArea      string `json:"sub_area"`
}

// ParameterConfiguration ..
type ParameterConfiguration struct {
	ID         uint
	Name       string
	ParamValue string
}

// ValidateAttendanceReq ..
type ValidateAttendanceReq struct {
	AttendanceId string `json:"attendance_id"`
	StatusBefore int    `json:"status_before"`
	StatusAfter  int    `json:"status_after"`
	Reason       string `json:"reason"`
}

// ExportAttendance ..
type ExportAttendance struct {
	No                     string `csv:"No"`
	AttendanceID           string `csv:"ID"`
	SalesID                string `csv:"Sales ID"`
	SalesPhone             string `csv:"Sales Phone"`
	SalesName              string `csv:"Sales Name"`
	SalesDepartment        string `csv:"Sales Department"`
	SalesAssignmentCode    string `csv:"Sales Assignment Code"`
	SalesAssignmentName    string `csv:"Sales Assignment Name"`
	SalesAssignmentCities  string `csv:"Sales Assignment Cities"`
	CompanyCode            string `csv:"Company Code"`
	AttendanceCategory     string `csv:"Attendance Category"`
	AttendanceCategoryType string `csv:"Attendance Category Type"`
	ClockTimeServer        string `csv:"Date & Time Server"`
	ClockTimeLocal         string `csv:"Date & Time Local"`
	TypeAttendance         string `csv:"Type Attendance"`
	Selfie                 string `csv:"Selfie"`
	Location               string `csv:"Location"`
	Long                   string `csv:"Long"`
	Lat                    string `csv:"Lat"`
	Notes                  string `csv:"Notes"`
	BranchOffices          string `csv:"Branch Offices"`
	PhotoProfile           string `csv:"Photo Profile"`
	PhotoAccuration        string `csv:"Photo Accuration"`
	MinAccuration          string `csv:"Min Accuration"`
	AccurationStatus       string `csv:"AccurationStatus"`
	StatusName             string `csv:"Status"`
}

// AttendanceDetailRes ..
type AttendanceDetailRes struct {
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
	Region 			   []string	 `json:"region"`
	SubAreaChanel 	   []string	 `json:"sub_area_channel"`
	SalesType 		   string  	 `json:"sales_type"`
	UpdatedAt		   time.Time `json:"updated_at"`
}
