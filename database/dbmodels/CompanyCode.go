package dbmodels

// CompanyCode ..
type CompanyCode struct {
	ID                int    `gorm:"column:id" json:"id"`
	Name              string `gorm:"column:name" json:"name"`
	Code              int    `gorm:"column:code" json:"code"`
	EmailVerification bool   `gorm:"column:email_verification" json:"email_verification"`
}

// TableName ..
func (t *CompanyCode) TableName() string {
	return "public.company_codes"
}
