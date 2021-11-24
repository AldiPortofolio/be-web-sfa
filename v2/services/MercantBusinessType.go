package services

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	db "ottosfa-api-web/database/postgres"
)

// MerchantBusinessType ..
func (svc *Service) MerchantBusinessType(res *models.Response) {
	fmt.Println(">>> MerchantBusinessType - Service <<<")

	database := db.DbPostgres{
		General: svc.General,
	}

	dataDB, err := database.MerchantBusinessTypeList()
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		return
	}

	data := []models.DataModule{}
	for _, val := range dataDB {
		a := models.DataModule{
			Code: val.Code,
			Name: val.Name,
		}
		data = append(data, a)
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)
	res.Data = data
	return
}
