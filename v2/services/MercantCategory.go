package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/hosts/rose"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MerchantCategory ..
func (svc *Service) MerchantCategory(res *models.Response) {
	fmt.Println(">>> MerchantCategory - Service <<<")

	dataRose, errRose := rose.UserCategory()
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := []models.DataModule{}
	for _, val := range dataRose {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data = append(data, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
