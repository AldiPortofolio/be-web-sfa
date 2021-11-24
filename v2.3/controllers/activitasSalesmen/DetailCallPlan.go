package activitasSalesmen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/activitasSalesmen"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Activitas Salesmen - Detail Callplan godoc
// @Summary Activitas Salesmen -  Detail Callplan
// @Description Activitas Salesmen -  Detail Callplan
// @ID Activitas Salesmen -  Detail Callplan
// @Tags Activitas Salesmen V2.3
// @Router /ottosfa/v2.3/activitas-salesmen/detail-callplan [GET]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{} "Activitas Salesmen - Detail Sales EXAMPLE"
func DetailCallPlan(ctx *gin.Context) {
	fmt.Println(">>> Activitas Salesmen - Detail Callplan - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	callplanMerchantID := ctx.Params.ByName("callplanMerchantID")

	activitasSalesmen.InitiateServicActivitasSalesmen(log).DetailCallPlan(callplanMerchantID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Activitas Salesmen-Detail-Sales Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
