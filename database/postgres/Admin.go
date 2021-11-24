package postgres

import (
	"errors"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"time"

	ottologger "ottodigital.id/library/logger"
)

// CheckAdminToken ..
func (database *DbPostgres) CheckAdminToken(token string) (uint, error) {
	sugarLogger := ottologger.GetLogger()
	var access dbmodels.AccessTokens
	sql := Dbcon.Raw("select * from access_tokens where token = ?", token).Scan(&access)

	if sql.Error != nil {
		sugarLogger.Info("Failed connect to  get account token: Invalid Token")
		return 0, errors.New("Invalid Token")
	}

	timeNowIn := time.Now().UnixNano() / int64(time.Millisecond)
	validFor := access.ValidFor + access.CreatedAt.UnixNano()
	if validFor < timeNowIn {
		sugarLogger.Info("Failed connect to  get account token: Token Expired")
		return 0, errors.New("Token Expired")
	}

	//sugarLogger.Info(fmt.Sprintf("Success get Admin token: %v", access.Token))

	return access.AdminID, nil
}

// CheckAdminToken ..
func (database *DbPostgres) FindAdminByEmail(email string) (res dbmodels.Admin, err error) {
	sugarLogger := ottologger.GetLogger()
	
	sql := Dbcon.Where("email = ?", email).Find(&res)

	if sql.Error != nil {
		sugarLogger.Info("Failed get User By Email")
		return res, errors.New("email tidak terdaftar")
	}

	return res, nil
}

// CheckAdminToken ..
func (database *DbPostgres) FindAdminById(id int64) (res dbmodels.Admin, err error) {
	sugarLogger := ottologger.GetLogger()
	
	sql := Dbcon.Where("id = ?", id).Find(&res)

	if sql.Error != nil {
		sugarLogger.Info("Failed get User By Id")
		return res, errors.New("data not found")
	}

	return res, nil
}

// CheckAdminToken ..
func (database *DbPostgres) FindAdminByAssignmentRole(roles []string, req models.RecipientReq) (res []dbmodels.Admin, err error) {
	sugarLogger := ottologger.GetLogger()

	where := ""
	if req.Keyword != "" {
		where = " coalesce(first_name , '') || coalesce(' ' , '') || coalesce(last_name , '') ilike '%" + req.Keyword + "%' and "
	}
	
	sql := Dbcon.Where(where + " assignment_role in (?)", roles).Find(&res)

	if sql.Error != nil {
		sugarLogger.Info("Failed get User By assigmenr role")
		return res, errors.New("data not found")
	}

	return res, nil
}