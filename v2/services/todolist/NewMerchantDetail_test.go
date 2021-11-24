package todolist

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestServiceTodolist_NewMerchantDetail(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.NewMerchantDetailReq{
		CustomerCode: "",
		PhoneNumber:  "",
	}
	go InitiateServiceTodolist(ottolog).NewMerchantDetail("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
