package village

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	village "ottosfa-api-web/v2.3/services/village"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Village - list by district godoc
// @Summary Village - List by district
// @Description Village - List by district
// @ID Village - List by district
// @Tags Village V2.3
// @Router /ottosfa/v2.3/village/list [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func ListByDistrict(ctx *gin.Context) {
	fmt.Println(">>> Village - List By District - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	districtID := ctx.Params.ByName("districtId")

	village.InitiateServiceVillage(log).ListByDistrict(districtID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Viilage-List-by-district Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
