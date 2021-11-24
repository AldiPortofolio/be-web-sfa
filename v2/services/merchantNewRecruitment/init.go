package merchantNewRecruitment

import (
	"ottosfa-api-web/database/dbmodels"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger/v2"
)

// ServiceMerchantNewRecruitment ..
type ServiceMerchantNewRecruitment struct {
	General  models.GeneralModel
	OttoLog  ottologger.OttologInterface
	Database db.DbPostgres
}

// ServiceMerchantNewRecruitmentInterface ..
type ServiceMerchantNewRecruitmentInterface interface {
	List(string, models.MerchantNewRecruitmentListReq, *models.ResponsePagination)
	Export(string, models.MerchantNewRecruitmentListReq, *models.Response)
	Create(string, dbmodels.MerchantNewRecruitments, *models.Response)
	Detail(string, string, *models.Response)
	Upload(string, []byte, *models.Response)
	DownloadTemplate(string, *models.Response)
}

// InitiateServiceMerchantNewRecruitment ..
func InitiateServiceMerchantNewRecruitment(log ottologger.OttologInterface) ServiceMerchantNewRecruitmentInterface {
	return &ServiceMerchantNewRecruitment{
		OttoLog: log,
	}
}
