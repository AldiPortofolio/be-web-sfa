package dbmodels

import "time"

// AdminSubArea ..
type AdminSubArea struct {
	ID         uint   `json:"id" gorm:"column:id"`
	AdminID    uint   `json:"admin_id" gorm:"column:admin_id"`
	SubAreaIDs string `json:"sub_area_ids" gorm:"column:sub_area_ids"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
