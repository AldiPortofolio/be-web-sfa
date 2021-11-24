package merchantNewRecruitment

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestService_Detail(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	go InitiateServiceMerchantNewRecruitment(ottolog).Detail("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", "10", res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
