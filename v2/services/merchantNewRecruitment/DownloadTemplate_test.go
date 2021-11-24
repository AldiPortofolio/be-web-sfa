package merchantNewRecruitment

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestServiceMerchantNewRecruitment_DownloadTemplate(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	go InitiateServiceMerchantNewRecruitment(ottolog).DownloadTemplate("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
