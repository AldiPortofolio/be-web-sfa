package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Village ..
func (svc *Service) Village(districtId string, req models.SearchReq, res *models.Response) {
	fmt.Println(">>> Village - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	data, err := database.Village(districtId, req)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
