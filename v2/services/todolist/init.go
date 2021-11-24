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
	Create(string, models.CreateTodolist, *models.Response)
	NewMerchantDetail(string, models.NewMerchantDetailReq, *models.Response)
	NewMerchantList(string, models.NewMerchantListReq, *models.Response)
	Detail(string, string, *models.Response)
	Upload(string, []byte, *models.Response)
	DownloadTemplate(string, *models.Response)
	Update(string, models.UpdateTodolist, *models.Response)
}

// InitiateServiceTodolist ..
func InitiateServiceTodolist(log ottologger.OttologInterface) ServiceTodolistInterface {
	return &ServiceTodolist{
		OttoLog: log,
	}
}
