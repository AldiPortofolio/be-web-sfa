package village

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceVillage ..
type ServiceVillage struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceVillageInterface ..
type ServiceVillageInterface interface {
	ListByDistrict(string, *models.Response)
}

// InitiateServiceVillage ..
func InitiateServiceVillage(log ottologger.OttologInterface) ServiceVillageInterface {
	return &ServiceVillage{
		OttoLog: log,
	}
}
