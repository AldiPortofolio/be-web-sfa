package activitasSalesmen

import (
	"fmt"
	"ottosfa-api-web/models"
)

func (svc *ServicActivitasSalesmen) DetailCallPlan(callplanMerchantID string, res *models.Response) {
	fmt.Println(">>> DetailActivitasSalesmen - DetailCallPlan <<<")

	data, err := svc.Database.DetailCallplan(callplanMerchantID)
	if err != nil {
		res.Meta.Status = false
		res.Meta.Code = 422
		res.Meta.Message = err.Error()
		return
	}

	res.Data = data
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200
}
