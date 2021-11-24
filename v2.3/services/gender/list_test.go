package gender

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_Listt(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	go InitiateServiceGender(ottolog).List(res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
