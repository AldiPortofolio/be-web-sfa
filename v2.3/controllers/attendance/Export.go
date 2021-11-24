package attendances

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/attendance"

	"ottosfa-api-web/constants"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Validate godoc
// @Summary Attendance Validate
// @Description Validate of Attendances
// @Tags Attendance
// @ID attendances-export
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param body body models.AttendanceReq true "request body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /ottosfa/v2.3/attendances/export [post]
func Export(ctx *gin.Context) {
	fmt.Println(">>> Attendances - Export - Controller <<<")

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
	log.Info("Attendances-Export Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	if req.Page == 0 {
		req.Page = 1
	}

	attendance.InitiateServiceAttendances(log).Export(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Attendances-Export Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
