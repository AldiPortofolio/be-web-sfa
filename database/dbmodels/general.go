package dbmodels

import (
	"time"
)

// DatabaseModel ..
type DatabaseModel struct {
	Id        int       `gorm:"column:id;primary_key" json:"id" example:"1"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
