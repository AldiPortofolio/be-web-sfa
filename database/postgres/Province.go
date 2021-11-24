package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"

	ottologger "ottodigital.id/library/logger"
	ottoutils "ottodigital.id/library/utils"
)

// Province ..
func (database *DbPostgres) Province(search models.SearchReq) (res []dbmodels.Provinces, err error) {
	fmt.Println(">>> Province - Province - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	limit := ottoutils.GetEnv("DB_LIMIT_DATA", " LIMIT 20 ")
	var query = "SELECT * FROM provinces "
	if search.Keyword != "" {
		query = query + " WHERE name ILIKE '%" + search.Keyword + "%' "
	}
	query = query + " ORDER BY updated_at desc " + limit
	err = Dbcon.Raw(query).Scan(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list province")
		return res, err
	}
	return res, nil
}

// Province ..
func (database *DbPostgres) ProvinceListByCountry(countryID string) (res []dbmodels.ProvincesByCountry, err error) {
	fmt.Println(">>> Province - Province - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Order("name ASC").Find(&res).Error
	if err != nil {
		sugarLogger.Info("Failed connect to database SFA when get list province")
		return res, err
	}
	return res, nil
}
