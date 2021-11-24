package merchantNewRecruitment

import (
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
	Upload(string, []byte, *models.Response)
}

// InitiateServiceMerchantNewRecruitment ..
func InitiateServiceMerchantNewRecruitment(log ottologger.OttologInterface) ServiceMerchantNewRecruitmentInterface {
	return &ServiceMerchantNewRecruitment{
		OttoLog: log,
	}
}
