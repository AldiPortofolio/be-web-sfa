package attendance

import (
	"fmt"
	"math"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
)

// List ..
func (svc *ServiceAttendance) List(token string, params models.AttendanceReq, res *models.Response){
	var resp models.Response
	fmt.Println(">>> List - ServiceAcquisition <<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		resp.Data = gin.H{}
		return
	}

	coverageArea, err := svc.Database.GetSessionCoverageArea(adminID)
	
	var attendanceListDb []dbmodels.Attendance
	var total int64
	var errs error
	if coverageArea != "" {
		attendanceListDb, total, errs = svc.Database.GetFilterAttendanceByCoverageArea(params, coverageArea)
	}else{
		attendanceListDb, total, errs = svc.Database.GetFilterAttendance(params)
	}

	if errs != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		resp.Data = gin.H{}
		return
	}

	metaPaging := models.MetaPagination{
		CurrentPage: int64(params.Page),
		NextPage:    int64(params.Page + 1),
		PrevPage:    int64(params.Page - 1),
		TotalPages:  int64(math.Ceil(float64(total) / 25)),
		TotalCount:  total,
	}

	attendances := map[string]interface{}{}
	if attendanceListDb == nil {
		attendances["attendances"] = []string{}
	} else {
		attendancesRes := []models.AttendanceRes{}
		for _, val := range attendanceListDb {
			a := models.AttendanceRes{
				ID:                 val.ID,
				SalesID:            val.SalesID,
				SalesPhone:         val.SalesPhone,
				SalesName:          val.SalesName,
				ClocktimeServer:    val.ClocktimeServer,
				AttendCategory:     val.AttendCategory,
				AttendCategoryType: val.AttendCategoryType,
				Notes:              val.Notes,
				Status: 			val.Status,
				StatusName:         utils.StatusAttendance(val.Status),
				TypeAttendance:     val.TypeAttendance,
			}
			attendancesRes = append(attendancesRes, a)
		}
		attendances["attendances"] = attendancesRes
	}

	attendances["meta"] = metaPaging

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = attendances

	return
}
