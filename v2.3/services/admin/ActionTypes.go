package admin

import (
	"fmt"
	"ottosfa-api-web/models"
)

func (svc *ServicAdmin) ActionTypes(res *models.Response) {
	fmt.Println(">>> ActionType - ServiceAdmin <<<")

	data := []map[string]interface{}{
		{"id": 0, "name": "New"},
		{"id": 1, "name": "Edit"},
		{"id": 2, "name": "Delete"},
		{"id": 3, "name": "Re-Upload"},
	}

	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200
}
