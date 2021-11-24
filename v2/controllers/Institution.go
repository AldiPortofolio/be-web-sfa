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

// Institution - List godoc
// @Summary Institution - List
// @Description Institution - List
// @ID Institution - List
// @Tags List
// @Router /ottosfa/v2/institution/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
func Institution(ctx *gin.Context) {
	fmt.Println(">>> Institution Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.SearchReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Institution Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).Institution(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Institution Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
