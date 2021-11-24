package postgres

import (
	"errors"
	"fmt"
	"ottosfa-api-web/database/dbmodels"

	ottologger "ottodigital.id/library/logger"
)

// DeleteByAdminID ..
func (database *DbPostgres) DeleteByAdmin(adminID string) error {
	sugarLogger := ottologger.GetLogger()
	var subarea dbmodels.AdminSubArea
	// id, _ := strconv.Atoi(adminID)
	fmt.Println("adminID", adminID)
	sErr := Dbcon.Where("admin_id = ?", adminID).Delete(&subarea).Error
	if sErr != nil {
		sugarLogger.Info(sErr.Error())
		return errors.New("Failed delete admin sub area")
	}

	return nil
}
