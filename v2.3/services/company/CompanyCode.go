package company

import (
	"fmt"
	"ottosfa-api-web/models"
	"strconv"
)

func (svc *ServiceCompany) CompanyCodes(res *models.Response) {
	fmt.Println(">>> List - ServiceCompany <<<")

	data, err := svc.Database.CompanyCodeList()
	if err != nil {
		res.Meta.Status = false
		res.Meta.Code = 422
		res.Meta.Message = err.Error()
		return
	}

	companyCodes := []string{}

	for _, v := range data {
		code := strconv.Itoa(v.Code) + "-" + v.Name
		companyCodes = append(companyCodes, code)
	}

	res.Data = map[string]interface{}{
		"company_codes": companyCodes,
	}
	res.Meta.Message = "success"
	res.Meta.Status = true
	res.Meta.Code = 200

	return
}
