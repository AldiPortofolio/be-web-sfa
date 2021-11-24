package todolist

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestServiceTodolist_NewMerchantList(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.NewMerchantListReq{
		Keyword: "",
	}
	go InitiateServiceTodolist(ottolog).NewMerchantList("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
