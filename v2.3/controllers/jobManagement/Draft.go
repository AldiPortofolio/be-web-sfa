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
// @Summary JobManagement - Draft
// @Description JobManagement - Draft
// @ID JobManagement - Draft
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/draft [POST]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.ReqFilterJobManagementDraft true "Body"
// @Success 200 {object} models.ResponsePagination{}
func Draft(ctx *gin.Context) {
	fmt.Println(">>> JobManagement - Draft - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.ResponsePagination{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	token := utils.GetToken(ctx.Request)

	req := models.ReqFilterJobManagementDraft{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Page == 0 {
		req.Page = 1
	}

	jobManagement.InitiateServiceJobManagements(log).Draft(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("JobManagement-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
