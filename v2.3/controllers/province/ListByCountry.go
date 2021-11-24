package province

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	province "ottosfa-api-web/v2.3/services/province"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Province - list by country godoc
// @Summary Province - list by country
// @Description Province - list by country
// @ID Province - list by country
// @Tags Province V2.3
// @Router /ottosfa/v2.3/Province/ [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func ListByCountry(ctx *gin.Context) {
	fmt.Println(">>> Gender - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	countryID := ctx.Params.ByName("countryId")
	province.InitiateServiceProvince(log).ListByCountry(countryID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Gender-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
