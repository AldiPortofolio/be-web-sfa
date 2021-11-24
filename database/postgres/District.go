package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

// District ..
func (database *DbPostgres) District(cityID string, search models.SearchReq) (res []dbmodels.District, err error) {
	fmt.Println(">>> District - District - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")
	var query = "SELECT * FROM districts WHERE city_id = " + cityID
	if search.Keyword != "" {
		query = query + " AND name ILIKE '%" + search.Keyword + "%' "
	}
	query = query + " ORDER BY updated_at desc " + limit
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list district")
		return res, err
	}
	return res, nil
}
