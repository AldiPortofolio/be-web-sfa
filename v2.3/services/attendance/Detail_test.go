package attendance

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_Detail(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	attendId := "1"
	go InitiateServiceAttendances(ottolog).Detail("zdlRDAhyJCumztWHlSUZsHCaPnRSNMou", attendId, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}

