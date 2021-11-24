package todolist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/models"

	todolistService "ottosfa-api-web/v2.3/services/todolist"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Create godoc
// @Summary Merchant Detail
// @Description Merchant Detail
// @Tags TodoList V2.3
// @ID merchant-detail
// @Accept  json
// @Produce  json
// @Param body body models.MerchantDetailReq true "request body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /ottosfa/v2.3/todolist/merchant-detail [post]
func MerchantDetail(ctx *gin.Context) {
	fmt.Println(">>> TodoList - MerchantDetail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.Response{}

	req := models.MerchantDetailReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	todolistService.InitiateServiceTodolist(log).MerchantDetail(req, &res)

	reqBytes, _ := json.Marshal(req)
	log.Info("Acquisition-Create Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	resBytes, _ := json.Marshal(res)
	log.Info("Acquisition-Create Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
