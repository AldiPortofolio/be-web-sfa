package merchantNewRecruitment

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Upload ..
func (svc *ServiceMerchantNewRecruitment) Upload(token string, fileBytes []byte, res *models.Response) {
	fmt.Println(">>> Upload - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, mErr := svc.Database.MerchantNewRecruitmentUpload(fileBytes)
	if mErr != nil {
		res.Meta = utils.GetMetaResponse("merchant.new.recruitment.upload.failed")
		res.Meta.Message = mErr.Error()
		res.Data = data
		return
	}

	res.Meta = utils.GetMetaResponse("merchant.new.recruitment.upload.success")
	res.Data = data
	res.Meta.Message = "All records Merchant Recruitment has been successfully Bulk Created"
	res.Meta.Status = true

	return
}
