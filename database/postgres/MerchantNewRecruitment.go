package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"strings"

	"github.com/jinzhu/gorm"
)

// MerchantNewRecruitmentList ..
func (database *DbPostgres) MerchantNewRecruitmentList(params models.MerchantNewRecruitmentListReq) ([]dbmodels.MerchantNewRecruitments, int64, error) {
	fmt.Println(">>> MerchantNewRecruitment - MerchantNewRecruitmentList - Postgres <<<")
	sugarLogger := database.General.OttoZaplog

	var merchantsNewRecruitments []dbmodels.MerchantNewRecruitments
	var total TotalRow

	limit := int64(25)
	page := params.Page

	sqlList := make(chan string)
	sqlCount := make(chan string)

	selectQuery := "SELECT * "

	go GenerateQueryListMerchantNewRecruitment(selectQuery, params, sqlList, " ORDER BY mnr.updated_at desc ")
	go GenerateQueryListMerchantNewRecruitment("SELECT count(*) total ", params, sqlCount, "")

	order := ""
	if limit != 0 {
		order = fmt.Sprintf(" OFFSET %d LIMIT %d", (page-1)*limit, limit)
	}

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &merchantsNewRecruitments, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rows.Error))
		return merchantsNewRecruitments, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rowsCount.Error))
		return merchantsNewRecruitments, 0, rowsCount.Error
	}

	return merchantsNewRecruitments, total.Total, nil
}

// GenerateQueryListMerchantNewRecruitment ..
func GenerateQueryListMerchantNewRecruitment(query string, req models.MerchantNewRecruitmentListReq, strChan chan string, order string) {

	sql := query + " FROM merchant_new_recruitments mnr"

	where := " WHERE "

	if req.Id != "" || req.Name != "" || req.CustomerCode != "" || req.PhoneNumber != "" || req.InstitutionCode != "" || req.SubAreaChannelName != "" || req.Status != "" {
		sql = sql + where
	}

	paramString := []string{}

	if req.Id != "" {
		id := " mnr.id =" + req.Id
		paramString = append(paramString, id)
	}

	if req.Name != "" {
		name := " mnr.name ILIKE '%" + req.Name + "%' "
		paramString = append(paramString, name)
	}

	if req.CustomerCode != "" {
		customerCode := "mnr.customer_code ILIKE '%" + req.CustomerCode + "%' "
		paramString = append(paramString, customerCode)
	}

	if req.PhoneNumber != "" {
		phoneNumber := "mnr.phone_number ILIKE '%" + req.PhoneNumber + "%' "
		paramString = append(paramString, phoneNumber)
	}

	if req.InstitutionCode != "" {
		paramString = append(paramString, fmt.Sprintf(" mnr.institution_code = '%s'", req.InstitutionCode))
	}

	if req.SubAreaChannelName != "" {
		SubAreaChannelName := "mnr.sub_area_channel_name ILIKE '%" + req.SubAreaChannelName + "%' "
		paramString = append(paramString, SubAreaChannelName)
	}

	if req.Status != "" {
		paramString = append(paramString, fmt.Sprintf(" mnr.status = '%s'", req.Status))
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	sql = sql + paramFilter + order

	strChan <- sql
}

// MerchantNewRecruitmentDataExport ..
func (database *DbPostgres) MerchantNewRecruitmentDataExport(req models.MerchantNewRecruitmentListReq) (res []models.MerchantNewRecruitmentExportRes, err error) {
	fmt.Println(">>> MerchantNewRecruitment - MerchantNewRecruitmentDataExport - Postgres <<<")
	sugarLogger := database.General.OttoZaplog

	sql := " SELECT mnr.*, " +
		" p.name province_name, c.name city_name, d.name district_name, v.name village_name " +
		" FROM merchant_new_recruitments mnr "
	join := " LEFT JOIN provinces p on p.id = mnr.province_id " +
		" LEFT JOIN cities c on c.id = mnr.city_id" +
		" LEFT JOIN districts d on d.id = mnr.district_id" +
		" LEFT JOIN villages v on v.id = mnr.village_id"
	where := " WHERE "
	order := " ORDER BY mnr.updated_at desc "

	sql = sql + join
	if req.Id != "" || req.Name != "" || req.CustomerCode != "" || req.PhoneNumber != "" || req.InstitutionCode != "" || req.SubAreaChannelName != "" || req.Status != "" {
		sql = sql + where
	}

	paramString := []string{}

	if req.Id != "" {
		id := " mnr.id =" + req.Id
		paramString = append(paramString, id)
	}

	if req.Name != "" {
		name := " mnr.name ILIKE '%" + req.Name + "%' "
		paramString = append(paramString, name)
	}

	if req.CustomerCode != "" {
		customerCode := "mnr.customer_code ILIKE '%" + req.CustomerCode + "%' "
		paramString = append(paramString, customerCode)
	}

	if req.PhoneNumber != "" {
		paramString = append(paramString, fmt.Sprintf(" mnr.phone_number = '%s'", req.PhoneNumber))
	}

	if req.InstitutionCode != "" {
		paramString = append(paramString, fmt.Sprintf(" mnr.institution_code = '%s'", req.InstitutionCode))
	}

	if req.SubAreaChannelName != "" {
		SubAreaChannelName := "mnr.sub_area_channel_name ILIKE '%" + req.SubAreaChannelName + "%' "
		paramString = append(paramString, SubAreaChannelName)
	}

	if req.Status != "" {
		paramString = append(paramString, fmt.Sprintf(" mnr.status = '%s'", req.Status))
	}

	paramFilter := strings.Join(paramString[:], " AND ")
	sql = sql + paramFilter + order

	err = Dbcon.Raw(sql).Scan(&res).Error
	if err != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database SFA when export:", err))
		return res, err
	}

	return res, nil
}

// MerchantNewRecruitmentCreate ..
func (database *DbPostgres) MerchantNewRecruitmentCreate(merchantNewRecruitment dbmodels.MerchantNewRecruitments) (res dbmodels.MerchantNewRecruitments, err error) {
	fmt.Println(">>> MerchantNewRecruitment - MerchantNewRecruitmentCreate - Postgres <<<")
	sugarLogger := database.General.OttoZaplog

	merchantNewRecruitment.Status = "Pending"
	err = Dbcon.Save(&merchantNewRecruitment).Scan(&res).Error
	if err != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database SFA create data:", err))
		return res, err
	}

	return res, nil
}

// MerchantNewRecruitmentDetail ..
func (database *DbPostgres) MerchantNewRecruitmentDetail(id string) (res models.MerchantNewRecruitmentExportRes, err error) {
	fmt.Println(">>> MerchantNewRecruitment - MerchantNewRecruitmentDetail - Postgres <<<")
	sugarLogger := database.General.OttoZaplog

	db := Dbcon.Table("merchant_new_recruitments mnr").
		Select("mnr.*,"+
			" p.name province_name, "+
			" c.name city_name, "+
			" d.name district_name, "+
			" v.name village_name").
		Joins(" LEFT JOIN provinces p on p.id = mnr.province_id").
		Joins(" LEFT JOIN cities c on c.id = mnr.city_id").
		Joins(" LEFT JOIN districts d on d.id = mnr.district_id").
		Joins(" LEFT JOIN villages v on v.id = mnr.village_id").
		Where(" mnr.id = ?", id)

	err = db.Scan(&res).Error
	if err != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database SFA when get data:", err))
		return res, err
	}

	return res, nil
}

// ValidationCheckCustomerCodeStatusPending ..
func (database *DbPostgres) ValidationCheckCustomerCodeStatusPending(customerCode string) (res dbmodels.MerchantNewRecruitments, err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - ValidationCheckCustomerCodeStatusPending - Postgres <<<")
	err = Dbcon.Where("customer_code = ? AND status = 'Pending'", customerCode).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// ValidationCheckCustomerCode ..
func (database *DbPostgres) ValidationCheckCustomerCode(customerCode string) (res dbmodels.MerchantNewRecruitments, err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - ValidationCheckCustomerCode - Postgres <<<")
	err = Dbcon.Where("customer_code = ?", customerCode).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// ValidationCheckPhoneNumberStatusPending ..
func (database *DbPostgres) ValidationCheckPhoneNumberStatusPending(phoneNumber string) (res dbmodels.MerchantNewRecruitments, err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - ValidationCheckPhoneNumberStatusPending - Postgres <<<")
	err = Dbcon.Where("phone_number = ? AND status = 'Pending'", phoneNumber).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// ValidationCheckPhoneNumber ..
func (database *DbPostgres) ValidationCheckPhoneNumber(phoneNumber string) (res dbmodels.MerchantNewRecruitments, err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - ValidationCheckPhoneNumber - Postgres <<<")
	err = Dbcon.Where("phone_number = ?", phoneNumber).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}

// UpdateByCustomerCode ..
func (database *DbPostgres) UpdateByCustomerCode(req dbmodels.MerchantNewRecruitments) (err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - UpdateByCustomerCode - Postgres <<<")
	err = Dbcon.Exec("UPDATE merchant_new_recruitments SET name = ? , phone_number = ?, institution_code = ?, "+
		"sub_area_channel_id =?, sub_area_channel_name = ?, owner_name = ?, address = ?, "+
		"longitude = ? , latitude = ?, province_id = ?, city_id = ?, district_id = ?, village_id =?, " +
		"sales_type_id=?, "+
		"updated_at = now() WHERE customer_code = ?",
		req.Name, req.PhoneNumber, req.InstitutionCode,
		req.SubAreaChannelID, req.SubAreaChannelName, req.OwnerName, req.Address,
		req.Longitude, req.Latitude, req.ProvinceId, req.CityId, req.DistrictId, req.VillageId,
		req.SalesTypeId,
		req.CustomerCode).Error
	if err != nil {
		return err
	}
	return nil
}

// UpdateByPhoneNumber (Customer Code is Null)
func (database *DbPostgres) UpdateByPhoneNumber(req dbmodels.MerchantNewRecruitments) (err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - UpdateByPhoneNumber - Postgres <<<")
	err = Dbcon.Exec("UPDATE merchant_new_recruitments SET name = ? , customer_code = ?, institution_code = ?, "+
		"sub_area_channel_id =?, sub_area_channel_name = ?, owner_name = ?, address = ?, "+
		"longitude = ? , latitude = ?, province_id = ?, city_id = ?, district_id = ?, village_id =?, "+
		"updated_at = now() WHERE phone_number = ?",
		req.Name, req.CustomerCode, req.InstitutionCode,
		req.SubAreaChannelID, req.SubAreaChannelName, req.OwnerName, req.Address,
		req.Longitude, req.Latitude, req.ProvinceId, req.CityId, req.DistrictId, req.VillageId,
		req.PhoneNumber).Error
	if err != nil {
		return err
	}
	return nil
}

// GetSalesTypeIDBySubArea ..
func (database *DbPostgres) GetSalesTypeIDBySubArea(subAreaChannelID int64) (res dbmodels.SalesAreaChannelsSubArea, err error) {
	fmt.Println(">>> MerchantNewRecruitmentCreate - GetSalesTypeIDBySubArea - Postgres <<<")
	err = Dbcon.Where("sub_area_id = ?", subAreaChannelID).First(&res).Error
	if err != nil {
		return res, err
	}
	return res, nil
}
