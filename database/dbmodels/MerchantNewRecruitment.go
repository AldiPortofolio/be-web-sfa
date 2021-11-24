package dbmodels

// MerchantNewRecruitments ..
type MerchantNewRecruitments struct {
	DatabaseModel
	Name               string `gorm:"column:name" json:"name" example:"Test Data 2"`
	PhoneNumber        string `gorm:"column:phone_number" json:"phone_number" example:"088999121367"`
	CustomerCode       string `gorm:"column:customer_code" json:"customer_code" example:"1234566543216.0"`
	InstitutionCode    string `gorm:"column:institution_code" json:"institution_code" example:"OP"`
	SubAreaChannelID   int64  `gorm:"column:sub_area_channel_id" json:"sub_area_channel_id" example:"4"`
	SubAreaChannelName string `gorm:"column:sub_area_channel_name" json:"sub_area_channel_name" example:"Sub Area Jabar 0004"`
	OwnerName          string `gorm:"column:owner_name" json:"owner_name" example:"Namaku Siapa Ya"`
	Address            string `gorm:"column:address" json:"address" example:"Jalan Buntu no 21"`
	Longitude          string `gorm:"column:longitude" json:"longitude" example:"106.81721122"`
	Latitude           string `gorm:"column:latitude" json:"latitude" example:"-6.2354219"`
	ProvinceId         int64  `gorm:"column:province_id" json:"province_id" example:"31"`
	CityId             int64  `gorm:"column:city_id" json:"city_id" example:"3171"`
	DistrictId         int64  `gorm:"column:district_id" json:"district_id" example:"3171100"`
	VillageId          int64  `gorm:"column:village_id" json:"village_id" example:"3171100008"`
	Status             string `gorm:"column:status" json:"status" example:"Activated"`
	IdCard             string `gorm:"column:id_card" json:"id_card" example:"12345678909887654"`
	SalesTypeId        int64  `gorm:"column:sales_type_id" json:"sales_type_id" example:"1"`
}

// TableName ..
func (t *MerchantNewRecruitments) TableName() string {
	return "public.merchant_new_recruitments"
}
