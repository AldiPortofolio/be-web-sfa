package postgres

import (
	"errors"
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"strconv"

	ottologger "ottodigital.id/library/logger"
)

// SaveJobCategories ..
func (database *DbPostgres) SaveJobCategories(req dbmodels.JobCategories) (err error) {
	fmt.Println(">>> JobCategories - SaveJobCategories - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Save(&req).Error
	if err != nil {
		sugarLogger.Info("Failed Save JobCategories")
		return err
	}
	return nil
}

// FindByID ...
func (database *DbPostgres) FindByID(id int64) (dbmodels.JobCategories, error) {
	fmt.Println(">>> JobCategories - FindJobCategoriesById - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var data dbmodels.JobCategories

	err := Dbcon.Where(dbmodels.JobCategories{ID: id}).First(&data).Error
	if err != nil {
		sugarLogger.Info("Data not Found")
		return data, err
	}
	return data, err
}

// FindByName ...
func (database *DbPostgres) FindByName(name string) (dbmodels.JobCategories, error) {
	fmt.Println(">>> JobCategories - FindJobCategoriesByName - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var data dbmodels.JobCategories

	err := Dbcon.Where("name = ?", name).First(&data).Error
	if err != nil {
		sugarLogger.Info("Data not Found")
		return data, err
	}
	return data, err
}

// FilterJobCategories ..
func (database *DbPostgres) FilterJobCategories(req models.ReqFilterJobCategories) (res []dbmodels.JobCategories, total int, err error) {
	fmt.Println(">>> JobCategories - FilterJobCategories - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	db := GetDbCon()

	page := req.Page
	limit := req.Limit

	var query string
	query = "1=1"
	if req.ID != 0 {
		query += " and id = " + strconv.Itoa(int(req.ID))
	}

	if req.Name != "" {
		query += " and name ilike '%" + req.Name + "%'"
	}

	err = db.Where(query).Limit(limit).Offset((page - 1) * limit).Order("id ASC").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		sugarLogger.Info("Failed Get List JobCategories")
		return res, 0, err
	}

	return res, total, nil
}

// DeleteJobCategory ..
func (database *DbPostgres) DeleteJobCategory(id int64) error {
	sugarLogger := ottologger.GetLogger()
	var jobcategory dbmodels.JobCategories
	// id, _ := strconv.Atoi(adminID)
	fmt.Println("id", id)
	sErr := Dbcon.Where("id = ?", id).Delete(&jobcategory).Error
	if sErr != nil {
		sugarLogger.Info(sErr.Error())
		return errors.New("Failed delete admin sub area")
	}

	return nil
}

// DetailJobCategory ..
func (database *DbPostgres) DetailJobCategory(ID string) (res dbmodels.JobCategories, err error) {
	fmt.Println(">>> JobCategories - DetailJobCategory - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	db := GetDbCon()

	err = db.Where("id = ?",ID).First(&res).Error

	if err != nil {
		sugarLogger.Info("Failed Get Detail JobCategories")
		return res, err
	}

	return res, nil
}
