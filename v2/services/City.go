package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// City ..
func (svc *Service) City(provinceId string, req models.SearchReq, res *models.Response) {
	fmt.Println(">>> City - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	data, err := database.City(provinceId, req)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
