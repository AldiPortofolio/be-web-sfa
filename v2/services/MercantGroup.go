package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/hosts/rose"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MerchantGroup ..
func (svc *Service) MerchantGroup(req models.MerchantGroupReq, res *models.Response) {
	fmt.Println(">>> MerchantGroup - Service <<<")

	if req.MerchantType == "" {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	dataRose, errRose := rose.MerchantGroupByMerchantType(req.MerchantType)
	if errRose != nil {
		res.Meta = utils.GetMetaResponse("default")
		return
	}

	data := []models.MerchantGroupRes{}
	for _, val := range dataRose {
		a := models.MerchantGroupRes{
			Id: val.ID,
			MerchantGroupName: val.MerchantGroupName,
		}
		data = append(data, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
