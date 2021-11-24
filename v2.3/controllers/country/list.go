package country

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	country "ottosfa-api-web/v2.3/services/country"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Country - list godoc
// @Summary Country - List
// @Description Country - List
// @ID Country - List
// @Tags Country V2.3
// @Router /ottosfa/v2.3/country/ [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func List(ctx *gin.Context) {
	fmt.Println(">>> Country - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	country.InitiateServiceCountry(log).List(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("Country-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
