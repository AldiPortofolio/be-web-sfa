package village

import (
	"fmt"
	"ottosfa-api-web/models"
)

func (svc *ServiceVillage) ListByDistrict(districtID string, res *models.Response) {
	fmt.Println(">>> List By District - ServiceVillage <<<")

	data, err := svc.Database.VillageListByDistrict(districtID)
	if err != nil {
		res.Meta.Status = false
		res.Meta.Code = 422
		res.Meta.Message = err.Error()
		return
	}

	res.Data = map[string]interface{}{
		"villages": data,
	}
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
