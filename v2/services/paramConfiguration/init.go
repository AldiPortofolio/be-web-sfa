package paramConfiguration

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceParamConfiguration ..
type ServiceParamConfiguration struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceParamConfigurationInterface ..
type ServiceParamConfigurationInterface interface {
	MinMaxPhone(string, *models.Response)
}

// InitiateServiceParamConfiguration ..
func InitiateServiceParamConfiguration(log ottologger.OttologInterface) ServiceParamConfigurationInterface {
	return &ServiceParamConfiguration{
		OttoLog: log,
	}
}
