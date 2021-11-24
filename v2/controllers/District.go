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

// District - List godoc
// @Summary District - List
// @Description District - List
// @ID District - List
// @Tags List
// @Router /ottosfa/v2/district/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.District} "District - List Response EXAMPLE"
func District(ctx *gin.Context) {
	fmt.Println(">>> District Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	cityId := ctx.DefaultQuery("city_id", "17")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.SearchReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if cityId == "" {
		go log.Error(fmt.Sprintf("Body request error"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("District Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).District(cityId, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("District Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
