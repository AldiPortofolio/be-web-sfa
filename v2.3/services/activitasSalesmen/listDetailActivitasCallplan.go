package activitasSalesmen

import (
	"fmt"
	"math"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
)

func (svc *ServicActivitasSalesmen) ListDetailActivitasSalesmenCallplan(req dbmodels.DetailListActivitasTodolistReq, res *models.ResponsePagination) {
	fmt.Println(">>> ListActivitasSalesmen - ServicActivitasSalesmen <<<")

	data, total, err := svc.Database.DetailListActivitasCallplan(req)
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
