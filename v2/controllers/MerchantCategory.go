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

// MerchantCategory - List godoc
// @Summary MerchantCategory - List
// @Description MerchantCategory - List
// @ID MerchantCategory - List
// @Tags List
// @Router /ottosfa/v2/merchant-category/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func MerchantCategory(ctx *gin.Context) {
	fmt.Println(">>> MerchantCategory Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("MerchantCategory Controller")

	services.InitiateService(log).MerchantCategory(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantCategory Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
