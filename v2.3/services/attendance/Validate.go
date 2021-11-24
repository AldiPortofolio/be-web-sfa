package attendance

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Validate ..
func (svc *ServiceAttendance) Validate(token string, req models.ValidateAttendanceReq, res *models.Response){
	var resp models.Response
	fmt.Println(">>> Validate - ServiceAttendance <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		resp.Data = gin.H{}
		return
	}

	if req.StatusBefore != 1 {
		res.Meta = utils.GetMetaResponse("attendance.validate.failed")
		return
	}

	err = svc.Database.UpdateStatusAttendance(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse("attendance.validate.failed")
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

	return
}
