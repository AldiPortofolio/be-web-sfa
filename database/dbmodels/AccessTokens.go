package dbmodels

import (
	"github.com/jinzhu/gorm"
)

// AccessTokens ..
type AccessTokens struct {
	gorm.Model
	AdminID  uint   `json:"admin_id"`
	ValidFor int64  `json:"valid_for"`
	Token    string `json:"token"`
	Revoked  bool   `json:"Revoked"`
}

// TableName ..
func (t *AccessTokens) TableName() string {
	return "access_tokens"
}
