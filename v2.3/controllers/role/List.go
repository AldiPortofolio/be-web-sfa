package role

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	role "ottosfa-api-web/v2.3/services/role"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Role - list godoc
// @Summary Role - List
// @Description Role - List
// @ID Role - List
// @Tags Role V2.3
// @Router /ottosfa/v2.3/admin/role/list [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func List(ctx *gin.Context) {
	fmt.Println(">>> Role - List - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	reqToken := ctx.Request.Header.Get("Authorization")
	if reqToken == "" {
		ctx.JSON(http.StatusUnauthorized, "Unauthorized")
		return
	}

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	role.InitiateServiceRole(log).List(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Role-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
