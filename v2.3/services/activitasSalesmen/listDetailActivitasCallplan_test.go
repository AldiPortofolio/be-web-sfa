package activitasSalesmen

import (
	"encoding/json"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"testing"

	ottologger "ottodigital.id/library/logger/v2"
)

func TestService_ListCallplan(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.ResponsePagination
	req := dbmodels.DetailListActivitasTodolistReq{
		Keyword:    "Warkop Bang Udin",
		Date:       "2020-11-20",
		ID:         "9957",
		CategoryID: 4,
		Status:     "Not Exist",
		Page:       1,
	}
	go InitiateServicActivitasSalesmen(ottolog).ListDetailActivitasSalesmenTodolist(req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
