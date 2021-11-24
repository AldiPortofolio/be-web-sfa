package jobmanagements

import (
	"fmt"
	"math"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// MerchantDetail ..
func (svc *ServiceJobManagements) FilterJobManagements(token string, req models.ReqFilterJobManagements, res *models.ResponsePagination ) {
	fmt.Println(">>> FilterJobManagements - ServiceJobManagements <<<")
	userId , err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}
	req.SenderId = int64(userId)
	req.RecipientId = int64(userId)
	data, total, err := svc.Database.FilterJobManagement(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("failed")
		res.Message= err.Error()
		return
	}

	// jobList := models.ResFilterJobManagements{
	// 	JobManagements: data,
	// 	Total: total,
	// }

	meta := models.MetaPagination{
		CurrentPage: int64(req.Page),
		NextPage:    int64(req.Page + 1),
		PrevPage:    int64(req.Page - 1),
		TotalPages:  int64(math.Ceil(float64(int64(total)) / float64(25))),
		TotalCount:  int64(total),
	}

	res.Data = data
	res.Meta = meta
	res.Message = "success"
	res.ResponseCode = "200"

}
