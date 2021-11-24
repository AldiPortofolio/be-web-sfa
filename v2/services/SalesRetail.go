package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	db "ottosfa-api-web/database/postgres"
)

// SalesRetail ..
func (svc *Service) SalesRetail(res *models.Response) {
	fmt.Println(">>> SalesRetail - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	data, err := database.SalesTypeList()
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
