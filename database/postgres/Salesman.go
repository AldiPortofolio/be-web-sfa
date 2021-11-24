package postgres

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
)

// GetSalesDetail ..
func (database *DbPostgres) GetSalesDetail(salesID int) (models.SalesmanResponse, error) {
	fmt.Println(">>> Attendance - GetSalesDetail - Postgres <<<")

	var sales dbmodels.Salesman
	var salesmanSer models.SalesmanResponse

	salesErr := Dbcon.Where("id = ?", salesID).Preload("SalesTypes").Preload("Positions").Find(&sales).Error

	if salesErr != nil {
		logs.Error(fmt.Sprintf("Failed to get sales : %d", salesErr))
		return salesmanSer, salesErr
	}

	salesmanSer = SerializedSales(sales)
	salesmanSer.SalesTypes = sales.SalesTypes
	var salesLevel dbmodels.SalesLevelList
	Dbcon.Raw("select id, name from sales_levels where id = ?", sales.SalesLevelID).Find(&salesLevel)
	salesmanSer.SalesLevel = salesLevel.Name

	return salesmanSer, nil
}

// GetParameterConfiguration ..
func (database *DbPostgres) GetParameterConfiguration(param string) (models.ParameterConfiguration, error) {
	fmt.Println(">>> Attendance - GetParameterConfiguration - Postgres <<<")
	var paramReg models.ParameterConfiguration

	err := Dbcon.Raw("select id, name, param_value from parameter_configurations where name = ? ", param).Scan(&paramReg).Error

	if err != nil {
		log.Println("Failed to show ParameterConfiguration: ", err)
		return paramReg, err
	}

	return paramReg, nil
}

