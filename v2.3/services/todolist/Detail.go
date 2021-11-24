package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Detail ..
func (svc *ServiceTodolist) Detail(token string, todolistID string, res *models.Response) {
	fmt.Println(">>> Detail - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TodolistDetailV2(todolistID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.detail.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.detail.success")
	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true

	return
}
