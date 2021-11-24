package jobmanagements

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Create ..
func (svc *ServiceJobManagements) CheckAdmin(token string,  res *models.Response) {
	fmt.Println(">>> Create - ServiceJobManagements <<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	admin, _ := svc.Database.FindAdminById(int64(adminID))

	res.Data = admin
	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true
}
