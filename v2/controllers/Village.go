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

// Village - List godoc
// @Summary Village - List
// @Description Village - List
// @ID Village - List
// @Tags List
// @Router /ottosfa/v2/village/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.Village} "Village - List Response EXAMPLE"
func Village(ctx *gin.Context) {
	fmt.Println(">>> Village Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	districtId := ctx.DefaultQuery("district_id", "17")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.SearchReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if districtId == "" {
		go log.Error(fmt.Sprintf("Body request error"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Village Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).Village(districtId, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Village Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
