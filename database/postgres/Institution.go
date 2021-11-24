package postgres

import (
	"fmt"
	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
)

// Institution ..
func (database *DbPostgres) Institution(search models.SearchReq) (res []dbmodels.Institutions, err error) {
	fmt.Println(">>> Institutions - Institution - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")

	var query = "SELECT * FROM institutions "

	if search.Keyword != "" {
		query = query + "WHERE code ILIKE '%" + search.Keyword + "%'"
	}

	query = query + "ORDER BY updated_at desc " + limit

	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list institutions")
		return res, err
	}
	return res, nil
}
