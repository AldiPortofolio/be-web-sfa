package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-web/database/dbmodels"
)

// SalesTypeList ..
func (database *DbPostgres) SalesTypeList() ([]dbmodels.SalesType, error) {
	fmt.Println(">>> SalesTypeList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.SalesType{}

	err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get sales type")
		return res, err
	}
	return res, err
}
