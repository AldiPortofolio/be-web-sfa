package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Institution ..
func (svc *Service) Institution(req models.SearchReq, res *models.Response) {
	fmt.Println(">>> Institution - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	data, err := database.Institution(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
