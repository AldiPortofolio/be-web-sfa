package role

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceRole ..
type ServiceRole struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceRoleInterface ..
type ServiceRoleInterface interface {
	List(string, *models.Response)
}

// InitiateServiceRole ..
func InitiateServiceRole(log ottologger.OttologInterface) ServiceRoleInterface {
	return &ServiceRole{
		OttoLog: log,
	}
}
