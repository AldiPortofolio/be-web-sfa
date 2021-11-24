package province

import (
	"fmt"
	"ottosfa-api-web/models"
)

func (svc *ServiceProvince) ListByCountry(countryID string, res *models.Response) {
	fmt.Println(">>> Detail - ServiceProvince <<<")

	data, err := svc.Database.ProvinceListByCountry(countryID)
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

	return
}
