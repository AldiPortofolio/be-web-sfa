package adminSubArea

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
)

func (svc *ServicAdminSubArea) DeleteByAdmin(token string, res *models.Response) {
	fmt.Println(">>> Delete - ServiceAdminSubArea <<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	admin_id := strconv.Itoa(int(adminID))
	
	errs := svc.Database.DeleteByAdmin(admin_id)

	if errs != nil {
		res.Meta.Status = false
		res.Meta.Code = 422
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200
}
