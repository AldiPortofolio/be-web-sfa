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

// Detail godoc
// @Summary Attendance Detail
// @Description Get Detail of Attendances
// @Tags Attendance
// @ID attendances-detail
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /ottosfa/v2.3/attendances/detail/:attendance_id [get]
func Detail(ctx *gin.Context) {
	fmt.Println(">>> Attendances - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	attendID := ctx.Params.ByName("attendance_id")

	log.Info("Attendances-Detail Controller",
		log.AddField("RequestBody:", "attendance_id:"+ attendID))

	attendance.InitiateServiceAttendances(log).Detail(token,attendID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Attendances-Detail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
