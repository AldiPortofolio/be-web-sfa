package activitasSalesmen

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_List(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.ResponsePagination
	token := "sXDMaHKmlbvphFbBDLLjuTNnvKdPHifC"
	req := dbmodels.ListActivitasSalesmenReq{
		Keyword:     "testing fai",
		PeriodFrom:  "2021-02-05",
		PeriodTo:    "2021-02-10",
		SalesTypeID: "",
		Page:        1,
	}
	go InitiateServicActivitasSalesmen(ottolog).ListActivitasSalesmen(token, req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
