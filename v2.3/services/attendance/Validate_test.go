package attendance

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_Validate(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.ValidateAttendanceReq{
		AttendanceId: "1",
		StatusBefore: 1,
		StatusAfter:  1,
		Reason:       "",
	}
	go InitiateServiceAttendances(ottolog).Validate("zdlRDAhyJCumztWHlSUZsHCaPnRSNMou", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
