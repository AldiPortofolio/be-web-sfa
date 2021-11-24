package dbmodels

// MerchantNewRecruitmentTemplate ..
type MerchantNewRecruitmentTemplate struct {
	ID           uint   `json:"id"`
	TemplateFile string `json:"template_file"`
}

// TableName ..
func (t *MerchantNewRecruitmentTemplate) TableName() string {
	return "public.merchant_new_recruitment_templates"
}

// TodolistTemplate ..
type TodolistTemplate struct {
	ID           uint   `json:"id"`
	TemplateFile string `json:"template_file"`
}

// TableName ..
func (t *TodolistTemplate) TableName() string {
	return "public.todolist_templates"
}
