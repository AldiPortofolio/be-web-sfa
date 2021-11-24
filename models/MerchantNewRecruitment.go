package models

import "ottosfa-api-web/database/dbmodels"

// MerchantNewRecruitmentListReq ..
type MerchantNewRecruitmentListReq struct {
	Id                 string `json:"id,omitempty" example:"2"`
	Name               string `json:"name,omitempty" example:"Data"`
	CustomerCode       string `json:"customer_code,omitempty" example:"566543"`
	PhoneNumber        string `json:"phone_number,omitempty" example:"088999121367"`
	InstitutionCode    string `json:"institution_code,omitempty" example:"OP"`
	SubAreaChannelName string `json:"sub_area_channel_name,omitempty" example:"0004"`
	Status             string `json:"status,omitempty" example:"Activated"`
	Page               int64  `json:"page,omitempty"`
}

// MerchantNewRecruitmentExportRes ..
type MerchantNewRecruitmentExportRes struct {
	dbmodels.MerchantNewRecruitments
	ProvinceName string `gorm:"column:province_name" json:"province_name" example:"DKI JAKARTA"`
	CityName     string `gorm:"column:city_name" json:"city_name" example:"KOTA JAKARTA SELATAN"`
	DistrictName string `gorm:"column:district_name" json:"district_name" example:"SETIA BUDI"`
	VillageName  string `gorm:"column:village_name" json:"village_name" example:"SETIA BUDI"`
}

// MerchantNewRecruitmentExportCSVRes ..
type MerchantNewRecruitmentExportCSVRes struct {
	No                 string
	Name               string `csv:"Nama Merchant"`
	OwnerName          string `csv:"Nama Pemilik"`
	CustomerCode       string `csv:"No Pelanggan"`
	PhoneNumber        string `csv:"No Hp Merchant"`
	InstitutionCode    string `csv:"Institusi"`
	Address            string `csv:"Alamat"`
	ProvinceName       string `csv:"Provinsi"`
	CityName           string `csv:"Kota/Kabupaten"`
	DistrictName       string `csv:"Kecamatan"`
	VillageName        string `csv:"Kelurahan"`
	SubAreaChannelName string `csv:"Sub Area Channel"`
	Longitude          string `csv:"Longtitute"`
	Latitude           string `csv:"Latitude"`
}
