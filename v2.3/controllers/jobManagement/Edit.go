package jobManagement

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"

	// "ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	jobManagement "ottosfa-api-web/v2.3/services/jobManagements"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Edit ..
// JobManagement - Edit godoc
// @Summary JobManagement - Edit
// @Description JobManagement - Edit
// @ID JobManagement - Edit
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/edit [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.ReqEditJobManagement true "Body"
// @Success 200 {object} models.Response{} "JobManagement - Edit EXAMPLE"
func Edit(ctx *gin.Context) {
	fmt.Println(">>> JobManagement - Edit - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReqEditJobManagement{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	if req.Id == 0 {
		go log.Error("Body request error: Id Cannot 0")
		res.Meta.Code = 400
		res.Meta.Message = "Body request error: Id Cannot 0"
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Todolist-Create  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	jobManagement.InitiateServiceJobManagements(log).Edit(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Create Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}