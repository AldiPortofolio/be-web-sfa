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

// Detail godoc
// @Summary JobManabagemnt - Detail
// @Description JobManabagemnt - Detail
// @ID JobManabagemnt - Detail
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/detail [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "JobManabagemnt - Detail EXAMPLE"
func Detail(ctx *gin.Context) {
	fmt.Println(">>> JobManabagemnt - Delete - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)


	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	id := ctx.Params.ByName("id")


	jobmanagements.InitiateServiceJobManagements(log).Detail(id, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("jobcategories-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
