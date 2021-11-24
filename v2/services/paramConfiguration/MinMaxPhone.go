package paramConfiguration

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MinMaxPhone ..
func (svc *ServiceParamConfiguration) MinMaxPhone(token string, res *models.Response) {
	fmt.Println(">>> MinMaxPhone - ServiceParamConfiguration <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.MinMaxPhone()
	if err != nil {
		res.Meta = utils.GetMetaResponse("param.min.max.phone.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("param.min.max.phone.success")
	res.Data = data
	res.Meta.Message = "Success"
	res.Meta.Status = true

	return
}
