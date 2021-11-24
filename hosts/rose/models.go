package rose

// Response ..
type Response struct {
	Rc   string      `json:"rc"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// LookUpGroupResponse ..
type LookUpGroupResponse struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	LookupGroup string `json:"lookupGroup"`
}

// UserCategoryResponse ..
type UserCategoryResponse struct {
	ID     int    `json:"id"`
	Code   string `json:"code"`
	Name   string `json:"name"`
	Logo   string `json:"logo"`
	AppId  string `json:"appId"`
	Notes  string `json:"notes"`
	Seq    int    `json:"seq"`
	Status string `json:"status"`
}

// MerchantGroupResponse ..
type MerchantGroupResponse struct {
	ID                      int    `json:"id"`
	MerchantGroupName       string `json:"merchantGroupName"`
	LogoPath                string `json:"logoPath"`
	enablePartnerCustomerId bool   `json:"v"`
}
