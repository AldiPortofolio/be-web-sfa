package services

import (
	"fmt"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/hosts/rose"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MerchantType ..
func (svc *Service) MerchantType(res *models.Response) {
	fmt.Println(">>> MerchantType - Service <<<")

	merchantType := ottoutils.GetEnv("LOOKUP_MERCHANT_TYPE", "TIPE_MERCHANT")
	dataRose, errRose := rose.LookUpGroup(merchantType)
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
