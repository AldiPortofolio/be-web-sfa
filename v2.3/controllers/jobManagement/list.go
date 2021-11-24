package jobManagement

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	jobManagement "ottosfa-api-web/v2.3/services/jobManagements"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// JobManagement - list godoc
// @Summary JobManagement - Filter
// @Description JobManagement - Filter
// @ID JobManagement - Filter
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/filter [POST]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.ReqFilterJobManagements true "Body"
// @Success 200 {object} models.ResponsePagination{}
func Filter(ctx *gin.Context) {
	fmt.Println(">>> JobManagement - Filter - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)
	token := utils.GetToken(ctx.Request)
	res := models.ResponsePagination{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReqFilterJobManagements{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.ResponseCode = "400"
		res.Message = fmt.Sprintf("Body request error: %v", err)
		ctx.JSON(http.StatusOK, res)
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}

	if req.Limit == 0 {
		req.Limit = 10
	}

	jobManagement.InitiateServiceJobManagements(log).FilterJobManagements(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("JobManagement-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}