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

// SubAreaChannel - List godoc
// @Summary SubAreaChannel - List
// @Description SubAreaChannel - List
// @ID SubAreaChannel - List
// @Tags List
// @Router /ottosfa/v2/sub-area-channel/list [post]
// @Accept json
// @Produce json
// @Param Body body models.SearchReq true "Body"
// @Success 200 {object} models.Response{data=[]dbmodels.SubArea} "Sub Area Channel - List Response EXAMPLE"
func SubAreaChannel(ctx *gin.Context) {
	fmt.Println(">>> SubAreaChannel Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	villageId := ctx.DefaultQuery("village_id", "17")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.SearchReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if villageId == "" {
		go log.Error(fmt.Sprintf("Body request error"))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("SubAreaChannel Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	services.InitiateService(log).SubAreaChannel(villageId, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("SubAreaChannel Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
