package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

// SubArea ..
func (database *DbPostgres) SubArea(villageID string, search models.SearchReq) (res []dbmodels.SubArea, err error) {
	fmt.Println(">>> SubAreaChannel - SubArea - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")
	var query = " SELECT sa.* " +
		" FROM sub_areas sa " +
		" LEFT JOIN sub_areas_villages sav on sav.sub_area_id = sa.id " +
		" WHERE sav.village_id = " + villageID
	if search.Keyword != "" {
		query = query + " AND sa.name ILIKE '%" + search.Keyword + "%' "
	}
	query = query + " ORDER BY updated_at desc " + limit
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list sub area")
		return res, err
	}
	return res, nil
}
