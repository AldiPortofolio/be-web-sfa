package jobcategory

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Update ..
func (svc *ServiceJobCategories) Detail(token string, ID string, res *models.Response) {
	fmt.Println(">>> Detail - ServiceJobCategories <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.DetailJobCategory(ID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("job.categories.update.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Data = data
	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true

}
