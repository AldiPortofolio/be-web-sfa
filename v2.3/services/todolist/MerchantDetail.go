package todolist

import (
	"fmt"
	"ottosfa-api-web/models"
)

// MerchantDetail ..
func (svc *ServiceTodolist) MerchantDetail(req models.MerchantDetailReq, res *models.Response) {
	fmt.Println(">>> MerchantDetail - ServiceTodolist <<<")

	data, err := svc.Database.GetMerchantDetailRose(req.MerchantID)
	if err != nil {
		res.Meta.Code = 401
		res.Meta.Message = err.Error()
		res.Meta.Status = false
		return
	}

	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
