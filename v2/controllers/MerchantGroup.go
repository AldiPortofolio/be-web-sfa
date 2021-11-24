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

// MerchantGroup - List godoc
// @Summary MerchantGroup - List
// @Description MerchantGroup - List
// @ID MerchantGroup - List
// @Tags List
// @Router /ottosfa/v2/merchant-group/list [post]
// @Accept json
// @Produce json
// @Param Body body models.MerchantGroupReq true "Body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func MerchantGroup(ctx *gin.Context) {
	fmt.Println(">>> MerchantGroup Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.MerchantGroupReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("MerchantGroup Controller")
		log.AddField("RequestBody:", string(reqBytes))
	services.InitiateService(log).MerchantGroup(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantGroup Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
