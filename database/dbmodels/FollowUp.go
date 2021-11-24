package dbmodels

// FollowUp ..
type FollowUp struct {
	ID          uint   `json:"id"`
	Label       string `json:"label"`
	ContentType string `json:"content_type"`
	Body        string `json:"body"`
	TaskID      uint   `json:"task_id"`
}
