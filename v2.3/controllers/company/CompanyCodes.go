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
// Company Codes - list godoc
// @Summary Company Codes - List
// @Description Company Codes - List
// @ID Company Codes - List
// @Tags Company Codes V2.3
// @Router /ottosfa/v2.3/company-codes [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func CompanyCodes(ctx *gin.Context) {
	fmt.Println(">>> Company Codes - list - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	company.InitiateServiceCompany(log).CompanyCodes(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("Company-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
