package admin

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServicAdmin ..
type ServicAdmin struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServicAdminInterface ..
type ServicAdminInterface interface {
	ActionTypes(*models.Response)
}

// InitiateServicAdmin ..
func InitiateServicAdmin(log ottologger.OttologInterface) ServicAdminInterface {
	return &ServicAdmin{
		OttoLog: log,
	}
}
