package jobmanagements

import (
	// "ottosfa-api-web/database/dbmodels"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceJobManagements ..
type ServiceJobManagements struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceJobManagementsInterface ..
type ServiceJobManagementsInterface interface {
	Create(string, models.ReqCreateJobManagement, *models.Response)
	Draft(string, models.ReqFilterJobManagementDraft, *models.ResponsePagination)
	Edit(string, models.ReqEditJobManagement, *models.Response)
	Delete(string, *models.Response)
	Detail(string, *models.Response)
	Upload(string, []byte, *models.Response)
	FilterJobManagements(string , models.ReqFilterJobManagements, *models.ResponsePagination)
	BulkUploadJob(fileBytes []byte, id uint, token string)
	CheckAdmin(string, *models.Response)
	RecipientList(string, models.RecipientReq, *models.Response)
}

// InitiateServiceJobManagements ..
func InitiateServiceJobManagements(log ottologger.OttologInterface) ServiceJobManagementsInterface {
	return &ServiceJobManagements{
		OttoLog: log,
	}
}
