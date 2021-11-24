package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// NewMerchantDetail ..
func (svc *ServiceTodolist) NewMerchantDetail(token string, req models.NewMerchantDetailReq, res *models.Response) {
	fmt.Println(">>> NewMerchantDetail - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TodolistNewMerchantDetail(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.merchant.detail.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.merchant.detail.success")
	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true

	return
}
