package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

// City ..
func (database *DbPostgres) City(provinceID string, search models.SearchReq) (res []dbmodels.Cities, err error) {
	fmt.Println(">>> City - City - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")
	var query = "SELECT * FROM cities WHERE province_id = " + provinceID
	if search.Keyword != "" {
		query = query + " AND name ILIKE '%" + search.Keyword + "%' "
	}
	query = query + " ORDER BY updated_at desc " + limit
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list city")
		return res, err
	}
	return res, nil
}
