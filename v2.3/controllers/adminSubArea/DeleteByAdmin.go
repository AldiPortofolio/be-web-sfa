package adminSubArea

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/adminSubArea"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// AdminSubArea - DeleteByAdmin godoc
// @Summary AdminSubArea - DeleteByAdmin
// @Description AdminSubArea - DeleteByAdmin
// @ID AdminSubArea - DeleteByAdmin
// @Tags AdminSubArea V2.3
// @Router /ottosfa/v2.3/admin-sub-area/delete-by-admin [Delete]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func DeleteByAdmin(ctx *gin.Context) {
	fmt.Println(">>> Admin Sub Area - list - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)
	token := utils.GetToken(ctx.Request)
	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	adminSubArea.InitiateServicAdminSubArea(log).DeleteByAdmin(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Admin Sub Area Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
