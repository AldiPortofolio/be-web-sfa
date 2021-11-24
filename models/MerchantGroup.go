package models

// MerchantGroupReq ..
type MerchantGroupReq struct {
	MerchantType string `json:"merchant_type" example:"UMKM"`
}

// MerchantGroupRes ..
type MerchantGroupRes struct {
	Id	 				int 	`json:"id" example:"111"`
	MerchantGroupName	string 	`json:"merchant_group_name" example:"bp grup learning"`
}
