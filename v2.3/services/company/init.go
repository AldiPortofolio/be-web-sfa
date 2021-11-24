package company

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceCompany ..
type ServiceCompany struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceCompanyInterface ..
type ServiceCompanyInterface interface {
	List(string, *models.Response)
	CompanyCodes(*models.Response)
}

// InitiateServiceCompany ..
func InitiateServiceCompany(log ottologger.OttologInterface) ServiceCompanyInterface {
	return &ServiceCompany{
		OttoLog: log,
	}
}
