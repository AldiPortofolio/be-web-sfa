package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// NewMerchantList ..
func (svc *ServiceTodolist) NewMerchantList(token string, req models.NewMerchantListReq, res *models.Response) {
	fmt.Println(">>> NewMerchantList - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TodolistNewMerchantList(req.Keyword)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.merchant.list.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.merchant.list.success")
	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true

	return
}
