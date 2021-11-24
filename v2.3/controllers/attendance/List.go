package attendances

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/models"
	"ottosfa-api-web/v2.3/services/attendance"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/constants"
)

// List godoc
// @Summary Attendance List
// @Description Get filter and list of Attendances
// @Tags Attendance
// @ID attendances-list
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param body body models.AttendanceReq true "request body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /ottosfa/v2.3/attendances/list [post]
func List(ctx *gin.Context) {
	fmt.Println(">>> Attendances - List - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.AttendanceReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Attendances-List Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	if req.Page == 0 {
		req.Page = 1
	}

	attendance.InitiateServiceAttendances(log).List(token,req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Attendances-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
