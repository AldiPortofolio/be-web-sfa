package jobManagement

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	jobmanagements "ottosfa-api-web/v2.3/services/jobManagements"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Recipient List ..
// JobManagement - Recipient List godoc
// @Summary JobManagement - Recipient List
// @Description JobManagement - Recipient List
// @ID JobManagement - Recipient List
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/recipient-list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.RecipientReq true "Body"
// @Success 200 {object} models.Response{} "JobManagement - Recipient List EXAMPLE"
func RecipientList(ctx *gin.Context) {
	fmt.Println(">>> JobManabagemnt - Delete - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.RecipientReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	jobmanagements.InitiateServiceJobManagements(log).RecipientList(token,req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("jobcategories-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
