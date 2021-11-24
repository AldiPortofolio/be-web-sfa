package postgres

import (
	"errors"
	"fmt"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
	"time"
)

// TodolistDetail ..
func (database *DbPostgres) TodolistDetail(todolistID string) (models.ShowTodolist, error) {
	fmt.Println(">>> Todolist - NewMerchantDetail - Postgres <<<")
	var todolist models.Todolist
	var todolistSer models.ShowTodolist
	// var tasks []dbmodels.Task
	var subArea dbmodels.SubArea
	var category dbmodels.TodolistCategory
	var salesmen []dbmodels.Salesman

	todoErr := Dbcon.Where("id = ?", todolistID).Preload("Tasks").Preload("TodolistHistories").Find(&todolist).Error
	histories := todolist.TodolistHistories
	// tasksErr := Dbcon.Raw("select * from tasks WHERE todolist_id = ? ORDER BY updated_at DESC", todolist_id).Scan(&tasks).Error

	data := map[string]interface{}{}
	var customerCode, merchantPhone string
	if todolist.TodolistCategoryID == 6 {
		merchantDetail, err := GetNewMerchantDetailByID(todolist.MerchantNewRecruitmentID)
		if err == nil {
			data["merchant"] = merchantDetail
			todolist.MerchantName = merchantDetail.MerchantPhone
			customerCode = merchantDetail.CustomerCode
			merchantPhone = merchantDetail.MerchantPhone
		}
	} else {
		merchantDetail := GetMerchantDetail(todolist.Mid)
		data["merchant"] = merchantDetail
		todolist.MerchantName = merchantDetail.MerchantName
	}

	categoryErr := Dbcon.Where("id = ?", todolist.TodolistCategoryID).Find(&category).Error
	subErr := Dbcon.Raw("select * from sub_areas INNER JOIN sub_areas_villages ON sub_areas.id = sub_areas_villages.sub_area_id"+
		" WHERE sub_areas_villages.village_id = ? ", todolist.VillageID).Scan(&subArea).Error

	if todoErr != nil {
		log.Println("Todolist not found : ", todoErr)
		return todolistSer, todoErr
	}

	if categoryErr != nil {
		log.Println("Failed to get category : ", categoryErr)
		return todolistSer, categoryErr
	}

	if subErr != nil {
		log.Println("Failed to get subArea : ", subErr)
	}

	var tasksAttributes []models.TaskDetail
	for _, task := range todolist.Tasks {
		var subCategory dbmodels.TodolistSubCategory
		var followUps []dbmodels.FollowUp
		var sales dbmodels.Salesman

		subErr := Dbcon.Where("id = ?", task.TodolistSubCategoryID).Find(&subCategory).Error
		if subErr != nil {
			log.Println("Failed to get sub category : ", subErr)
		}

		fwErr := Dbcon.Raw("SELECT * FROM follow_ups WHERE task_id = ?", task.ID).Scan(&followUps).Error
		if fwErr != nil {
			log.Println("Failed to get followUps : ", fwErr)
		}

		if task.ActionBy != "" {
			salesErr := Dbcon.Raw("SELECT s.first_name, s.last_name from salesmen s WHERE s.phone_number = ?", task.ActionBy).Scan(&sales).Error
			if salesErr != nil {
				log.Println("Failed to get sales : ", salesErr)
			}
		}

		var taskSer models.TaskDetail
		taskSer = SerializedTaskFollowUp(task, followUps, subCategory)
		taskSer.ActionBy = sales.FirstName + " " + sales.LastName
		tasksAttributes = append(tasksAttributes, taskSer)
	}

	var sales dbmodels.Salesman
	if todolist.SalesPhone != "" {
		serr := Dbcon.Where("phone_number = ?", todolist.SalesPhone).Find(&sales).Error
		if serr != nil {
			log.Println("Failed to get sales : ", serr)
		}
	}

	if todolist.SalesPhone == "" {
		salesErr := Dbcon.Raw("SELECT * FROM salesmen INNER JOIN positions ON positions.salesman_id = salesmen.id"+
			" WHERE positions.regionable_id = ? AND positions.regionable_type = ?", subArea.ID, "SubArea").Scan(&salesmen).Error

		if salesErr != nil {
			log.Println("Failed to get sales : ", salesErr)
		}

	}

	todolistSer = SerializedTodolistDetail(todolist, data["merchant"], tasksAttributes, histories, salesmen, sales)
	todolistSer.CategoryName = category.Name
	todolistSer.CreatedAt = todolist.CreatedAt
	todolistSer.ActionBy = sales.FirstName + " " + sales.LastName
	todolistSer.SalesPhone = todolist.SalesPhone
	todolistSer.CustomerCode = customerCode
	todolistSer.MerchantPhone = merchantPhone

	return todolistSer, nil
}

// TodolistNewMerchantDetail ..
func (database *DbPostgres) TodolistNewMerchantDetail(params models.NewMerchantDetailReq) (models.NewMerchantDetail, error) {
	fmt.Println(">>> Todolist - NewMerchantDetail - Postgres <<<")
	var res models.NewMerchantDetail

	defQuery := "select id, name as merchant_name, customer_code, owner_name, sub_area_channel_name as sub_area, address, " +
		"phone_number as merchant_phone, longitude, latitude, village_id " +
		" from merchant_new_recruitments "
	var filter string
	if params.CustomerCode != "" {
		filter = " where customer_code = '" + params.CustomerCode + "'"
	}

	if params.CustomerCode == "" && params.PhoneNumber != "" {
		filter = " where phone_number = '" + params.PhoneNumber + "'"
	}

	query := defQuery + filter
	err := Dbcon.Raw(query).Scan(&res).Error

	if err != nil {
		log.Println("Failed to get New Merchant detail: ", err)
		return res, err
	}

	res.MerchantID = strconv.Itoa(int(res.ID))
	return res, err
}

// TodolistCreate ..
func (database *DbPostgres) TodolistCreate(todo models.CreateTodolist) (models.TodolistDetail, error) {
	fmt.Println(">>> Todolist - TodolistCreate - Postgres <<<")
	var todolist dbmodels.Todolist

	taskDate, _ := time.Parse("2006-01-02", todo.TaskDate)
	todolist.Mid = todo.Mid
	todolist.SalesPhone = todo.SalesPhone
	todolist.TodolistCategoryID = todo.TodolistCategoryID
	todolist.VillageID = todo.VillageID
	todolist.Notes = todo.Notes
	// todolist.CustomerCode = todo.CustomerCode
	// todolist.MerchantPhone = todo.MerchantPhone
	todolist.MerchantName = todo.MerchantName
	todolist.TaskDate = taskDate
	todolist.Status = "Open"

	if todo.TodolistCategoryID == 6 {
		params := models.NewMerchantDetailReq{
			CustomerCode: todo.CustomerCode,
			PhoneNumber:  todo.MerchantPhone,
		}
		merchantDetail, err := database.TodolistNewMerchantDetail(params)
		if err == nil {
			todolist.MerchantNewRecruitmentID = merchantDetail.ID
		}
	}

	tasksAttributes := []models.TaskDetail{}
	var todolistSer models.TodolistDetail
	var tasks []dbmodels.Task

	for _, taskAttribute := range todo.TasksAttributes {
		var fileEdukasi string
		var taskSer models.TaskDetail
		var subCategory dbmodels.TodolistSubCategory

		if taskAttribute.FileEdukasi != "" {
			fileEdukasi = utils.UploadFileEdukasi(taskAttribute.FileEdukasi, todolist.Mid)
		}

		subCategoryID, _ := strconv.Atoi(taskAttribute.SubCategoryID)
		task := dbmodels.Task{
			SupplierName:          taskAttribute.SupplierName,
			TodolistSubCategoryID: uint(subCategoryID),
			FileEdukasi:           fileEdukasi,
		}

		subErr := Dbcon.Where("id = ?", subCategoryID).Find(&subCategory).Error
		if subErr != nil {
			log.Println("Failed to get sub category : ", subErr)
			return todolistSer, subErr
		}
		if subCategory.Code == "CS04" && taskAttribute.SupplierName == "" {
			log.Println("Failed to get create todolist ", errors.New("Supplier Name must be exist for sub category "+subCategory.Name))
			return todolistSer, errors.New("Supplier Name must be exist for sub category " + subCategory.Name)
		}

		tasks = append(tasks, task)
		taskSer = SerializedTask(task, subCategory)
		tasksAttributes = append(tasksAttributes, taskSer)
	}

	err := Dbcon.Save(&todolist).Error
	if err != nil {
		log.Println("Failed connect to database SFA create data: ", err)
		return todolistSer, err
	}

	for i, task := range tasks {
		task.TodolistID = todolist.ID
		err := Dbcon.Create(&task).Error
		if err != nil {
			log.Println("Failed to create Task : ", err)
			return todolistSer, err
		}
		tasksAttributes[i].TodolistID = todolist.ID
		tasksAttributes[i].ID = task.ID
		tasksAttributes[i].CreatedAt = task.CreatedAt
		tasksAttributes[i].UpdatedAt = task.UpdatedAt
	}

	todolistSer = SerializedTodolist(todolist, tasksAttributes, todo.CustomerCode, todo.MerchantPhone)

	return todolistSer, nil
}

// GetMerchantDetail ..
func GetMerchantDetail(merchantID string) models.MerchantDetail {
	var res models.MerchantDetail

	sql := Dbcon.Raw("select a.id, a.name as merchant_name, a.merchant_id, a.owner_name, c.name as sub_area, a.address, a.note, a.image_merchant as merchant_image, a.phone_number as merchant_phone"+
		" from merchants a"+
		" left join public.sub_areas_villages b on a.village_id = b.village_id"+
		" left join public.sub_areas c on b.sub_area_id = c.id"+
		" where merchant_id = ?", merchantID).Scan(&res)

	if sql.Error != nil {
		log.Println("Failed to get admin detail: ", sql.Error)
	}
	return res
}

// TodolistNewMerchantList ..
func (database *DbPostgres) TodolistNewMerchantList(keyword string) ([]models.NewMerchantList, error) {
	var res []models.NewMerchantList

	keywordValue := "%" + keyword + "%"
	sql := Dbcon.Raw("select * from merchant_new_recruitments where name LIKE ? OR phone_number LIKE ? OR customer_code LIKE ? limit 5", keywordValue, keywordValue, keywordValue).Scan(&res)

	if sql.Error != nil {
		log.Println("Failed to get merchant list: ", sql.Error)
		return res, sql.Error
	}
	return res, nil
}

// GetNewMerchantDetailByID ..
func GetNewMerchantDetailByID(id uint) (models.NewMerchantDetail, error) {
	fmt.Println(">>> Todolist - GetNewMerchantDetailByID - Postgres <<<")
	var res models.NewMerchantDetail

	defQuery := "select id, name as merchant_name, customer_code, owner_name, sub_area_channel_name as sub_area, address, " +
		"phone_number as merchant_phone, longitude, latitude, village_id " +
		" from merchant_new_recruitments where id = ? "
	err := Dbcon.Raw(defQuery, id).Scan(&res).Error

	if err != nil {
		log.Println("Failed to get New Merchant detail: ", err)
		return res, err
	}

	res.MerchantID = strconv.Itoa(int(res.ID))
	return res, err
}

// TodolistUpdate ..
func (database *DbPostgres) TodolistUpdate(todo models.UpdateTodolist) (models.TodolistDetail, error) {
	fmt.Println(">>> Todolist - TodolistUpdate - Postgres <<<")

	// categoryID, _ := strconv.Atoi(todo.CategoryID)
	taskDate, _ := time.Parse("2006-01-02", todo.TaskDate)
	fmt.Println("task date before  ---------", todo.TaskDate)
	fmt.Println("task date   ---------", taskDate)
	fmt.Println("task date after ---------", taskDate.Format("02-01-2006"))

	var todolist dbmodels.Todolist
	var todolistSer models.TodolistDetail
	var category dbmodels.TodolistCategory
	todoErr := Dbcon.Where("id = ? AND status = 'Open'", todo.ID).Find(&todolist).Error

	if todoErr != nil {
		log.Println(" Todolist not found or Todolist has been done : ", todoErr)
		return todolistSer, todoErr
	}

	cErr := Dbcon.Where("id = ? ", todo.CategoryID).Find(&category).Error
	if cErr != nil {
		log.Println(" Todolist Category not found : ", cErr)
		return todolistSer, cErr
	}

	var merchErr error
	var dataMerchant models.DataMerchant
	if category.Code == "CR" {
		merchantID, _ := strconv.Atoi(todo.MerchantID)
		merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number FROM merchant_new_recruitments as m WHERE id = ?", merchantID).Scan(&dataMerchant).Error
		todolist.MerchantNewRecruitmentID = dataMerchant.ID
	} else {
		merchErr = Dbcon.Raw("SELECT m.id, m.name, m.merchant_id, m.village_id, m.phone_number FROM merchants as m WHERE merchant_id = ?", todo.MerchantID).Scan(&dataMerchant).Error
		todolist.Mid = dataMerchant.MerchantID
	}

	if merchErr != nil {
		log.Println(" Merchant not found : ", merchErr)
		return todolistSer, merchErr
	}

	todolist.MerchantName = dataMerchant.Name
	todolist.Status = "Open" // Open, Pending, Done, Not Exist
	todolist.TaskDate = taskDate
	todolist.TodolistCategoryID = category.ID
	todolist.VillageID = dataMerchant.VillageID
	todolist.Notes = todo.Notes
	todolist.SalesPhone = todo.SalesPhone

	err := Dbcon.Save(&todolist).Error
	tasksAttributes := []models.TaskDetail{}
	var taskVar []dbmodels.Task

	Dbcon.Raw("DELETE from tasks WHERE todolist_id = ?", todolist.ID).Scan(&taskVar)

	for _, taskAttribute := range todo.TasksAttributes {
		var fileEdukasi string
		if taskAttribute.FileEdukasi != "" {
			fileEdukasi = utils.UploadFileEdukasi(taskAttribute.FileEdukasi, todolist.Mid)
		}

		subCategoryID, _ := strconv.Atoi(taskAttribute.SubCategoryID)
		task := dbmodels.Task{
			SupplierName:          taskAttribute.SupplierName,
			TodolistSubCategoryID: uint(subCategoryID),
			TodolistID:            todolist.ID,
			FileEdukasi:           fileEdukasi,
		}

		var subCategory dbmodels.TodolistSubCategory
		subErr := Dbcon.Where("id = ?", subCategoryID).Find(&subCategory).Error
		if subErr != nil {
			log.Println("Failed to get sub category : ", subErr)
		}

		err := Dbcon.Create(&task).Error
		if err != nil {
			log.Println("Failed to create Task : ", err)
		}

		var taskSer models.TaskDetail
		taskSer = SerializedTask(task, subCategory)

		tasksAttributes = append(tasksAttributes, taskSer)
	}

	todolistSer = SerializedTodolist(todolist, tasksAttributes, dataMerchant.CustomerCode, dataMerchant.PhoneNumber)
	todolistSer.TodolistCategoryName = category.Name

	if err != nil {
		log.Println("Failed to update Todolist : ", err)
		return todolistSer, err
	}

	return todolistSer, err
}
