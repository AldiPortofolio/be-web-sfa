package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/models"
	"ottosfa-api-web/v2/services"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// HealthCheck ..
// Health Check godoc
// @Summary Health Check
// @Description Health Check
// @ID Health Check
// @Tags OTTO SFA WEB
// @Router /ottosfa/v2/health-check [get]
// @Accept json
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func HealthCheck(ctx *gin.Context) {
	fmt.Println(">>> HealthCheck Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{}

	log.Info("HealthCheck Controller")

	services.InitiateService(log).HealthCheck(&res)

	resBytes, _ := json.Marshal(res)
	log.Info("HealthCheck Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
