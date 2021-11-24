package company

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	company "ottosfa-api-web/v2.3/services/company"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Company - list godoc
// @Summary Company - List
// @Description Company - List
// @ID Company - List
// @Tags Company V2.3
// @Router /ottosfa/v2.3/company//list [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func List(ctx *gin.Context) {
	fmt.Println(">>> Company - Detail - Controller <<<")

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

	company.InitiateServiceCompany(log).List(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Company-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
