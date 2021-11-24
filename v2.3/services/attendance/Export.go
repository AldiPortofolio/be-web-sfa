package attendance

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"io/ioutil"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// Export ..
func (svc *ServiceAttendance) Export(token string, req models.AttendanceReq, res *models.Response){
	var resp models.Response
	fmt.Println(">>> Export - ServiceAttendance <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		resp.Data = gin.H{}
		return
	}

	data, total, err := svc.Database.GetFilterAttendanceExport(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.export.failed")
		return
	}
	fmt.Println("Total --- ", total)

	dataExport, err := svc.Database.ExportAttendance(data)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.export.failed")
		return
	}

	csvContent, err := gocsv.MarshalString(&dataExport) // Get all clients as CSV string
	fmt.Println("Export ------ ", csvContent)
	// err = gocsv.MarshalFile(&res, clientsFile) // Use this to save the CSV back to the file
	if err != nil {
		panic(err)
	}

	r := strings.NewReader(csvContent)
	file, err := ioutil.ReadAll(r)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.export.failed")
		return
	}

	res.Meta = utils.GetMetaResponse("attendance.export.success")
	res.Data = file

	return
}
