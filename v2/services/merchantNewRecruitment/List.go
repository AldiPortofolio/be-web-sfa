package merchantNewRecruitment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// List ..
func (svc *ServiceMerchantNewRecruitment) List(token string, req models.MerchantNewRecruitmentListReq, res *models.ResponsePagination) {
	fmt.Println(">>> List - ServiceMerchantNewRecruitment <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, total, err := svc.Database.MerchantNewRecruitmentList(req)
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
		res.Data = gin.H{}
		//return
	}

	meta := models.MetaPagination{
		CurrentPage: req.Page,
		NextPage:    req.Page + 1,
		PrevPage:    req.Page - 1,
		TotalPages:  int64(math.Ceil(float64(total) / float64(25))),
		TotalCount:  total,
	}

	res.Meta = meta
	res.Data = data
	res.ResponseCode = "00"
	res.Message = "success"

	return
}
