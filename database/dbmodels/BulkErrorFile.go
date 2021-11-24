package dbmodels

import "github.com/jinzhu/gorm"

// BulkErrorFile models
type BulkErrorFile struct {
	gorm.Model
	ErrorFile string `json:"error_file"`
	BulkType  string `json:"bulk_type"`
	Message   string `json:"message"`
}

// TableName ..
func (t *BulkErrorFile) TableName() string {
	return "public.bulk_error_files"
}
