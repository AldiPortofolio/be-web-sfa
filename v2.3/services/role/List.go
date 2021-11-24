package role

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

func (svc *ServiceRole) List(token string, res *models.Response) {
	fmt.Println(">>> Detail - ServiceRole <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.RoleList()
	if err != nil {
		res.Meta.Status = false
		res.Meta.Code = 422
		res.Meta.Message = err.Error()
		return
	}

	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
