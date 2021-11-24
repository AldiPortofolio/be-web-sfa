package merchantNewRecruitment

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestService_Upload(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := []byte{}
	go InitiateServiceMerchantNewRecruitment(ottolog).Upload("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
