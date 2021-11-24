package ottosfanotif

// Response ..
type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

// Meta ...
type Meta struct {
	Status  bool   `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// ReqCreateNotif ...
type ReqCreateNotif struct {
	SenderName  string `json:"sender_name"`
	ObjectId    int64  `json:"object_id"`
	ObjectType  string `json:"object_type"`
	RecipientId int64  `json:"recipient_id"`
	SenderId    int64  `json:"sender_id"`
	Message     string `json:"message"`
	NotifCode   int    `json:"notif_code"`
}
