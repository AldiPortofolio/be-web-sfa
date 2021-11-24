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

// Create ..
// JobManagement - Create godoc
// @Summary JobManagement - Create
// @Description JobManagement - Create
// @ID JobManagement - Create
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/save [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.JobManagements true "Body"
// @Success 200 {object} models.Response{} "JobManagement - Create EXAMPLE"
func Create(ctx *gin.Context) {
	fmt.Println(">>> JobManagement - Create - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReqCreateJobManagement{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	if req.Name == "" || len(req.RecipientId) == 0 || req.JobCategoryId == 0 || req.JobPriority == "" || req.AssignmentDate == "" || req.Deadline == "" {
		res.Meta.Code = 400
		res.Meta.Message = "nama tugas, penerima tugas, tugas kategori, prioritas , tanggal selesai tidak boleh kosong "
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Todolist-Create  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	jobManagement.InitiateServiceJobManagements(log).Create(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Create Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}