package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// CountryList ..
func (database *DbPostgres) CountryList() (res []dbmodels.Countries, err error) {
	fmt.Println(">>> Country - Country - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list country")
		return res, err
	}
	return res, nil
}
