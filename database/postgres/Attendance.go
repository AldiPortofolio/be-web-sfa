package postgres

import (
	"fmt"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/jinzhu/gorm"
)

// GetSessionCoverageArea ..
func (database *DbPostgres) GetSessionCoverageArea(adminID uint) (subAreaId string, err error) {
	fmt.Println(">>> Attendance - GetSessionCoverageArea - Postgres <<<")
	// sugarLogger := database.General.OttoZaplog

	var adminSubArea dbmodels.AdminSubArea
	err = Dbcon.Where("admin_id= ?", adminID).Find(&adminSubArea).Error
	if err != nil {
		fmt.Println("err", err)
		// sugarLogger.Info(fmt.Sprintf("Failed connect to database SFA when get data coverage area sakes:", err))
		return "", err
	}
	return adminSubArea.SubAreaIDs, nil
}

// GetFilterAttendanceByCoverageArea ..
func (database *DbPostgres) GetFilterAttendanceByCoverageArea(params models.AttendanceReq, coverageArea string) ([]dbmodels.Attendance, int64, error) {
	fmt.Println(">>> Attendance - GetFilterAttendanceByCoverageArea - Postgres <<<")

	var attendances []dbmodels.Attendance
	var total TotalRow
	fmt.Println("trims ---> ", coverageArea)

	salesIDs := GetSalesBySubArea(coverageArea)

	limit := int64(25)
	page := params.Page

	sqlList := make(chan string)
	sqlCount := make(chan string)

	if len(salesIDs) > 0 {
		go GenerateQueryListAttendanceCoverageArea("SELECT * ", params, sqlList, " ORDER BY created_at desc ", salesIDs)
		go GenerateQueryListAttendanceCoverageArea("SELECT count(*) total ", params, sqlCount, "", salesIDs)
	} else {
		go GenerateQueryListAttendance("SELECT * ", params, sqlList, " ORDER BY created_at desc ")
		go GenerateQueryListAttendance("SELECT count(*) total ", params, sqlCount, "")
	}

	order := ""
	if limit != 0 {
		order = fmt.Sprintf(" OFFSET %d LIMIT %d", (page-1)*limit, limit)
	}

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &attendances, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		logs.Error("Failed connect to database", rows.Error)
		return attendances, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		logs.Error("Failed connect to database", rowsCount.Error)
		return attendances, 0, rowsCount.Error
	}

	return attendances, total.Total, nil
}

// GetSalesBySubArea ...
func GetSalesBySubArea(coverageArea string) []uint {
	fmt.Println(">>> Attendance - GetSalesBySubArea - Postgres <<<")

	subAreaIDs := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(coverageArea)), ", "), "[]")
	var sales []dbmodels.Salesman
	query := "select s.id from salesmen s join positions p on p.salesman_id = s.id join sub_areas sa on p.regionable_id = sa.id where p.regionable_type = 'SubArea' AND sa.id IN (" + subAreaIDs + ") "
	Dbcon.Raw(query).Find(&sales)

	ids := []uint{}
	for _, s := range sales {
		ids = append(ids, s.ID)
	}
	fmt.Println("salesid =", ids)
	return ids
}

// GenerateQueryListAttendanceCoverageArea ..
func GenerateQueryListAttendanceCoverageArea(query string, req models.AttendanceReq, strChan chan string, order string, salesIDs []uint) {
	fmt.Println(">>> Attendance - GenerateQueryListAttendanceCoverageArea - Postgres <<<")

	sID := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(salesIDs)), ", "), "[]")
	sql := query + " from attendances where sales_id in (" + sID + ") "

	// where := " WHERE "

	if req.ID != "" || req.SalesName != "" || req.SalesPhone != "" || req.Category != "" || req.Type != "" || req.DateFrom != "" || req.DateTo != "" || req.Notes != "" || req.Status != "" || req.Keyword != ""{
		sql = sql + " AND "
	}

	paramString := []string{}

	if req.ID != "" {
		queryID := " id = " + req.ID
		paramString = append(paramString, queryID)
	}

	if req.SalesName != "" {
		querySalesName := " sales_name LIKE '%" + req.SalesName + "%'"
		paramString = append(paramString, querySalesName)
	}

	if req.DateFrom != "" && req.DateTo != "" {
		paramString = append(paramString, fmt.Sprintf(" clocktime_server >= '%s' and clocktime_server <= '%s'", req.DateFrom, req.DateTo))
	}

	if req.SalesPhone != "" {
		querySalesPhone := " sales_phone LIKE '%" + req.SalesPhone + "%'"
		paramString = append(paramString, querySalesPhone)
	}

	if req.Category != "" {
		queryCategory := " attend_category ilike '%" + req.Category + "%'"
		paramString = append(paramString, queryCategory)
	}

	if req.Type != "" {
		paramString = append(paramString, fmt.Sprintf(" type_attendance = '%s'", req.Type))
	}

	if req.Notes != "" {
		queryNotes := " notes LIKE '%" + req.Notes + "%'"
		paramString = append(paramString, queryNotes)
	}

	if req.Status != "" {
		paramString = append(paramString, fmt.Sprintf(" status = '%s'", req.Status))
	}

	if req.Keyword != ""{
		queryKeyword := " (sales_phone ilike '%" + req.Keyword + "%' or sales_name ilike '%" + req.Keyword + "%' or id::text = '"+req.Keyword+"' )"
		paramString = append(paramString, queryKeyword)
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	sql = sql + paramFilter + order

	strChan <- sql
}

// GenerateQueryListAttendance ..
func GenerateQueryListAttendance(query string, req models.AttendanceReq, strChan chan string, order string) {
	fmt.Println(">>> Attendance - GenerateQueryListAttendance - Postgres <<<")

	sql := query + " from attendances "

	where := " WHERE"

	if req.ID != "" || req.SalesName != "" || req.SalesPhone != "" || req.Category != "" || req.Type != "" || req.DateFrom != "" || req.DateTo != "" || req.Notes != "" || req.Status != "" || req.Keyword != "" {
		sql = sql + where
	}

	paramString := []string{}

	if req.ID != "" {
		queryID := " id = " + req.ID
		paramString = append(paramString, queryID)
	}

	if req.SalesName != "" {
		querySalesName := " sales_name LIKE '%" + req.SalesName + "%'"
		paramString = append(paramString, querySalesName)
	}

	if req.DateFrom != "" && req.DateTo != "" {
		paramString = append(paramString, fmt.Sprintf(" clocktime_server >= '%s' and clocktime_server <= '%s'", req.DateFrom, req.DateTo))
	}

	if req.SalesPhone != "" {
		querySalesPhone := " sales_phone LIKE '%" + req.SalesPhone + "%'"
		paramString = append(paramString, querySalesPhone)
	}

	if req.Category != "" {
		queryCategory := " attend_category ilike '%" + req.Category + "%'"
		paramString = append(paramString, queryCategory)
	}

	if req.Type != "" {
		paramString = append(paramString, fmt.Sprintf(" type_attendance = '%s'", req.Type))
	}

	if req.Notes != "" {
		queryNotes := " notes LIKE '%" + req.Notes + "%'"
		paramString = append(paramString, queryNotes)
	}

	if req.Status != "" {
		paramString = append(paramString, fmt.Sprintf(" status = '%s'", req.Status))
	}

	if req.Keyword != ""{
		queryKeyword := " (sales_phone ilike '%" + req.Keyword + "%' or sales_name ilike '%" + req.Keyword + "%' or id::text = '"+req.Keyword+"' )"
		paramString = append(paramString, queryKeyword)
	}

	paramFilter := strings.Join(paramString[:], " AND ")

	sql = sql + paramFilter + order

	strChan <- sql
}

// GetFilterAttendance ..
func (database *DbPostgres) GetFilterAttendance(params models.AttendanceReq) ([]dbmodels.Attendance, int64, error) {
	fmt.Println(">>> Attendance - GetFilterAttendance - Postgres <<<")

	var attendances []dbmodels.Attendance
	var total TotalRow

	limit := int64(25)
	page := params.Page

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListAttendance("SELECT * ", params, sqlList, " ORDER BY created_at desc ")
	go GenerateQueryListAttendance("SELECT count(*) total ", params, sqlCount, "")

	order := ""
	if limit != 0 {
		order = fmt.Sprintf(" OFFSET %d LIMIT %d", (page-1)*limit, limit)
	}

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &attendances, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		logs.Error("Failed connect to database", rows.Error)
		return attendances, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		logs.Error("Failed connect to database", rowsCount.Error)
		return attendances, 0, rowsCount.Error
	}

	return attendances, total.Total, nil
}

// GetAttendanceDetail ..
func (database *DbPostgres) GetAttendanceDetail(attendID int) (dbmodels.Attendance, error) {
	fmt.Println(">>> Attendance - GetAttendanceDetail - Postgres <<<")

	var attendance dbmodels.Attendance

	err := Dbcon.Where("id = ? ", attendID).Find(&attendance).Error

	if err != nil {
		log.Println("Failed to show attendance detail: ", err)
		logs.Error(fmt.Sprintf("Failed to show attendance detail: %v", err))
		return attendance, err
	}

	return attendance, nil
}

// UpdateStatusAttendance ..
func (database *DbPostgres) UpdateStatusAttendance(req models.ValidateAttendanceReq) error {
	fmt.Println(">>> Attendance - UpdateStatusAttendance - Postgres <<<")
	sugarLogger := database.General.OttoZaplog
	var err error
	fmt.Println("req", req)
	err = Dbcon.Exec("update attendances set status = ?, reason = ?, updated_at = now() where id = ?", req.StatusAfter, req.Reason, req.AttendanceId).Error
	if err != nil {
		fmt.Println("err", err)
		sugarLogger.Error("Failed connect to database SFA when updating long lat call plan merchant")
		return err
	}
	return nil
}

// GetFilterAttendanceExport ..
func (database *DbPostgres) GetFilterAttendanceExport(params models.AttendanceReq) ([]dbmodels.Attendance, int64, error) {
	fmt.Println(">>> Attendance - GetFilterAttendanceExport - Postgres <<<")
	var attendances []dbmodels.Attendance
	var total TotalRow

	limit := int64(params.Limit)
	page := params.Page

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListAttendance("SELECT * ", params, sqlList, " ORDER BY created_at desc ")
	go GenerateQueryListAttendance("SELECT count(*) total ", params, sqlCount, "")

	order := ""
	if limit != 0 {
		order = fmt.Sprintf(" OFFSET %d LIMIT %d", (page-1)*limit, limit)
	}

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &attendances, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		logs.Error("Failed connect to database", rows.Error)
		return attendances, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		logs.Error("Failed connect to database", rowsCount.Error)
		return attendances, 0, rowsCount.Error
	}

	return attendances, total.Total, nil
}

// ExportAttendance ..
func (database *DbPostgres) ExportAttendance(data []dbmodels.Attendance) ([]models.ExportAttendance, error) {
	fmt.Println(">>> Attendance - v - Postgres <<<")
	sugarLogger := database.General.OttoZaplog
	var res []models.ExportAttendance
	var paramReg models.ParameterConfiguration

	Dbcon.Raw("select id, name, param_value from parameter_configurations where name = 'face_recog_atd_percentage' ").Scan(&paramReg)

	for i, val := range data {
		var assignSales []dbmodels.AssignSales
		var sales dbmodels.Salesman

		serr := Dbcon.Where("phone_number = ?", val.SalesPhone).Find(&sales).Error
		if serr != nil {
			sugarLogger.Error(fmt.Sprintf("Failed to get sales : %v", serr))
		}

		salesID := strconv.Itoa(int(sales.ID))

		query := "SELECT r.id assign_id, r.code assign_code, r.name assign_name, string_agg(distinct c.name, ';') city_name, string_agg(distinct b.branch_office, ';') branch_office " +
			"from positions p inner join regions r on p.regionable_id = r.id " +
			"inner join branches b on r.id = b.region_id " +
			"inner join branches_cities bc on b.id = bc.branch_id " +
			"inner join cities c on bc.city_id = c.id " +
			"where p.salesman_id = " + salesID + " group by r.id union all " +
			"select b.id assign_id, b.code assign_code, b.name assign_name, string_agg(distinct c.name, ';') city_name, string_agg(distinct b.branch_office, ';') branch_office " +
			"from positions p inner join branches b on p.regionable_id = b.id " +
			"inner join branches_cities bc on b.id = bc.branch_id  " +
			"inner join cities c on bc.city_id = c.id " +
			"where p.salesman_id = " + salesID + " group by b.id union all " +
			"SELECT a.id assign_id, a.code assign_code, a.name assign_name, string_agg(distinct c.name, ';') city_name, string_agg(distinct b.branch_office, ';') branch_office " +
			"from positions p inner join areas a on p.regionable_id = a.id " +
			"inner join branches b on a.branch_id = b.id " +
			"inner join branches_cities bc on b.id = bc.branch_id  " +
			"inner join cities c on bc.city_id = c.id " +
			"where p.salesman_id = " + salesID + " group by a.id union all " +
			"select s.id assign_id, s.code assign_code, s.name assign_name, string_agg(distinct c.name, ';') city_name, string_agg(distinct b.branch_office, ';') branch_office " +
			"from positions p inner join sub_areas s on p.regionable_id = s.id " +
			"inner join areas a on s.area_id = a.id " +
			"inner join branches b on a.branch_id = b.id " +
			"inner join branches_cities bc on b.id = bc.branch_id  " +
			"inner join cities c on bc.city_id = c.id " +
			"where p.salesman_id = " + salesID + " group by s.id "

		perr := Dbcon.Raw(query).Scan(&assignSales).Error
		if perr != nil {
			sugarLogger.Error(fmt.Sprintf("Failed to get assigment sales : %v", perr))
		}

		var salesAssignCodes []string
		var salesAssignNames []string
		var salesCityNames []string
		var salesBranchOffices []string

		for _, assign := range assignSales {
			salesAssignCodes = append(salesAssignCodes, assign.AssignCode)
			salesAssignNames = append(salesAssignNames, assign.AssignName)
			cityNames := strings.Split(assign.CityName, ";")
			branchOffices := strings.Split(assign.BranchOffice, ";")
			salesCityNames = append(salesCityNames, cityNames...)
			salesBranchOffices = append(salesBranchOffices, branchOffices...)
		}

		no := strconv.Itoa(i + 1)
		attendanceID := int(val.ID)

		var code string
		switch val.AttendCategoryType {
		case "0":
			code = "All"
		case "1":
			code = "In"
		case "2":
			code = "Out"
		default:
			code = "All"
		}

		reqStatus := ""

		accuration, _ := strconv.ParseFloat(val.PhotoAccuration, 64)
		paramValue, _ := strconv.ParseFloat(paramReg.ParamValue, 64)
		if accuration != 0 && accuration < paramValue {
			reqStatus = "NOT PASS"
		} else if accuration > paramValue {
			reqStatus = "PASS"
		}

		a := models.ExportAttendance{
			No:                    no,
			AttendanceID:          strconv.Itoa(attendanceID),
			SalesID:               strconv.Itoa(val.SalesID),
			SalesPhone:            val.SalesPhone,
			SalesName:             val.SalesName,
			SalesDepartment:       sales.Occupation,
			SalesAssignmentCode:   strings.Join(salesAssignCodes, "|"),
			SalesAssignmentName:   strings.Join(salesAssignNames, "|"),
			SalesAssignmentCities: strings.Join(salesCityNames, "|"),
			//SalesAssignmentCities:  strings.Join(UniqStr(salesCityNames), "|"),
			CompanyCode:            sales.CompanyCode,
			AttendanceCategory:     val.AttendCategory,
			AttendanceCategoryType: code,
			ClockTimeServer:        val.ClocktimeServer.Format("2006-01-02 15:04:05"),
			ClockTimeLocal:         val.ClocktimeLocal.Format("2006-01-02 15:04:05"),
			TypeAttendance:         val.TypeAttendance,
			Selfie:                 val.Selfie,
			Location:               val.Location,
			Long:                   val.Longitude,
			Lat:                    val.Latitude,
			Notes:                  val.Notes,
			BranchOffices:          strings.Join(salesBranchOffices, "|"),
			PhotoAccuration:        val.PhotoAccuration,
			MinAccuration:          paramReg.ParamValue,
			AccurationStatus:       reqStatus,
			StatusName:             utils.StatusAttendance(val.Status),
		}

		res = append(res, a)
	}

	return res, nil
}
