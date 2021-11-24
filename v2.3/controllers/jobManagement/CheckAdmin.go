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


func CheckAdmin(ctx *gin.Context) {
	fmt.Println(">>> JobManabagemnt - Delete - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	jobmanagements.InitiateServiceJobManagements(log).CheckAdmin(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("jobcategories-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
