package postgres

import (
	"errors"
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"strconv"
)

// TemplateTodolist ..
func (database *DbPostgres) TemplateTodolist() (interface{}, error) {
	fmt.Println(">>> Todolist - Download Template - Postgres <<<")

	var template dbmodels.TodolistTemplate
	templateFile := map[string]interface{}{}
	err := Dbcon.Last(&template).Error

	if err != nil || template.TemplateFile == "" {
		err = errors.New("Template File tidak ditemukan")
		return templateFile, err
	}

	templateFile["template"] = "/uploads/todolist_template/template_file/" + strconv.Itoa(int(template.ID)) + "/" + template.TemplateFile

	return templateFile, nil
}
