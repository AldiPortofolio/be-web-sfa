package dbmodels

import (
	"time"
)

// Salesman ..
type Salesman struct {
	ID                   uint
	FirstName            string      `json:"first_name"`
	LastName             string      `json:"last_name"`
	Email                string      `json:"email"`
	IDNumber             string      `json:"id_number"`
	Dob                  time.Time   `json:"dob"`
	PhoneNumber          string      `json:"phone_number"`
	Gender               int         `json:"gender"`
	CompanyCode          string      `json:"company_code"`
	PhoneArea            string      `json:"phone_area"`
	Address              string      `json:"address"`
	Postcode             string      `json:"postcode"`
	BirthPlace           string      `json:"birth_place"`
	Status               int         `json:"status"`
	IDCard               string      `json:"id_card"`
	Photo                string      `json:"photo"`
	SalesID              string      `json:"sales_id"`
	SfaID                string      `json:"sfa_id"`
	Occupation           string      `json:"occupation"`
	WorkDate             time.Time   `json:"work_date"`
	PasswordDigest       string      `json:"password_digest" gorm:"column:password_digest"`
	SessionToken         string      `json:"session_token"`
	ProvinceID           uint        `json:"province_id"`
	FunctionalPositionID uint        `json:"functional_position_id"`
	SalesLevelID         uint        `json:"sales_level_id"`
	SalesTypeID          uint        `json:"sales_type_id"`
	SalesTypes           []SalesType `json:"sales_types" gorm:"many2many:sales_types_salesmen;association_jointable_foreignkey:sales_type_id;jointable_foreignkey:salesman_id;"`
	Positions            []Position
	CreatedAt            time.Time
	UpdatedAt            time.Time
}

// Position ..
type Position struct {
	ID             uint      `json:"id"`
	PostCode       string    `json:"post_code"`
	RoleName       string    `json:"role_name"`
	RegionableType string    `json:"regionable_type"`
	RegionableID   uint      `json:"regionable_id"`
	SalesRoleID    uint      `json:"sales_role_id"`
	SalesmanID     uint      `json:"salesman_id"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

// GetGender ..
func GetGender(g int) string {
	gen := [2]string{"male", "female"}
	return gen[g]
}

// GetStatus ..
func GetStatus(s int) string {
	status := [5]string{"Unregistered", "Registered", "Verified", "Inactive", "Pending"}
	return status[s]
}

// SalesLevelList ..
type SalesLevelList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}