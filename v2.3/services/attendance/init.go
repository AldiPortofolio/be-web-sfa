package attendance

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"
	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceAttendance ..
type ServiceAttendance struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceAttendanceInterface ..
type ServiceAttendanceInterface interface {
	List(string, models.AttendanceReq, *models.Response)
	Detail(string, string, *models.Response)
	Validate(string, models.ValidateAttendanceReq, *models.Response)
	Export(string, models.AttendanceReq, *models.Response)
}

// InitiateServiceAttendances ..
func InitiateServiceAttendances(log ottologger.OttologInterface) ServiceAttendanceInterface {
	return &ServiceAttendance{
		OttoLog: log,
	}
}