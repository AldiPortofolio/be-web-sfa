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

// MerchantBusinessType - List godoc
// @Summary MerchantBusinessType - List
// @Description MerchantBusinessType - List
// @ID MerchantBusinessType - List
// @Tags List
// @Router /ottosfa/v2/merchant-business-type/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func MerchantBusinessType(ctx *gin.Context) {
	fmt.Println(">>> MerchantBusinessType Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("MerchantBusinessType Controller")

	services.InitiateService(log).MerchantBusinessType(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantBusinessType Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
