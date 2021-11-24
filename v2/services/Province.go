package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	db "ottosfa-api-web/database/postgres"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Province ..
func (svc *Service) Province(req models.SearchReq, res *models.Response) {
	fmt.Println(">>> Province - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	data, err := database.Province(req)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data

	return
}
