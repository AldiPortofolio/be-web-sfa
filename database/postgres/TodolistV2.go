package postgres

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/gocarina/gocsv"
)

// TodolistCreateV2 ..
func (database *DbPostgres) TodolistCreateV2(todo models.CreateTodolistV2) (models.TodolistDetail, error) {
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
	todolist.MerchantPhone = todo.MerchantPhone
	todolist.Address = todo.Address
	todolist.AddressBenchmark = todo.AddressBenchmark
	todolist.OwnerName = todo.OwnerName
	todolist.MerchantID = todo.MerchantID
	todolist.SalesTypeID = todo.SalesTypeID

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

// TodolistUpdateV2 ..
func (database *DbPostgres) TodolistUpdateV2(todo models.UpdateTodolistV2) (models.TodolistDetail, error) {
	fmt.Println(">>> Todolist - TodolistUpdate - Postgres <<<", todo.MerchantID)

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
	var dataMerchant models.MerchantDetailV2
	if category.Code == "CR" {
		merchantID, _ := strconv.Atoi(todo.MerchantID)
		fmt.Println(">>> Todolist - Merchan ID - Postgres <<<", merchantID)
		merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number as merchant_phone FROM merchant_new_recruitments as m WHERE id = ?", merchantID).Scan(&dataMerchant).Error
		todolist.MerchantNewRecruitmentID = uint(dataMerchant.ID)
	} else {
		query := `select 
				m.id,
				m.store_name as merchant_name, 
				m.merchant_outlet_id as merchant_id, 
				m.alamat as address, 
				m.notes as note, 
				m.store_phone_number as merchant_phone, 
				m.kelurahan as village_id,
				o.owner_first_name as owner_name,
				m.patokan as address_benchmark,
				m.partner_customer_id as customer_code,
				m.sr_id as sales_type_id
				from merchant m
				left join "owner" o on o.id = m.owner_id 
				where m.merchant_outlet_id = ?`
		merchErr = DbConRose.Raw(query, todo.MerchantID).Scan(&dataMerchant).Error
		todolist.Mid = dataMerchant.MerchantID
	}

	if merchErr != nil {
		log.Println(" Merchant not found : ", merchErr)
		return todolistSer, merchErr
	}

	villageID, _ := strconv.Atoi(dataMerchant.VillageID)
	todolist.MerchantName = dataMerchant.MerchantName
	todolist.Status = "Open" // Open, Pending, Done, Not Exist
	todolist.TaskDate = taskDate
	todolist.TodolistCategoryID = category.ID
	todolist.VillageID = int64(villageID)
	todolist.Notes = todo.Notes
	todolist.SalesPhone = todo.SalesPhone
	todolist.MerchantPhone = dataMerchant.MerchantPhone
	todolist.Address = dataMerchant.Address
	todolist.AddressBenchmark = dataMerchant.AddressBenchmark
	todolist.OwnerName = dataMerchant.OwnerName
	todolist.MerchantID = dataMerchant.ID
	// todolist.MerchantID = todo.MerchantID
	todolist.SalesTypeID = dataMerchant.SalesTypeID

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

	todolistSer = SerializedTodolist(todolist, tasksAttributes, dataMerchant.CustomerCode, dataMerchant.MerchantPhone)
	todolistSer.TodolistCategoryName = category.Name

	if err != nil {
		log.Println("Failed to update Todolist : ", err)
		return todolistSer, err
	}

	return todolistSer, err
}

// TodolistUploadV2 ..
func (database *DbPostgres) TodolistUploadV2(fileBytes []byte) (models.BulkLinkError, error) {
	fmt.Println(">>> Todolist - Upload - Postgres <<<")
	// sugarLogger := database.General.OttoZaplog
	var todoCSV []models.BulkTodolist
	todoErrResponse := models.BulkLinkError{
		ErrorFile: "",
	}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})

	err := gocsv.UnmarshalBytes(fileBytes, &todoCSV)
	if err != nil {
		log.Println("Failed to extract file : ", err)
		return todoErrResponse, err
	}

	var todoSubCategories []interface{}
	var todolists []models.Todolist
	var suppliers []string
	var dataErrors []models.DataErrorByRow

	for i, todo := range todoCSV {
		fmt.Println("csv todo ------>>", todo)
		fmt.Println("csv category =------>>", todo.CategoryID)
		fmt.Println("csv merchant phone =------>>", todo.MerchantPhoneNo)
		fmt.Println("csv customer code =------>>", todo.CustomerCode)
		fmt.Println("csv Number Row Value =------->>", i)
		var dataMerchant models.MerchantDetailV2
		var dataCategory dbmodels.TodolistCategory
		var todolist models.Todolist
		var subCategories []dbmodels.TodolistSubCategory
		var errorMessages []string
		var rowError models.DataErrorByRow
		var sales dbmodels.Salesman

		merchantPhone := todo.MerchantPhoneNo
		merchantID := strings.ReplaceAll(todo.MerchantID, " ", "")
		customerCode := todo.CustomerCode
		category := todo.CategoryID
		subCategory := todo.SubCategoryID
		taskDate := todo.TaskDate
		salesPhone := todo.SalesPhone
		notes := todo.Notes
		supplier := todo.SupplierName
		subCategoryIds := strings.Split(subCategory, "$")

		if category == "" || subCategory == "" || taskDate == "" {
			err = errors.New("category, sub category, atau taskdate tidak boleh kosong")
		}

		if category == "CR" && customerCode == "" {
			err = errors.New("customer code untuk category CR tidak boleh kosong")
		}

		if category != "CR" && merchantID == "" {
			err = errors.New("MID tidak boleh kosong")
		}

		if merchantPhone != "" && merchantPhone[0:1] != "0" {
			err = errors.New("Pastikan nomor HP yang di masukan berawalan nol")
		}

		if err != nil {
			data, _ := json.Marshal(todo)
			json.Unmarshal(data, &rowError)
			rowError.NoRow = i + 1
			rowError.ErrorMessages = err.Error()
			dataErrors = append(dataErrors, rowError)
			continue
		}

		var merchErr error
		if category == "CR" {
			if customerCode != "" {
				merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number as merchant_phone FROM merchant_new_recruitments as m WHERE customer_code = ?", customerCode).Scan(&dataMerchant).Error
			}
			// else if merchantPhone != "" {
			// 	merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number FROM merchant_new_recruitments as m WHERE phone_number = ?", merchantPhone).Scan(&dataMerchant).Error
			// }

			todolist.MerchantNewRecruitmentID = uint(dataMerchant.ID)
		} else {
			query := `select 
				m.id,
				m.store_name as merchant_name, 
				m.merchant_outlet_id as merchant_id, 
				m.alamat as address, 
				m.notes as note, 
				m.store_phone_number as merchant_phone, 
				m.kelurahan as village_id,
				m.sr_id as sales_type_id,
				CONCAT(o.owner_first_name, ' ', o.owner_last_name) as owner_name,
				m.patokan as address_benchmark
				from merchant m
				left join "owner" o on o.id = m.owner_id 
				where m.merchant_outlet_id = ?`
			merchErr = DbConRose.Raw(query, merchantID).Scan(&dataMerchant).Error
		}
		catErr := Dbcon.Raw("SELECT c.name, c.id FROM todolist_categories as c WHERE code = ?", category).Scan(&dataCategory).Error
		subCatErr := Dbcon.Raw("SELECT sc.name, sc.id, sc.code FROM todolist_sub_categories as sc WHERE sc.code IN (?)", subCategoryIds).Scan(&subCategories).Error

		if salesPhone != "" && salesPhone != "-" {
			salesErr := Dbcon.Where("phone_number = ?", salesPhone).Find(&sales).Error
			if salesErr != nil {
				errorMessages = append(errorMessages, "Sales with phone "+salesPhone+" not found")
			}
		}

		if merchErr != nil {
			logs.Error("Find Merchant rose : ", merchErr)
			errorMessages = append(errorMessages, "Merchant not found")
		}

		if dataMerchant == (models.MerchantDetailV2{}) {
			logs.Error("Find Merchant rose : ", merchErr)
			errorMessages = append(errorMessages, "Merchant not found")
		}

		if catErr != nil {
			errorMessages = append(errorMessages, "Category with code "+category+" not found")
		}

		if subCatErr != nil {
			errorMessages = append(errorMessages, "Failed to get sub categories --- "+strings.Join(subCategoryIds, ","))
		}

		for _, sub := range subCategoryIds {
			if !contains(subCategories, sub) {
				errorMessages = append(errorMessages, "Faild to get Sub Category with code "+sub)
			}
		}

		newTaskDate, _ := time.Parse("01/02/2006 15:04:05", taskDate+" 23:59:59")
		if newTaskDate.Format("01/02/2006") == "01/01/0001" {
			errorMessages = append(errorMessages, "Invalid Task Date format")
		}
		todayDate := time.Now().Format("01/02/2006")

		today, _ := time.Parse("01/02/2006", todayDate)
		if newTaskDate.Before(today) {
			errorMessages = append(errorMessages, "Task Date tidak boleh hari sebelum hari ini")
		}

		if len(errorMessages) > 0 {
			data, _ := json.Marshal(todo)
			json.Unmarshal(data, &rowError)
			rowError.NoRow = i + 1
			rowError.ErrorMessages = strings.Join(errorMessages, "|")
			dataErrors = append(dataErrors, rowError)
			continue
		}

		villageID, _ := strconv.Atoi(dataMerchant.VillageID)
		fmt.Println("before task date ---------", taskDate)
		fmt.Println("task date ---------", newTaskDate)
		todolist.Mid = dataMerchant.MerchantID
		todolist.MerchantName = dataMerchant.MerchantName
		todolist.TodolistCategoryID = dataCategory.ID
		todolist.VillageID = int64(villageID)
		todolist.Notes = notes
		todolist.Status = "Open"
		todolist.TaskDate = newTaskDate
		todolist.SalesPhone = salesPhone
		todolist.Address = dataMerchant.Address
		todolist.AddressBenchmark = dataMerchant.AddressBenchmark
		todolist.OwnerName = dataMerchant.OwnerName
		todolist.MerchantID = dataMerchant.ID
		todolist.MerchantPhone = dataMerchant.MerchantPhone
		todolist.SalesTypeID = dataMerchant.SalesTypeID

		// todolist.MerchantPhone = merchantPhone
		// todolist.CustomerCode = customerCode

		todoSubCategories = append(todoSubCategories, subCategories)
		todolists = append(todolists, todolist)
		suppliers = append(suppliers, supplier)
	}

	if len(todolists) > 0 {
		for i, todolist := range todolists {
			var tasks []interface{}

			todoErr := Dbcon.Create(&todolist).Error
			if todoErr != nil {
				log.Println("Failed to create todolist : ", todoErr)
				return todoErrResponse, todoErr
			}
			fmt.Println("====>>>>>", todolist.ID)

			tasksError := make(chan error)

			for _, sub := range todoSubCategories[i].([]dbmodels.TodolistSubCategory) {
				var task dbmodels.Task
				fmt.Println("Supply -=-=-=-=- ", suppliers[i])
				task.TodolistSubCategoryID = sub.ID
				if sub.ID == 15 {
					task.SupplierName = suppliers[i]
				}
				task.TodolistID = todolist.ID

				tasks = append(tasks, task)
			}

			go CreateBulkTask(tasks, tasksError)

			if taskErr := <-tasksError; taskErr != nil {
				log.Println("Failed connect to database : ", taskErr)
				return todoErrResponse, taskErr
			}
		}
	}

	if len(dataErrors) > 0 {
		dataErrCount := strconv.Itoa(len(dataErrors))
		todoCount := strconv.Itoa(len(todolists))
		todoErrResponse.ErrorFile = generateUrlErrorFile(dataErrors, todoCount, dataErrCount)
		return todoErrResponse, errors.New(todoCount + " data berhasil diproses. " + dataErrCount + " data gagal diproses.")
	}
	//else {
	return todoErrResponse, nil
	//}
}

// TodolistDetail ..
func (database *DbPostgres) TodolistDetailV2(todolistID string) (models.ShowTodolist, error) {
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
		merchantDetail := GetMerchantDetailFromTodolist(todolist)
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

// GetMerchantDetailV2 ..
func GetMerchantDetailV2(merchantID string) models.MerchantDetail {
	var res models.MerchantDetail

	query := `select 
			m.id,
			m.store_name as merchant_name, 
			m.merchant_outlet_id as merchant_id, 
			m.alamat as address, 
			m.notes as note, 
			m.store_phone_number as merchant_phone, 
			m.kelurahan as village_id,
			o.owner_first_name as owner_name,
			m.patokan as address_benchmark
			from merchant m
			left join "owner" o on o.id = m.owner_id 
			where m.merchant_outlet_id = ?
`

	sql := Dbcon.Raw(query, merchantID).Scan(&res)

	if sql.Error != nil {
		log.Println("Failed to get admin detail: ", sql.Error)
	}
	return res
}

// GetMerchantDetailFromTodolist ..
func GetMerchantDetailFromTodolist(todolist models.Todolist) models.MerchantDetail {
	merchantDetail := models.MerchantDetail{
		ID:            uint(todolist.MerchantID),
		MerchantName:  todolist.MerchantName,
		MerchantPhone: todolist.MerchantPhone,
		MerchantID:    todolist.Mid,
		OwnerName:     todolist.OwnerName,
		SubArea:       todolist.OwnerName,
		Address:       todolist.Address,
		Note:          todolist.Notes,
	}

	return merchantDetail
}
