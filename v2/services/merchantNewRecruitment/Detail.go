package merchantNewRecruitment

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Detail ..
func (svc *ServiceMerchantNewRecruitment) Detail(token string, id string, res *models.Response) {
	fmt.Println(">>> Detail - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.MerchantNewRecruitmentDetail(id)
	if err != nil {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.detail.failed")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
