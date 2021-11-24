package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Update ..
func (svc *ServiceTodolist) Update(token string, req models.UpdateTodolist, res *models.Response) {
	fmt.Println(">>> Update - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TodolistUpdate(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.update.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.update.success")
	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true

	return
}
