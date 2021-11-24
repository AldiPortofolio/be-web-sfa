package todolist

import (
	"fmt"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MerchantDetail ..
func (svc *ServiceTodolist) MerchantList(req models.MerchantListReq, res *models.Response) {
	fmt.Println(">>> MerchantList - ServiceTodolist <<<")

	data, err := svc.Database.GetMerchantListRose(req.Keyword)
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.create.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
