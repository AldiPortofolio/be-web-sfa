package todolist

import (
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceTodolist ..
type ServiceTodolist struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceTodolistInterface ..
type ServiceTodolistInterface interface {
	Create(string, models.CreateTodolistV2, *models.Response)
	Update(string, models.UpdateTodolistV2, *models.Response)
	MerchantDetail(models.MerchantDetailReq, *models.Response)
	MerchantList(models.MerchantListReq, *models.Response)
	Upload(string, []byte, *models.Response)
	Detail(string, string, *models.Response)
}

// InitiateServiceTodolist ..
func InitiateServiceTodolist(log ottologger.OttologInterface) ServiceTodolistInterface {
	return &ServiceTodolist{
		OttoLog: log,
	}
}
