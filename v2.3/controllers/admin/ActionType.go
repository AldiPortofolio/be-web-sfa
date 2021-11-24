package admin

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/admin"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Admin - ActionType godoc
// @Summary Admin - ActionType
// @Description Admin - ActionType
// @ID Admin - ActionType
// @Tags Admin V2.3
// @Router /ottosfa/v2.3/action-types [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func ActionTypes(ctx *gin.Context) {
	fmt.Println(">>> Company Codes - list - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	admin.InitiateServicAdmin(log).ActionTypes(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("Company-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
