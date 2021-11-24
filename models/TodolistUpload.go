package models

// BulkTodolist ..
type BulkTodolist struct {
	MerchantPhoneNo string `csv:"Merchant Phone No"`
	MerchantID      string `csv:"MID"`
	CustomerCode    string `csv:"Id Pelanggan"`
	CategoryID      string
	SubCategoryID   string `csv:"Sub-CategoryID"`
	TaskDate        string `csv:"Task Date"`
	SalesPhone      string `csv:"Sales Phone"`
	Notes           string
	SupplierName    string `csv:"Supplier Name"`
}

// DataErrorByRow ..
type DataErrorByRow struct {
	NoRow           int    `csv:"Nomor Row"`
	MerchantPhoneNo string `csv:"Merchant Phone No"`
	MerchantID      string `csv:"MID"`
	CategoryID      string
	SubCategoryID   string `csv:"Sub-CategoryID"`
	TaskDate        string `csv:"Task Date"`
	Notes           string
	SupplierName    string `csv:"Supplier Name"`
	SalesPhone      string `csv:"SalesPhone"`
	ErrorMessages   string `csv:"Error Messages"`
}

// DataMerchant ..
type DataMerchant struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	MerchantID   string `json:"merchant_id"`
	VillageID    int64  `json:"village_id"`
	PhoneNumber  string `json:"phone_number"`
	CustomerCode string `json:"customer_code"`
}

// BulkTask ..
type BulkTask struct {
	SubCategories []string `json:"sub_categories"`
	SupplierName  string   `json:"supplier_name"`
}
