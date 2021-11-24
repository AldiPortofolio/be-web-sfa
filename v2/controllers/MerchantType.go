package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services"
)

// MerchantType - List godoc
// @Summary MerchantType - List
// @Description MerchantType - List
// @ID MerchantType - List
// @Tags List
// @Router /ottosfa/v2/merchant-type/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func MerchantType(ctx *gin.Context) {
	fmt.Println(">>> MerchantType Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("MerchantType Controller")

	services.InitiateService(log).MerchantType(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantType Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
