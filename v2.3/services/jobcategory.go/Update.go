package jobcategory

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Update ..
func (svc *ServiceJobCategories) Update(token string, req dbmodels.JobCategories, res *models.Response) {
	fmt.Println(">>> Update - ServiceJobCategories <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	err = svc.Database.SaveJobCategories(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("job.categories.update.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("job.categories.update.success")
	res.Meta.Message = "success"
	res.Meta.Status = true

}
