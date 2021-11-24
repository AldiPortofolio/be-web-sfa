package activitasSalesmen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/activitasSalesmen"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Activitas Salesmen - Detail Sales godoc
// @Summary Activitas Salesmen -  Detail Sales
// @Description Activitas Salesmen -  Detail Sales
// @ID Activitas Salesmen -  Detail Sales
// @Tags Activitas Salesmen V2.3
// @Router /ottosfa/v2.3/activitas-salesmen/detail-sales [POST]
// @Accept json
// @Produce json
// @Param Body body dbmodels.DetailActivitasSalesmenReq{} true "Body"
// @Success 200 {object} models.Response{} "Activitas Salesmen - Detail Sales EXAMPLE"
func DetailActivitasSalesmen(ctx *gin.Context) {
	fmt.Println(">>> Activitas Salesmen - list - Sales - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.DetailActivitasSalesmenReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	activitasSalesmen.InitiateServicActivitasSalesmen(log).DetailActivitasSalesmen(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Activitas Salesmen-Detail-Sales Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
