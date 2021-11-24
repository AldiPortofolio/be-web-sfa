package dbmodels

// SalesAreaChannelsSubArea ..
type SalesAreaChannelsSubArea struct {
	SalesAreaChannelID int64 `gorm:"column:sales_area_channel_id" json:"sales_area_channel_id" example:"1"`
	SubAreaID          int64 `gorm:"column:sub_area_id" json:"sub_area_id" example:"234"`
}
