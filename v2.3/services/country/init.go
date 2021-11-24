package country

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceCountry ..
type ServiceCountry struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceCountryInterface ..
type ServiceCountryInterface interface {
	List(*models.Response)
}

// InitiateServiceCountry ..
func InitiateServiceCountry(log ottologger.OttologInterface) ServiceCountryInterface {
	return &ServiceCountry{
		OttoLog: log,
	}
}
