package paramConfiguration

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services/paramConfiguration"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// MinMaxPhone ..
// ParamConfiguration - MinMaxPhone godoc
// @Summary ParamConfiguration - MinMaxPhone
// @Description ParamConfiguration - Get Min Max Phone
// @ID ParamConfiguration - MinMaxPhone
// @Tags Param Config
// @Router /ottosfa/v2/param-configuration/min-max-phone [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response{} "ParamConfiguration - MinMaxPhone EXAMPLE"
func MinMaxPhone(ctx *gin.Context) {
	fmt.Println(">>> ParamConfiguration - MinMaxPhone - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)
	token := utils.GetToken(ctx.Request)
	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	paramConfiguration.InitiateServiceParamConfiguration(log).MinMaxPhone(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("ParamConfiguration-MinMaxPhone Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
