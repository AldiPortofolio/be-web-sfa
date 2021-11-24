package dbmodels

import "time"

// JobManagements ..
type JobManagements struct {
	Id              int64             `json:"id"`
	Name            string            `json:"name"`
	JobCategoryId   int64             `json:"jobCategoryId"`
	JobCategoryName JobCategories     `json:"jobCategoryName" gorm:"foreignKey:jobCategoryId;references:ID" `
	SenderId        int64             `json:"senderId"`
	SenderName      Admin             `json:"senderName" gorm:"foreignKey:senderId;references:ID"`
	RecipientId     int64             `json:"recipientId"`
	RecipientName   Admin             `json:"recipientName" gorm:"foreignKey:RecipientId;references:ID" `
	Deadline        time.Time         `json:"deadline"`
	JobPriority     string            `json:"jobPriority"`
	Status          int               `json:"status"`
	AssignmentDate  time.Time         `json:"assignmentDate"`
	StatusStorage   bool              `json:"statusStorage"`
	CreatedAt       time.Time         `json:"createdAt"`
	UpdatedAt       time.Time         `json:"updatedAt"`
	AcceptDate      time.Time         `json:"acceptDate"`
	DeliverDate     time.Time         `json:"deliverDate"`
	CompleteDate    time.Time         `json:"completeDate"`
	CancelDate      time.Time         `json:"cancelDate"`
	ResendDate      time.Time         `json:"resendDate"`
	Reason          string            `json:"reason"`
	JobDescriptions []JobDescriptions `json:"jobDescriptions" gorm:"Foreignkey:job_management_id;association_foreignkey:ID;"`
}

// JobManagements ..
type DraftJobManagements struct {
	Id              int64     `json:"id"`
	Name            string    `json:"name"`
	JobCategoryId   int64     `json:"job_category_id"`
	JobCategoryName string    `json:"job_category_name"`
	AssignmentDate  time.Time `json:"assignment_date"`
}

// TableName ..
func (t *JobManagements) TableName() string {
	return "public.job_managements"
}
