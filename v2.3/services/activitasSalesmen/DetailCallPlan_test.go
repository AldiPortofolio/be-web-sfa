package activitasSalesmen

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_DetailCallplan(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := "1"
	go InitiateServicActivitasSalesmen(ottolog).DetailCallPlan(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
