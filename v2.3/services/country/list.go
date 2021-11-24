package country

import (
	"fmt"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

func (svc *ServiceCountry) List(res *models.Response) {
	fmt.Println(">>> Detail - ServiceCountry <<<")

	data, err := svc.Database.CountryList()
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.detail.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.detail.success")
	res.Data = map[string]interface{}{"countries": data}
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
