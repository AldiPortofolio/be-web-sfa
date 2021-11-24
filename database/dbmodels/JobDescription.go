package dbmodels

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

// JobDescriptions ..
type JobDescriptions struct {
	Id              int64  `json:"id"`
	JobManagementId int64  `json:"jobManagementId"`
	Description     string `json:"description"`
	Note            string `json:"note"`
	DescLink        JsonB  `json:"descLink"`
	NoteLink        JsonB  `json:"noteLink"`
}

// TableName ..
func (t *JobDescriptions) TableName() string {
	return "public.job_descriptions"
}

// DescLink
type JsonB []map[string]string

// Make the Attrs struct implement the driver.Valuer interface. This method
// simply returns the JSON-encoded representation of the struct.
func (a JsonB) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Make the Attrs struct implement the sql.Scanner interface. This method
// simply decodes a JSON-encoded value into the struct fields.
func (a *JsonB) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &a)
}

type Link struct {
	FileName string `json:"fileName"`
	Url  string `json:"url"`
	Size string `json:"size"`
}
