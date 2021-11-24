package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"strings"

	"github.com/astaxie/beego/logs"
)

// MinMaxPhone ..
func (database *DbPostgres) MinMaxPhone() ([]dbmodels.ParameterConfiguration, error) {
	fmt.Println(">>> ParamConfiguration - MinMaxPhone - Postgres <<<")

	var paramReg []dbmodels.ParameterConfiguration
	names := strings.Split("min_char_merchant_phone max_char_merchant_phone", " ")
	err := Dbcon.Where("name IN (?) ", names).Find(&paramReg).Error

	if err != nil {
		logs.Error("Failed to get ParameterConfiguration: ", err)
	}
	return paramReg, nil
}
