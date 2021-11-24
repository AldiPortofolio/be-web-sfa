package activitasSalesmen

import (
	"ottosfa-api-web/database/dbmodels"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServicActivitasSalesmen ..
type ServicActivitasSalesmen struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServicActivitasSalesmenInterface ..
type ServicActivitasSalesmenInterface interface {
	ListActivitasSalesmen(string, dbmodels.ListActivitasSalesmenReq, *models.ResponsePagination)
	DetailActivitasSalesmen(dbmodels.DetailActivitasSalesmenReq, *models.Response)
	ListDetailActivitasSalesmenTodolist(dbmodels.DetailListActivitasTodolistReq, *models.ResponsePagination)
	ListDetailActivitasSalesmenCallplan(dbmodels.DetailListActivitasTodolistReq, *models.ResponsePagination)
	DetailCallPlan(string, *models.Response)
	DetailTodolist(string, *models.Response)
}

// InitiateServicActivitasSalesmen ..
func InitiateServicActivitasSalesmen(log ottologger.OttologInterface) ServicActivitasSalesmenInterface {
	return &ServicActivitasSalesmen{
		OttoLog: log,
	}
}
