package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-web/database/dbmodels"
)

// MerchantBusinessTypeList ..
func (database *DbPostgres) MerchantBusinessTypeList() ([]dbmodels.MerchantBusinessType, error) {
	fmt.Println(">>> MerchantBusinessTypeList - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	res := []dbmodels.MerchantBusinessType{}

	err := Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get business type")
		return res, err
	}
	return res, err
}
