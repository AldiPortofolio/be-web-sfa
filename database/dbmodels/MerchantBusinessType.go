package dbmodels

// MerchantBusinessType ..
type MerchantBusinessType struct {
	Code string `gorm:"column:code" json:"code"`
	Name string `gorm:"column:name" json:"name"`
}

// TableName ..
func (t *MerchantBusinessType) TableName() string {
	return "public.merchant_business_types"
}
