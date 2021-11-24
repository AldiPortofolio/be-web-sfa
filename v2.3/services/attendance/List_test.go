package attendance

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_List(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.AttendanceReq{
		ID: "",
		SalesPhone: "",
		SalesName: "",
		Category: "",
		Type: "",
		DateFrom: "",
		DateTo: "",
		Notes: "",
		Page: 1,
		Limit: 25,
	}
	go InitiateServiceAttendances(ottolog).List("zdlRDAhyJCumztWHlSUZsHCaPnRSNMou", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}

