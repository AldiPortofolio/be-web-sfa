package attendance

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
)

// Detail ..
func (svc *ServiceAttendance) Detail(token string, attendid string, res *models.Response){
	var resp models.Response
	fmt.Println(">>> Detail - ServiceAttendance <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		resp.Data = gin.H{}
		return
	}

	attendID, _ := strconv.Atoi(attendid)
	attendDetail, err :=  svc.Database.GetAttendanceDetail(attendID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.detail.failed")
		return
	}

	salesDetail, err := svc.Database.GetSalesDetail(attendDetail.SalesID)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.detail.failed")
		return
	}

	paramRecog, err := svc.Database.GetParameterConfiguration("face_recog_atd_percentage")
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.detail.failed")
		return
	}

	//attendDetail.PhotoProfile = salesDetail.Photo
	//attendDetail.MinAccPercentage = paramRecog.ParamValue

	region := []string{}
	subAreaChannel := []string{}
	for _, val := range salesDetail.Positions {
		region = append(region, val.Region)
		subAreaChannel = append(subAreaChannel, val.SubArea)
	}

	//salesType := []string{}
	//for _, val := range salesDetail.SalesTypes {
	//	salesType = append(salesType, val.Name)
	//}

	dataRes := models.AttendanceDetailRes{
		ID:                 attendDetail.ID,
		SalesID:            attendDetail.SalesID,
		SalesPhone:         attendDetail.SalesPhone,
		SalesName:          attendDetail.SalesName,
		Selfie:             attendDetail.Selfie,
		ClocktimeServer:    attendDetail.ClocktimeServer,
		ClocktimeLocal:     attendDetail.ClocktimeLocal,
		Location:           attendDetail.Location,
		Latitude:           attendDetail.Latitude,
		Longitude:          attendDetail.Longitude,
		AttendCategory:     attendDetail.AttendCategory,
		AttendCategoryType: attendDetail.AttendCategoryType,
		TypeAttendance:     attendDetail.TypeAttendance,
		Notes:              attendDetail.Notes,
		PhotoAccuration:    attendDetail.PhotoAccuration,
		PhotoProfile:       salesDetail.Photo,
		MinAccPercentage:   paramRecog.ParamValue,
		Status:             attendDetail.Status,
		Region: 			region,
		SubAreaChanel:      subAreaChannel,
		SalesType: 			salesDetail.SalesType,
		UpdatedAt:			attendDetail.UpdatedAt,
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = dataRes

	return
}
