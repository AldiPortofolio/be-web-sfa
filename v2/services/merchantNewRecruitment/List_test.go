package merchantNewRecruitment

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestService_List(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.ResponsePagination
	req := models.MerchantNewRecruitmentListReq{
		Id:                 "",
		Name:               "",
		CustomerCode:       "",
		PhoneNumber:        "",
		InstitutionCode:    "",
		SubAreaChannelName: "",
		Status:             "",
		Page:               0,
	}
	go InitiateServiceMerchantNewRecruitment(ottolog).List("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
