package models

// BulkNewRecruitment models
type BulkNewRecruitment struct {
	StoreName      string `csv:"Nama Toko"`
	Owner          string `csv:"Nama Pemilik"`
	CustomerCode   string `csv:"ID Pelanggan"`
	MerchantPhone  string `csv:"No Hp"`
	Institution    string `csv:"Institusi"`
	Address        string `csv:"Alamat"`
	VillageID      string `csv:"KelurahanID"`
	SubAreaChannel string `csv:"SubAreaChannelCode"`
	Longitude      string `csv:"Longitude"`
	Latitude       string `csv:"Latitude"`
	IDCard         string `csv:"NIK"`
	SalesType      string `csv:"Sales Type"`
}

// BulkLinkError models
type BulkLinkError struct {
	ErrorFile string `json:"error_file"`
}

// DataNewRecruitmentErrorByRow models
type DataNewRecruitmentErrorByRow struct {
	NoRow          int    `csv:"Nomor Row"`
	StoreName      string `csv:"Nama Toko"`
	Owner          string `csv:"Nama Pemilik"`
	CustomerCode   string `csv:"ID Pelanggan"`
	MerchantPhone  string `csv:"No Hp"`
	Institution    string `csv:"Institusi"`
	Address        string `csv:"Alamat"`
	VillageID      string `csv:"KelurahanID"`
	SubAreaChannel string `csv:"SubAreaChannelCode"`
	Longitude      string `csv:"Longitude"`
	Latitude       string `csv:"Latitude"`
	IDCard         string `csv:"NIK"`
	SalesType      string `csv:"Sales Type"`
	ErrorMessages  string `csv:"Error Messages"`
}

// UploadRequest ..
type UploadRequest struct {
	BucketName  string
	Data        string
	NameFile    string
	ContentType string
}
