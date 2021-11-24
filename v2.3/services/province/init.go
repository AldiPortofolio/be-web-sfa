package province

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceProvince ..
type ServiceProvince struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceProvinceInterface ..
type ServiceProvinceInterface interface {
	ListByCountry(string, *models.Response)
}

// InitiateServiceProvince ..
func InitiateServiceProvince(log ottologger.OttologInterface) ServiceProvinceInterface {
	return &ServiceProvince{
		OttoLog: log,
	}
}
