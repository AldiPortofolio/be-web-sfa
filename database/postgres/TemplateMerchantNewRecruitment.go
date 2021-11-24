package postgres

import (
	"errors"
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"strconv"
)

// TemplateMerchantNewRecruitment ..
func (database *DbPostgres) TemplateMerchantNewRecruitment() (interface{}, error) {
	fmt.Println(">>> MerchantNewRecruitment - Download Template - Postgres <<<")

	var template dbmodels.MerchantNewRecruitmentTemplate
	templateFile := map[string]interface{}{}
	err := Dbcon.Last(&template).Error

	if err != nil || template.TemplateFile == "" {
		err = errors.New("Template File tidak ditemukan")
		return templateFile, err
	}

	templateFile["template"] = "/uploads/merchant_new_recruitment_template/template_file/" + strconv.Itoa(int(template.ID)) + "/" + template.TemplateFile

	return templateFile, nil
}
