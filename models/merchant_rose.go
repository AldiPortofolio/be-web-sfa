package models

// MerchantDetailReq ..
type MerchantDetailReq struct {
	MerchantID string `json:"merchant_id"`
}

// MerchantDetailV2 ..
type MerchantDetailV2 struct {
	ID               int    `json:"id"`
	MerchantName     string `json:"merchant_name"`
	MerchantPhone    string `json:"merchant_phone"`
	MerchantID       string `json:"merchant_id"`
	OwnerName        string `json:"owner_name"`
	SubArea          string `json:"sub_area"`
	Address          string `json:"address"`
	Note             string `json:"note"`
	VillageID        string `json:"village_id"`
	AddressBenchmark string `json:"address_benchmark"`
	CustomerCode     string `json:"partner_customer_id"`
	SalesTypeID      int    `json:"sales_type_id"`
}

// MerchantList ..
type MerchantList struct {
	MerchantID  string `json:"merchant_id"`
	PhoneNumber string `json:"phone_number"`
	Name        string `json:"name"`
}

// MerchantListReq ..
type MerchantListReq struct {
	Keyword string `json:"keyword"`
}
