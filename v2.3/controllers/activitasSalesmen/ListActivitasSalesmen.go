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

// Activitas Salesmen godoc
// @Summary Activitas Salesmen - List
// @Description Activitas Salesmen - List
// @Tags Activitas Salesmen V2.3
// @ID Activitas Salesmen - List
// @Accept json
// @Produce json
// @Param Body body dbmodels.ListActivitasSalesmenReq{} true "Body"
// @Success 200 {object} models.ResponsePagination
// @Failure 400 {object} models.ResponsePagination
// @Router /ottosfa/v2.3/activitas-salesmen/list [post]
func ListActivitasSalesmen(ctx *gin.Context) {
	fmt.Println(">>> Activitas Salesmen - list - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.ResponsePagination{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.ListActivitasSalesmenReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	activitasSalesmen.InitiateServicActivitasSalesmen(log).ListActivitasSalesmen(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Activitas Salesmen-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
