package jobcategory

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Update ..
func (svc *ServiceJobCategories) Delete(token string, categoryID int64, res *models.Response) {
	fmt.Println(">>> Delete - ServiceJobCategories <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	jobManagements, err := svc.Database.FindByCategoryID(categoryID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("job.categories.update.failed")
		res.Meta.Message = err.Error()
		return
	}

	if len(jobManagements) > 0 {
		res.Meta.Code = 400
		res.Meta.Status = false
		res.Meta.Message = "Kategori tugas sudah di gunakan"
		return
	}

	errs := svc.Database.DeleteJobCategory(categoryID)
	if errs != nil {
		res.Meta.Code = 400
		res.Meta.Status = false
		res.Meta.Message = err.Error()
		return
	}

	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true

}
