package dbmodels

import "time"

// SalesAreaChannel ..
type SalesAreaChannel struct {
	ID          uint
	Name        string `json:"name"`
	SalesTypeID uint   `json:"sales_type_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// SubAreas    []SubArea `gorm:"many2many:sales_area_channels_sub_areas;"`
}

// SalesAreaChannelList ..
type SalesAreaChannelList struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
