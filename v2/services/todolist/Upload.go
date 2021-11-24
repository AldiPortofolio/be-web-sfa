package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Upload ..
func (svc *ServiceTodolist) Upload(token string, fileBytes []byte, res *models.Response) {
	fmt.Println(">>> Upload - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, mErr := svc.Database.TodolistUpload(fileBytes)
	if mErr != nil {
		res.Meta = utils.GetMetaResponse("todolist.upload.failed")
		res.Meta.Message = mErr.Error()
		res.Data = data
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.upload.success")
	res.Data = data
	res.Meta.Message = "All records Todolist has been successfully Bulk Created"
	res.Meta.Status = true

	return
}
