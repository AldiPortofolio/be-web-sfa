package gender

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	gender "ottosfa-api-web/v2.3/services/gender"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Gender - list godoc
// @Summary Gender - List
// @Description Gender - List
// @ID Gender - List
// @Tags Gender V2.3
// @Router /ottosfa/v2.3/gender/ [GET]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func List(ctx *gin.Context) {
	fmt.Println(">>> Gender - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	gender.InitiateServiceGender(log).List(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("Gender-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
