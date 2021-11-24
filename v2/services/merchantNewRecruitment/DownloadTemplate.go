package merchantNewRecruitment

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// DownloadTemplate ..
func (svc *ServiceMerchantNewRecruitment) DownloadTemplate(token string, res *models.Response) {
	fmt.Println(">>> DownloadTemplate - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TemplateMerchantNewRecruitment()
	if err != nil {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.download.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("merchant.new.recruitment.download.success")
	res.Data = data
	res.Meta.Message = "Success"
	res.Meta.Status = true

	return
}
