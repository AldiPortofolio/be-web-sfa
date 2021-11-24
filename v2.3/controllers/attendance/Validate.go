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

// Validate godoc
// @Summary Attendance Validate
// @Description Validate of Attendances
// @Tags Attendance
// @ID attendances-validate
// @Accept  json
// @Produce  json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param body body models.ValidateAttendanceReq true "request body"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /ottosfa/v2.3/attendances/validate [post]
func Validate(ctx *gin.Context) {
	fmt.Println(">>> Attendances - Validate - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ValidateAttendanceReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Attendances-Validate Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	attendance.InitiateServiceAttendances(log).Validate(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Attendances-Validate Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
