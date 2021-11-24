package activitasSalesmen

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_DetailSales(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := dbmodels.DetailActivitasSalesmenReq{
		ID:   "9957",
		Date: "2020-11-20",
	}
	go InitiateServicActivitasSalesmen(ottolog).DetailActivitasSalesmen(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
