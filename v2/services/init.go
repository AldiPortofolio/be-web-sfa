package services

import (
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
)

// Service ..
type Service struct {
	General models.GeneralModel
	OttoLog ottologger.OttologInterface
}

// ServiceInterface ..
type ServiceInterface interface {
	Institution(models.SearchReq, *models.Response)
	Province(models.SearchReq, *models.Response)
	City(string, models.SearchReq, *models.Response)
	District(string, models.SearchReq, *models.Response)
	Village(string, models.SearchReq, *models.Response)
	SubAreaChannel(string, models.SearchReq, *models.Response)

	MerchantType(*models.Response)
	MerchantCategory(*models.Response)
	MerchantGroup(models.MerchantGroupReq, *models.Response)
	MerchantBusinessType(*models.Response)
	SalesRetail(*models.Response)

	HealthCheck(*models.Response)
}

// InitiateService ..
func InitiateService(log ottologger.OttologInterface) ServiceInterface {
	return &Service{
		OttoLog: log,
	}
}
