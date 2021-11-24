package activitasSalesmen

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
)

func (svc *ServicActivitasSalesmen) DetailActivitasSalesmen(req dbmodels.DetailActivitasSalesmenReq, res *models.Response) {
	fmt.Println(">>> DetailActivitasSalesmen - ServicActivitasSalesmen <<<")

	data, err := svc.Database.DetailActivitasSalesmen(req)
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
