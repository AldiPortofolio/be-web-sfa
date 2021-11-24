package postgres

import (
	"encoding/json"
	"errors"
	"fmt"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"gorm.io/gorm/clause"
	ottologger "ottodigital.id/library/logger"
)

// SaveJobManagement ..
func (database *DbPostgres) SaveJobManagement(req dbmodels.JobManagements) (id int64, err error) {
	fmt.Println(">>> JobManagement - SaveJobManagement - Postgres <<<")
	sugarLogger := ottologger.GetLogger()

	err = Dbcon.Preload(clause.Associations).Save(&req).Error
	if err != nil {
		sugarLogger.Info("Failed Save JobManagement")
		return req.Id, err
	}
	return req.Id, nil
}

// FilterJobManagement ..
func (database *DbPostgres) FilterJobManagement(req models.ReqFilterJobManagements) (res []dbmodels.JobManagements, total int, err error) {
	fmt.Println(">>> JobManagement - FilterJobManagement - Postgres <<<")
	request, _ := json.Marshal(req)
	fmt.Println("Request : ", string(request))
	sugarLogger := ottologger.GetLogger()
	db := GetDbCon()

	page := req.Page
	limit := req.Limit
	
	db = db.Preload("JobDescriptions").Preload("JobCategoryName").Preload("SenderName").Preload("RecipientName")

	db = db.Preload("JobDescriptions").Preload("JobCategoryName").Preload("SenderName").Preload("RecipientName").Find(&res)

	if req.ID != 0 {
		db = db.Where("id = ?", req.ID)
	}

	if req.Status  != 0 {
        db = db.Where("status = ?",  req.Status )
    }

	if req.JobCategoryId != 0 {
		db = db.Where("job_category_id = ?", req.JobCategoryId)
	}

	if req.SenderId != 0 || req.RecipientId != 0 {
		db = db.Where("sender_id = ? or recipient_id = ?",  req.SenderId, req.RecipientId )
	}

	// if req.RecipientId != 0 {
	// 	db = db.Where("recipient_id = ?",  req.RecipientId )
	// }

	if req.JobPriority != "" {
		db = db.Where("job_priority = ?",  req.JobPriority )
	}

	if req.Name != "" {
		db = db.Where("name ilike '%" + req.Name + "%'")
	}

	db = db.Where("status_storage = 'false'")
	// if !req.StatusStorage {
	// 	db = db.Where("status_storage = ?", req.StatusStorage)
	// } else {
	// 	db = db.Where("status_storage = ?", true)
	// }

	if strings.TrimSpace(req.DateStart) != "" && strings.TrimSpace(req.DateEnd) != "" {
		transactionDateLayout := "2006-01-02T15:04:05.000000Z07:00"
		dateStart, errStart := time.Parse("2006-01-02", req.DateStart)
		if errStart != nil {
			fmt.Println("Failed to parse request start date to time:", errStart)
		}

		dateEnd, errEnd := time.Parse("2006-01-02", req.DateEnd)
		dateEnd = dateEnd.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
		if errEnd != nil {
			fmt.Println("Failed to parse request start date to time:", errEnd)
		}

		db = db.Where("created_at >= ? and created_at <= ?", dateStart.Format(transactionDateLayout), dateEnd.Format(transactionDateLayout))
	}

	err = db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		sugarLogger.Info("Failed Get List JobManagement")
		return res, 0, err
	}

	return res, total, nil
}

// FindByCategoryID ...
func (database *DbPostgres) FindByCategoryID(id int64) ([]dbmodels.JobManagements, error) {
	fmt.Println(">>> JobCategories - FindJobCategoriesById - Postgres <<<")
	sugarLogger := ottologger.GetLogger()
	var res []dbmodels.JobManagements
	err := Dbcon.Where(dbmodels.JobManagements{JobCategoryId: id}).Find(&res).Error
	if err != nil {
		sugarLogger.Info("Data not Found")
		return res, err
	}
	return res, err
}

// ListActivitasSalesmen ..
func (database *DbPostgres) ListDraftJobManagements(req models.ReqFilterJobManagementDraft) ([]dbmodels.DraftJobManagements, int64, error) {
	fmt.Println(">>> Activitas Salesmen - ListActivitasSalesmen - DB <<<")
	sugarLogger := database.General.OttoZaplog
	var res []dbmodels.DraftJobManagements
	var total TotalRow
	limit := int64(25)

	sqlList := make(chan string)
	sqlCount := make(chan string)

	go GenerateQueryListJobManagementDraft(false, req, sqlList, " order by jm.created_at ASC")
	go GenerateQueryListJobManagementDraft(true, req, sqlCount, "")

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

// GenerateQueryListJobManagementDraft ..
func GenerateQueryListJobManagementDraft(isCount bool, req models.ReqFilterJobManagementDraft, strChan chan string, order string) {

	sql := ""
	if !isCount {
		sql = `select 
				jm.id, 
				jm.name, 
				jm.job_category_id, 
				jc.name as job_category_name, 
				jm.assignment_date  
			`
	} else {
		sql = "select count(*) as total "
	}

	sql += `from job_managements jm
				left join job_categories jc on jc.id = jm.job_category_id
			where jm.status_storage = true and jm.sender_id = `+req.AdminID+`
			`
	if req.Keyword != "" {
		sql += ` and (lower(jm.name) like lower('%` + req.Keyword + `%') or lower(jc."name") like lower('%` + req.Keyword + `%')) `
	}
	sql += order

	strChan <- sql
}


// DeleteJobCategory ..
func (database *DbPostgres) DeleteJobManagement(id int64) error {
	sugarLogger := ottologger.GetLogger()
	var jobManagements dbmodels.JobManagements
	fmt.Println("id", id)
	sErr := Dbcon.Where("id = ?", id).Delete(&jobManagements).Error
	if sErr != nil {
		sugarLogger.Info(sErr.Error())
		return errors.New("Failed delete admin sub area")
	}

	return nil
}


// FilterJobManagement ..
func (database *DbPostgres) DetailJobManagement(jobID string) (res dbmodels.JobManagements, err error) {
	fmt.Println(">>> JobManagement - FilterJobManagement - Postgres <<<")
	// request, _ := json.Marshal(req)
	// fmt.Println("Request : ", string(request))
	sugarLogger := ottologger.GetLogger()
	db := GetDbCon()

	
	err = db.Where("id = ?", jobID).Preload("JobDescriptions").Preload("JobCategoryName").Preload("SenderName").Preload("RecipientName").First(&res).Error



	if err != nil {
		sugarLogger.Info("Failed Get List JobManagement")
		return res, err
	}

	return res, nil
}

// FilterJobManagement ..
func (database *DbPostgres) FilterJobManagementEdit(req models.ReqFilterJobManagements) (res []dbmodels.JobManagements, total int, err error) {
	fmt.Println(">>> JobManagement - FilterJobManagement - Postgres <<<")
	request, _ := json.Marshal(req)
	fmt.Println("Request : ", string(request))
	sugarLogger := ottologger.GetLogger()
	db := GetDbCon()

	page := req.Page
	limit := req.Limit
	
	db = db.Preload("JobDescriptions").Preload("JobCategoryName").Preload("SenderName").Preload("RecipientName")

	db = db.Preload("JobDescriptions").Preload("JobCategoryName").Preload("SenderName").Preload("RecipientName").Find(&res)

	if req.ID != 0 {
		db = db.Where("id = ?", req.ID)
	}

	if req.Status  != 0 {
        db = db.Where("status = ?",  req.Status )
    }

	if req.JobCategoryId != 0 {
		db = db.Where("job_category_id = ?", req.JobCategoryId)
	}

	if req.SenderId != 0 || req.RecipientId != 0 {
		db = db.Where("sender_id = ? or recipient_id = ?",  req.SenderId, req.RecipientId )
	}

	// if req.RecipientId != 0 {
	// 	db = db.Where("recipient_id = ?",  req.RecipientId )
	// }

	if req.JobPriority != "" {
		db = db.Where("job_priority = ?",  req.JobPriority )
	}

	if req.Name != "" {
		db = db.Where("name ilike '%" + req.Name + "%'")
	}

	// db = db.Where("status_storage = 'false'")
	// if !req.StatusStorage {
	// 	db = db.Where("status_storage = ?", req.StatusStorage)
	// } else {
	// 	db = db.Where("status_storage = ?", true)
	// }

	if strings.TrimSpace(req.DateStart) != "" && strings.TrimSpace(req.DateEnd) != "" {
		transactionDateLayout := "2006-01-02T15:04:05.000000Z07:00"
		dateStart, errStart := time.Parse("2006-01-02", req.DateStart)
		if errStart != nil {
			fmt.Println("Failed to parse request start date to time:", errStart)
		}

		dateEnd, errEnd := time.Parse("2006-01-02", req.DateEnd)
		dateEnd = dateEnd.Add(time.Hour * 23).Add(time.Minute * 59).Add(time.Second * 59)
		if errEnd != nil {
			fmt.Println("Failed to parse request start date to time:", errEnd)
		}

		db = db.Where("created_at >= ? and created_at <= ?", dateStart.Format(transactionDateLayout), dateEnd.Format(transactionDateLayout))
	}

	err = db.Limit(limit).Offset((page - 1) * limit).Order("id desc").Find(&res).Limit(-1).Offset(0).Count(&total).Error

	if err != nil {
		sugarLogger.Info("Failed Get List JobManagement")
		return res, 0, err
	}

	return res, total, nil
}
