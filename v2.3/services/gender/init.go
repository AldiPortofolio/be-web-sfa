package gender

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceGender ..
type ServiceGender struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceGenderInterface ..
type ServiceGenderInterface interface {
	List(*models.Response)
}

// InitiateServiceGender ..
func InitiateServiceGender(log ottologger.OttologInterface) ServiceGenderInterface {
	return &ServiceGender{
		OttoLog: log,
	}
}
