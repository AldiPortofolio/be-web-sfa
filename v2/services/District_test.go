package services

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestService_District(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.SearchReq{
		Keyword: "",
	}
	go InitiateService(ottolog).District("3100", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
