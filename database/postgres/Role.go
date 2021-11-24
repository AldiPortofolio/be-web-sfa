package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// RoleList ..
func (database *DbPostgres) RoleList() (res []dbmodels.Role, err error) {
	fmt.Println(">>> Role - List - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list role")
		return res, err
	}
	return res, nil
}
