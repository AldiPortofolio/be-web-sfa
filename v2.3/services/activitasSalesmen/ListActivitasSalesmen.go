package activitasSalesmen

import (
	"fmt"
	"math"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
)

func (svc *ServicActivitasSalesmen) ListActivitasSalesmen(token string, req dbmodels.ListActivitasSalesmenReq, res *models.ResponsePagination) {
	fmt.Println(">>> ListActivitasSalesmen - ServicActivitasSalesmen <<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	coverageArea, err := svc.Database.GetSessionCoverageArea(adminID)

	var data []dbmodels.ActivitasSalesmenList
	var total int64
	var errs error
	if coverageArea != "" {
		data, total, errs = svc.Database.ListActivitasSalesmenByCoverage(req, coverageArea)
	} else {
		data, total, errs = svc.Database.ListActivitasSalesmen(req)
	}
	
	if errs != nil {
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
