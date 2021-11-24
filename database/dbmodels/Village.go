package dbmodels

// Village ..
type Village struct {
	DatabaseModel
	Name        string `json:"name"`
	VillageCode string `json:"village_code"`
	DistrictId  uint   `json:"district_id"`
}

// VillageByDistrict ..
type VillageByDistrict struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// TableName ..
func (t *Village) TableName() string {
	return "public.villages"
}

// TableName ..
func (t *VillageByDistrict) TableName() string {
	return "public.villages"
}
