package models

import "ottosfa-api-web/database/dbmodels"

// ReqCreateJobManagement
type ReqCreateJobManagement struct {
	Id            int64  `json:"id"`
	Name          string `json:"name" example:"test job 1"`
	JobCategoryId int64  `json:"jobCategoryId" example:"1"`
	// SenderId        int64                      `json:"senderId" example:"3"`
	RecipientId    []int64 `json:"recipientId" example:"[4,5,6]"`
	Deadline       string  `json:"deadline" example:"2021-10-30"`
	JobPriority    string  `json:"jobPriority" example:"High"`
	Status         int     `json:"status" example:"0"`
	AssignmentDate string  `json:"assignmentDate" example:"2021-10-25"`
	StatusStorage  bool    `json:"statusStorage" example:"false"`
	// CreatedAt       string                     `json:"createdAt" example:"2021-10-25"`
	// UpdatedAt       string                     `json:"updatedAt" example:"2021-10-25"`
	// AcceptDate      string                     `json:"acceptDate" example:"2021-10-25"`
	// DeliverDate     string                     `json:"deliverDate" example:"2021-10-25"`
	// CompleteDate    string                     `json:"completeDate" example:"2021-10-25"`
	// CancelDate      string                     `json:"cancelDate" example:"2021-10-25"`
	// ResendDate      string                     `json:"resendDate" example:"2021-10-25"`
	JobDescriptions []dbmodels.JobDescriptions `json:"jobDescriptions" gorm:"Foreignkey:job_management_id;association_foreignkey:ID;"`
}

type ReqBulkUploadJobManagement struct {
	No              string `json:"no" csv:"No"`
	Name            string `json:"name" csv:"Nama Tugas"`
	JobCategory     string `json:"jobCategory" csv:"Kategori Tugas"`
	RecipientEmail  string `json:"recipientId" csv:"Penerima Tugas (E-mail)"`
	JobPriority     string `json:"jobPriority" csv:"Prioritas"`
	Deadline        string `json:"deadline" csv:"Tanggal Selesai"`
	DeskripsiTugas  string `json:"deskripsiTugas" csv:"Deskripsi Tugas"`
	LabelAttachment string `json:"labelAttachment" csv:"Label Attachment"`
	LinkAttachment  string `json:"linkAttachment" csv:"Link Attachment"`
	Keterangan      string `json:"keterangan" csv:"Keterangan"`
}

// ReqEditJobManagement
type ReqEditJobManagement struct {
	Id int64 `json:"id" example:"1"`
	Name            string                     `json:"name"`
	JobCategoryId   int64                      `json:"jobCategoryId"`
	SenderId    int64 `json:"senderId"`
	RecipientId int64 `json:"recipientId"`
	Deadline        string                     `json:"deadline"`
	JobPriority     string                     `json:"jobPriority"`
	Status int `json:"status" example:"3"`
	// AssignmentDate  string                     `json:"assignmentDate"`
	// StatusStorage   bool                       `json:"statusStorage"`
	// CreatedAt       string                     `json:"createdAt"`
	// UpdatedAt       string                     `json:"updatedAt" example:"2021-10-25"`
	AcceptDate      string                     `json:"acceptDate" example:"2021-10-25"`
	DeliverDate     string                     `json:"deliverDate" example:"2021-10-25"`
	CompleteDate    string                     `json:"completeDate" example:"2021-10-25"`
	CancelDate      string                     `json:"cancelDate" example:"2021-10-25"`
	ResendDate      string                     `json:"resendDate" example:"2021-10-25"`
	Reason          string                     `json:"reason" example:"dibatalkan karena salah tanggal"`
	JobDescriptions []dbmodels.JobDescriptions `json:"jobDescriptions" gorm:"Foreignkey:job_management_id;association_foreignkey:ID;"`
	StatusStorage   bool                       `json:"statusStorage" example:"false"`
}

// ReqFilterJobManagements ..
type ReqFilterJobManagements struct {
	ID            int64  `json:"id" example:"1"`
	Name          string `json:"name" example:"Kategori A"`
	JobCategoryId int64  `json:"jobCategoryId" example:"1"`
	DateStart     string `json:"dateStart" example:"2021-10-19"`
	JobPriority   string `json:"jobPriority" example:"High"`
	Status        int    `json:"status" example:"3"`
	StatusStorage bool   `json:"statusStorage" example:"false"`
	DateEnd       string `json:"dateEnd" example:"2021-10-19"`
	Limit         int    `json:"limit" example:"10"`
	Page          int    `json:"page"  example:"1"`
	SenderId      int64  `json:"senderId"`
	RecipientId   int64  `json:"recipientId"`
}

// ReqFilterJobManagementDraft ..
type ReqFilterJobManagementDraft struct {
	Keyword string `json:"keyword" example:"Kategory A"`
	AdminID string `json:"admin_id" example:"2"`
	Page    int64  `json:"page" example:"1"`
}

type ResFilterJobManagements struct {
	JobManagements []dbmodels.JobManagements `json:"data"`
	Total          int                       `json:"total"`
}
