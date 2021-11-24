package todolist

import (
	"encoding/json"
	"log"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/models"
	"testing"
)

func TestService_Create(t *testing.T) {
	var ottolog ottologger.OttologInterface
	var res *models.Response
	req := models.CreateTodolist{
		Mid:                "",
		TaskDate:           "",
		SalesPhone:         "",
		TodolistCategoryID: 0,
		VillageID:          0,
		Notes:              "",
		CustomerCode:       "",
		MerchantPhone:      "",
		MerchantName:       "",
		TasksAttributes:    nil,
	}
	go InitiateServiceTodolist(ottolog).Create("HVzwlFFaSfdfuJhmOEqFNLxkoqZQlRSX", req, res)

	byteRes, _ := json.Marshal(res)
	log.Println("res --> ", string(byteRes))
}
