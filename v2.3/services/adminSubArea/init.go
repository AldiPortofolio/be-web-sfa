package adminSubArea

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServicAdminSubArea ..
type ServicAdminSubArea struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServicAdminSubAreaInterface ..
type ServicAdminSubAreaInterface interface {
	DeleteByAdmin(string, *models.Response)
}

// InitiateServicAdminSubArea ..
func InitiateServicAdminSubArea(log ottologger.OttologInterface) ServicAdminSubAreaInterface {
	return &ServicAdminSubArea{
		OttoLog: log,
	}
}
