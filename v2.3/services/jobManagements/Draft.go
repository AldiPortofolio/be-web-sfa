package jobmanagements

import (
	"fmt"
	"math"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (svc *ServiceJobManagements) Draft(token string, req models.ReqFilterJobManagementDraft, res *models.ResponsePagination) {
	fmt.Println(">>> Draft - ServicJobManagement<<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	req.AdminID = strconv.Itoa(int(adminID))
	data, total, err := svc.Database.ListDraftJobManagements(req)
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
		res.Data = gin.H{}
		return
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
	res.ResponseCode = "200"
	res.Message = "success"

}
