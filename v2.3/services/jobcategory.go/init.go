package jobcategory

import (
	"ottosfa-api-web/database/dbmodels"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceJobCategories ..
type ServiceJobCategories struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceJobCategoriesInterface ..
type ServiceJobCategoriesInterface interface {
	Create(string, dbmodels.JobCategories, *models.Response)
	Update(string, dbmodels.JobCategories, *models.Response)
	FilterJobCategories(models.ReqFilterJobCategories, *models.ResponsePagination)
	Delete(string, int64, *models.Response)
	Detail(string, string, *models.Response)
}

// InitiateServiceJobCategories ..
func InitiateServiceJobCategories(log ottologger.OttologInterface) ServiceJobCategoriesInterface {
	return &ServiceJobCategories{
		OttoLog: log,
	}
}
