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

// SalesRetail - List godoc
// @Summary SalesRetail - List
// @Description SalesRetail - List
// @ID SalesRetail - List
// @Tags List
// @Router /ottosfa/v2/sales-retail/list [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func SalesRetail(ctx *gin.Context) {
	fmt.Println(">>> SalesRetail Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("SalesRetail Controller")

	services.InitiateService(log).SalesRetail(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("SalesRetail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
