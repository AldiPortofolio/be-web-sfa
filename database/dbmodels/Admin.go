package dbmodels

// Admin ..
type Admin struct {
	Id             int64  `json:"id"`
	Email          string `json:"email"`
	FirstName      string `json:"firstName"`
	LastName       string `json:"lastName"`
	AssignmentRole string `json:"assignmentRole"`
}

// TableName ..
func (t *Admin) TableName() string {
	return "admins"
}
