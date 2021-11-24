package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// CompanyCodeList ..
func (database *DbPostgres) CompanyCodeList() (res []dbmodels.CompanyCode, err error) {
	fmt.Println(">>> Company - List - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Order("id ASC").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list company code")
		return res, err
	}
	return res, nil
}
