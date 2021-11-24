package merchantNewRecruitment

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"testing"
)

// TestService_Create ..
func TestService_Create(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := dbmodels.MerchantNewRecruitments{
		DatabaseModel:      dbmodels.DatabaseModel{},
		Name:               "",
		PhoneNumber:        "",
		CustomerCode:       "",
		InstitutionCode:    "",
		SubAreaChannelID:   0,
		SubAreaChannelName: "",
		OwnerName:          "",
		Address:            "",
		Longitude:          "",
		Latitude:           "",
		ProvinceId:         0,
		CityId:             0,
		DistrictId:         0,
		VillageId:          0,
		Status:             "",
	}
	go InitiateServiceMerchantNewRecruitment(ottolog).Create("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
