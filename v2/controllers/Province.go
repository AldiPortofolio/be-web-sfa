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

// Province - List godoc
// @Summary Province - List
// @Description Province - List
// @ID Province - List
// @Tags List
// @Router /ottosfa/v2/province/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.Provinces} "Province - List Response EXAMPLE"
func Province(ctx *gin.Context) {
	fmt.Println(">>> Province Controller <<<")

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
	log.Info("Province Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).Province(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Province Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
