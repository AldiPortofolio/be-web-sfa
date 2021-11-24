package gender

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

func (svc *ServiceGender) List(res *models.Response) {
	fmt.Println(">>> Detail - ServiceGender <<<")

	data := []dbmodels.Gender{
		{ID: 0, Name: "male"},
		{ID: 1, Name: "female"},
	}

	res.Meta = utils.GetMetaResponse("todolist.detail.success")
	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
