package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

// Village ..
func (database *DbPostgres) Village(districtID string, search models.SearchReq) (res []dbmodels.Village, err error) {
	fmt.Println(">>> Village - Village - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")
	var query = "SELECT * FROM villages WHERE district_id = " + districtID
	if search.Keyword != "" {
		query = query + " AND name ILIKE '%" + search.Keyword + "%' "
	}
	query = query + " ORDER BY updated_at desc " + limit
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list village")
		return res, err
	}
	return res, nil
}

// VillageByDistrict ..
func (database *DbPostgres) VillageListByDistrict(districtID string) (res []dbmodels.VillageByDistrict, err error) {
	fmt.Println(">>> Village - VillageListByDistrict - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	query := "select v.id, INITCAP (v.name) AS name from villages v where v.district_id = ? order by v.name"

	err = Dbcon.Raw(query, districtID).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list village by district")
		return res, err
	}
	return res, nil
}
