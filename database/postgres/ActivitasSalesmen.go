package postgres

import (
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

// ListActivitasSalesmen ..
func (database *DbPostgres) ListActivitasSalesmen(req dbmodels.ListActivitasSalesmenReq) ([]dbmodels.ActivitasSalesmenList, int64, error) {
	fmt.Println(">>> Activitas Salesmen - ListActivitasSalesmen - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res []dbmodels.ActivitasSalesmenList
	var total TotalRow
	limit := int64(25)

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListActivitasSalesmen(false, req, sqlList, " order by s.action_date desc")
	go GenerateQueryListActivitasSalesmen(true, req, sqlCount, "")

	if req.Page == 0 {
		req.Page = 1
	}

	order := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*limit, limit)

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &res, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rows.Error))
		return res, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rowsCount.Error))
		return res, 0, rowsCount.Error
	}

	return res, total.Total, nil
}

// ListActivitasSalesmenByCoverage ..
func (database *DbPostgres) ListActivitasSalesmenByCoverage(req dbmodels.ListActivitasSalesmenReq, coverageArea string) ([]dbmodels.ActivitasSalesmenList, int64, error) {
	fmt.Println(">>> Activitas Salesmen - ListActivitasSalesmen - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res []dbmodels.ActivitasSalesmenList
	var total TotalRow
	limit := int64(25)

	sqlList := make(chan string)
	sqlCount := make(chan string)

	salesIDs := GetSalesBySubArea(coverageArea)

	if len(salesIDs) > 0 {
		go GenerateQueryListActivitasSalesmenByCoverageArea(false, req, sqlList, " order by s.action_date desc", salesIDs)
		go GenerateQueryListActivitasSalesmenByCoverageArea(true, req, sqlCount, "", salesIDs)
	} else {
		go GenerateQueryListActivitasSalesmen(false, req, sqlList, " order by s.action_date desc")
		go GenerateQueryListActivitasSalesmen(true, req, sqlCount, "")
	}

	if req.Page == 0 {
		req.Page = 1
	}

	order := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*limit, limit)

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, order, &res, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rows.Error))
		return res, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rowsCount.Error))
		return res, 0, rowsCount.Error
	}

	return res, total.Total, nil
}

// GenerateQueryListActivitasSalesmen ..
func GenerateQueryListActivitasSalesmenByCoverageArea(isCount bool, req dbmodels.ListActivitasSalesmenReq, strChan chan string, order string, salesIDs []uint) {

	currentTime := time.Now()
	periodFrom := currentTime.AddDate(0, -2, 0).Format("2006-01-02")
	periodTo := currentTime.Format("2006-01-02")

	if req.PeriodFrom != "" && req.PeriodTo != "" {
		periodFrom = req.PeriodFrom
		periodTo = req.PeriodTo
	}

	where := ""
	if req.Keyword != "" {
		where += ` and (
					s.name ilike LOWER('%` + req.Keyword + `%') or
					s.phone_number ilike '%` + req.Keyword + `%' or s.id::text ilike '%` + req.Keyword + `%'
				)`
	}

	if req.SalesTypeID != "" {
		where += ` and st.id  = '` + req.SalesTypeID + `'`
	}

	sID := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(salesIDs)), ", "), "[]")

	sql := ""
	if !isCount {
		sql = `
			WITH salesmen_table AS (
				SELECT 
				i::date as action_date,  
				s.id,  
				s.sales_id,  
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as name,
				s.phone_number,
				s.sales_type_id
				FROM salesmen s
				CROSS JOIN generate_series('` + periodFrom + `'::date, '` + periodTo + `', '1 day') t(i)
			) 
			SELECT 
			s.action_date, 
			s.id, 
			s.sales_id, 
			s.name, 
			s.phone_number,
			st."name" as sales_type,
			(select count(m.*) as total from merchants m  where date(m.created_at) = s.action_date and m.salesman_id = s.id) as akusisi,
			(
			select count(distinct s2.id) as total from suppliers s2 
			left join ottopay_orders oo on oo.supplier_id = s2.id
			where date(s2.created_at) = s.action_date and oo.salesman_id = s.id
			) as noo,
			(
				select count(t.*) as total from todolists t
				where date(t.action_date) = s.action_date and t.sales_phone = s.phone_number 
			) as todolist_count,
			cp.success_call,
			(
				select Count(*) as total from call_plans cp 
				left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
				where 
				cpm.status = 'Completed'
				and
				date(cpm.action_date) = s.action_date and cp.sales_id = s.id
			) as success_callplan_count,
			(
				select Count(*) as total from call_plans cp 
				left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
				where 
				cpm.status in ('Completed', 'Visited')
				and
				date(cpm.action_date) = s.action_date and cp.sales_id = s.id
			) as total_callplan_count
			FROM salesmen_table AS s
			left join sales_types st on st.id = s.sales_type_id
			left join call_plans cp on cp.sales_id = s.id and date(cp.call_plan_date) = s.action_date 
			where s.id in (` + sID + `) ` + where + `
		`
	} else {
		sql = `
			WITH salesmen_table AS (
				SELECT 
				i::date as action_date,  
				s.id,  
				s.sales_id,  
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as name,
				s.phone_number,
				s.sales_type_id
				FROM salesmen s
				CROSS JOIN generate_series('` + periodFrom + `'::date, '` + periodTo + `', '1 day') t(i)
			) 
			SELECT 
			 count(s.*) as total
			FROM salesmen_table AS s
			left join sales_types st on st.id = s.sales_type_id
			left join call_plans cp on cp.sales_id = s.id and date(cp.call_plan_date) = s.action_date 
			where s.id in (` + sID + `) ` + where + `
		`
	}

	sql += order

	strChan <- sql
}

// GenerateQueryListActivitasSalesmen ..
func GenerateQueryListActivitasSalesmen(isCount bool, req dbmodels.ListActivitasSalesmenReq, strChan chan string, order string) {

	currentTime := time.Now()
	periodFrom := currentTime.AddDate(0, -2, 0).Format("2006-01-02")
	periodTo := currentTime.Format("2006-01-02")

	if req.PeriodFrom != "" && req.PeriodTo != "" {
		periodFrom = req.PeriodFrom
		periodTo = req.PeriodTo
	}

	where := ""
	if req.Keyword != "" {
		where += ` and (
					s.name ilike LOWER('%` + req.Keyword + `%') or
					s.phone_number ilike '%` + req.Keyword + `%' or s.id::text ilike '%` + req.Keyword + `%'
				)`
	}

	if req.SalesTypeID != "" {
		where += ` and st.id  = '` + req.SalesTypeID + `'`
	}

	sql := ""
	if !isCount {
		sql = `
			WITH salesmen_table AS (
				SELECT 
				i::date as action_date,  
				s.id,  
				s.sales_id,  
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as name,
				s.phone_number,
				s.sales_type_id
				FROM salesmen s
				CROSS JOIN generate_series('` + periodFrom + `'::date, '` + periodTo + `', '1 day') t(i)
			) 
			SELECT 
			s.action_date, 
			s.id, 
			s.sales_id, 
			s.name, 
			s.phone_number,
			st."name" as sales_type,
			(select count(m.*) as total from merchants m  where date(m.created_at) = s.action_date and m.salesman_id = s.id) as akusisi,
			(
			select count(distinct s2.id) as total from suppliers s2 
			left join ottopay_orders oo on oo.supplier_id = s2.id
			where date(s2.created_at) = s.action_date and oo.salesman_id = s.id
			) as noo,
			(
				select count(t.*) as total from todolists t
				where date(t.action_date) = s.action_date and t.sales_phone = s.phone_number 
			) as todolist_count,
			cp.success_call,
			(
				select Count(*) as total from call_plans cp 
				left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
				where 
				cpm.status = 'Completed'
				and
				date(cpm.action_date) = s.action_date and cp.sales_id = s.id
			) as success_callplan_count,
			(
				select Count(*) as total from call_plans cp 
				left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
				where 
				cpm.status in ('Completed', 'Visited')
				and
				date(cpm.action_date) = s.action_date and cp.sales_id = s.id
			) as total_callplan_count
			FROM salesmen_table AS s
			left join sales_types st on st.id = s.sales_type_id
			left join call_plans cp on cp.sales_id = s.id and date(cp.call_plan_date) = s.action_date 
			where 1=1 ` + where + `
		`
	} else {
		sql = `
			WITH salesmen_table AS (
				SELECT 
				i::date as action_date,  
				s.id,  
				s.sales_id,  
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as name,
				s.phone_number,
				s.sales_type_id
				FROM salesmen s
				CROSS JOIN generate_series('` + periodFrom + `'::date, '` + periodTo + `', '1 day') t(i)
			) 
			SELECT 
			 count(s.*) as total
			FROM salesmen_table AS s
			left join sales_types st on st.id = s.sales_type_id
			left join call_plans cp on cp.sales_id = s.id and date(cp.call_plan_date) = s.action_date
			where 1=1 ` + where + `
		`
	}

	sql += order

	strChan <- sql
}

// DetailActivitasSalesmen ..
func (database *DbPostgres) DetailActivitasSalesmen(req dbmodels.DetailActivitasSalesmenReq) (dbmodels.ActivitasSalesmenDetail, error) {
	fmt.Println(">>> Activitas Salesmen - DetailActivitasSalesmen - DB <<<")
	// sugarLogger := database.General.OttoZaplog
	var res dbmodels.ActivitasSalesmenDetail
	query := `select 
				s.id, 
				s.sales_id,
				s.photo,
				cp.call_plan_date, 
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as name,
				st."name" as sales_type,
				s.phone_number,
				(select count(m.id) from merchants m where m.salesman_id = s.id and date(m.created_at) = date('` + req.Date + `') ) as akusisi,
				(select count(distinct s2.id) as total from suppliers s2 
				left join ottopay_orders oo on oo.supplier_id = s2.id
				where date(s2.created_at) = date('` + req.Date + `') and oo.salesman_id  = s.id
				group by oo.salesman_id) as noo,
				(select count(*) from todolists tl where tl.sales_phone = s.phone_number and  date(tl.action_date) =  date('` + req.Date + `') ) as todolist_count,
				(select sum(cpm.amount) as total from call_plans cp 
					left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
					where 
					date(cpm.action_date) = date('` + req.Date + `') and cp.sales_id = s.id
				) as amount,
				cp.success_call as success_callplan_percentage,
				(
					select Count(*) as total from call_plans cp 
					left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
					where 
					cpm.status = 'Completed'
					and
					date(cpm.action_date) = date('` + req.Date + `') and cp.sales_id = s.id
				) as success_callplan_count,
				(
					select Count(*) as total from call_plans cp 
					left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
					where 
					cpm.status in ('Completed', 'Visited')
					and
					date(cpm.action_date) = date('` + req.Date + `') and cp.sales_id = s.id
				) as total_callplan_count,
				cp.sub_area,
				string_agg(distinct v.name, ', ') as villages,
				sa.name as sac 
				from salesmen s 
				left join call_plans cp on cp.sales_id = s.id and date(cp.call_plan_date) = date('` + req.Date + `') 
				left join sales_types st on st.id = s.sales_type_id 
				left join positions p on p.salesman_id = s.id 
				left join sub_areas sa on p.regionable_id = sa.id and p.regionable_type = 'SubArea' 
				left join sub_areas_villages sv on sv.sub_area_id = sa.id 
				left join villages v on v.id = sv.village_id
				where 
				s.id = '` + req.ID + `'
				group by s.id, cp.call_plan_date, st."name", cp.id ,sa.name
			`

	err := Dbcon.Raw(query).Scan(&res).Error

	if err != nil {
		fmt.Println("Failed connect to database :", err.Error)
		return res, err
	}

	return res, nil
}

// DetailListActivitasTodoList ..
func (database *DbPostgres) DetailListActivitasTodoList(req dbmodels.DetailListActivitasTodolistReq) ([]dbmodels.ActivitasSalesmenDetailTodolist, int64, error) {
	fmt.Println(">>> Activitas Salesmen - DetailActivitasSalesmen - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res []dbmodels.ActivitasSalesmenDetailTodolist
	var total TotalRow
	limit := int64(25)

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListTodoListActivitasSalesmen(false, req, sqlList, " order by t.task_date ASC")
	go GenerateQueryListTodoListActivitasSalesmen(true, req, sqlCount, "")

	if req.Page == 0 {
		req.Page = 1
	}

	paging := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*limit, limit)

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, paging, &res, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rows.Error))
		return res, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rowsCount.Error))
		return res, 0, rowsCount.Error
	}

	return res, total.Total, nil
}

// GenerateQueryListTodoListActivitasSalesmen ..
func GenerateQueryListTodoListActivitasSalesmen(isCount bool, req dbmodels.DetailListActivitasTodolistReq, strChan chan string, order string) {

	sql := ""
	if !isCount {
		sql = `select 
				t.id, 
				t.task_date as date, 
				t.merchant_name, 
				t.mid, 
				mt."name" as merchant_type_name, 
				tc."name" as category, 
				t.status 
			`
	} else {
		sql = "select count(*) as total"
	}

	sql += ` from todolists t 
				left join salesmen s on s.phone_number = t.sales_phone
				left join merchants m on m.merchant_id = t.mid 
				left join merchant_new_recruitments mnr on mnr.id  = t.merchant_new_recruitment_id 
				left join merchant_types mt on mt.id  = m.merchant_type_id 
				left join todolist_categories tc on tc.id = t.todolist_category_id 
				where s.id = '` + req.ID + `' and action_date = date('` + req.Date + `')
			`

	if req.Keyword != "" {
		sql += ` and (lower(t.merchant_name) like lower('%` + req.Keyword + `%') or t.mid like '%` + req.Keyword + `%' or t.id::text like '%` + req.Keyword + `%')`
	}

	category := strconv.Itoa(req.CategoryID)
	if category != "0" {
		sql += ` and t.todolist_category_id = '` + category + `'`
	}

	if req.Status != "" {
		sql += ` and t.status = '` + req.Status + `'`
	}

	sql += order

	strChan <- sql
}

// DetailListActivitasCallplan ..
func (database *DbPostgres) DetailListActivitasCallplan(req dbmodels.DetailListActivitasTodolistReq) ([]dbmodels.ActivitasSalesmenDetailCallplan, int64, error) {
	fmt.Println(">>> Activitas Salesmen - DetailActivitasSalesmen - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res []dbmodels.ActivitasSalesmenDetailCallplan
	var total TotalRow
	limit := int64(25)

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListCallplanActivitasSalesmen(false, req, sqlList, " order by cp.call_plan_date  ASC")
	go GenerateQueryListCallplanActivitasSalesmen(true, req, sqlCount, "")

	if req.Page == 0 {
		req.Page = 1
	}

	paging := fmt.Sprintf(" OFFSET %d LIMIT %d", (req.Page-1)*limit, limit)

	rowsc := make(chan *gorm.DB)
	rowsCountc := make(chan *gorm.DB)

	go AsyncRawQuery(<-sqlList, paging, &res, rowsc)
	go AsyncRawQuery(<-sqlCount, "", &total, rowsCountc)

	if rows := <-rowsc; rows.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rows.Error))
		return res, 0, rows.Error
	}

	if rowsCount := <-rowsCountc; rowsCount.Error != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", rowsCount.Error))
		return res, 0, rowsCount.Error
	}

	return res, total.Total, nil
}

// GenerateQueryListTodoListActivitasSalesmen ..
func GenerateQueryListCallplanActivitasSalesmen(isCount bool, req dbmodels.DetailListActivitasTodolistReq, strChan chan string, order string) {

	sql := ""
	if !isCount {
		sql = `select 
				cpm.id, 
				cp.sales_id,
				cp.call_plan_date, 
				date(cpm.action_date) as action_date,
				cpm.merchant_name, 
				cpm.mid, 
				cpm.merchant_type_id, 
				mt."name" as merchant_type_name, 
				cpm.merchant_status, 
				v."name" as kelurahan,
				cpm.status 
			`
	} else {
		sql = "select count(*) as total"
	}

	sql += ` from call_plans cp 
				left join call_plan_merchants cpm on cpm.call_plan_id = cp.id 
				left join merchant_types mt on mt.id = cpm.merchant_type_id 
				left join merchants m on m.merchant_id = cpm.mid
				left join villages v on v.id = m.village_id 
				where cp.sales_id = '` + req.ID + `' and date(cpm.action_date) = date('` + req.Date + `')
			`

	if req.Keyword != "" {
		sql += ` and (lower(cpm.merchant_name) like lower('%` + req.Keyword + `%') or cpm.mid like '%` + req.Keyword + `%')`
	}

	category := strconv.Itoa(req.CategoryID)
	if category != "0" {
		sql += ` and cpm.merchant_type_id = '` + category + `'`
	}

	if req.Status != "" {
		sql += ` and cpm.status = '` + req.Status + `'`
	} else {
		sql += ` and cpm.status in ('Visited', 'Completed')`
	}

	sql += order

	strChan <- sql
}

// DetailActivitasSalesmenCallplan ..
func (database *DbPostgres) DetailCallplan(callPlanMerchantID string) (dbmodels.CallPlanMerchant, error) {
	fmt.Println(">>> Activitas Salesmen - DetailCallplan - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res dbmodels.CallPlanMerchant
	query := `select 
				cpm.id,
				cpm.merchant_name,
				cpm.mid,
				m.owner_name,
				m.merchant_type_id,
				mt."name" as merchant_type_name,
				cpm.merchant_address as address,
				cpm.longitude ,
				cpm.latitude ,
				cpm.merchant_status,
				cp.call_plan_date,
				cpm.action_date,
				cpm.clock_time,
				cpm.status,
				cpm.notes,
				cpm.photo_location
			from call_plan_merchants cpm 
				left join merchants m on m.merchant_id = cpm.mid
				left join merchant_types mt on mt.id  = m.merchant_type_id 
				left join call_plans cp on cp.id = cpm.call_plan_id 
			where cpm.id = ?
			`

	err := Dbcon.Raw(query, callPlanMerchantID).Preload("CallPlanActions").Find(&res).Error

	if err != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", err.Error))
		return res, err
	}

	return res, nil
}

// DetailActivitasSalesmenTodolist ..
func (database *DbPostgres) DetailTodoList(todolisID string) (dbmodels.TodoListDetail, error) {
	fmt.Println(">>> Activitas Salesmen - DetailCallplan - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res dbmodels.TodoListDetail
	query := `select
				t.id,
				t.merchant_name,
				t.mid,
				t.owner_name,
				case 
				when mnr.address is null or mnr.address = '' then m.address
				else mnr.address end as address, 
				t.notes,
				mt."name" as merchant_type_name,
				t.task_date,
				t.action_date,
				t.longitude,
				t.latitude,
				t.status,
				t.created_at,
				t.todolist_category_id,
				coalesce(s.first_name , '') || coalesce(' ' , '') || coalesce(s.last_name , '') as action_by
			from todolists t 
				left join merchant_new_recruitments mnr on t.merchant_new_recruitment_id = mnr.id
				left join merchants m on m.merchant_id = t.mid 
				left join merchant_types mt on mt.id = m.merchant_type_id 
				left join salesmen s on s.phone_number = t.sales_phone
			where t.id = ?
			`

	err := Dbcon.Raw(query, todolisID).Preload("TodolistCategory").Preload("Tasks.FollowUps").Preload("Tasks.TodolistSubCategory").Preload("Tasks.ActionByName").Preload("Tasks.LabelTask").Preload("TodoListHistories").Find(&res).Error

	if err != nil {
		sugarLogger.Info(fmt.Sprintf("Failed connect to database :", err.Error))
		return res, err
	}

	return res, nil
}
