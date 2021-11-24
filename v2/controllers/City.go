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

// City - List godoc
// @Summary City - List
// @Description City - List
// @ID City - List
// @Tags List
// @Router /ottosfa/v2/city/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.Cities} "City - List Response EXAMPLE"
func City(ctx *gin.Context) {
	fmt.Println(">>> City Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	provinceId := ctx.DefaultQuery("province_id", "17")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.SearchReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if provinceId == "" {
		go log.Error(fmt.Sprintf("Body request error"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("City Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).City(provinceId, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("City Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
