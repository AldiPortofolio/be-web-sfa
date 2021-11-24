package postgres

import (
	"encoding/base64"
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

	"github.com/gocarina/gocsv"
	gormbulk "github.com/t-tiger/gorm-bulk-insert"
)

// TodolistUpload ..
func (database *DbPostgres) TodolistUpload(fileBytes []byte) (models.BulkLinkError, error) {
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
		var dataMerchant models.DataMerchant
		var dataCategory dbmodels.TodolistCategory
		var todolist models.Todolist
		var subCategories []dbmodels.TodolistSubCategory
		var errorMessages []string
		var rowError models.DataErrorByRow
		var sales dbmodels.Salesman

		merchantPhone := todo.MerchantPhoneNo
		merchantID := todo.MerchantID
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
			err = errors.New("Pastikan nomor HP yang di masukan sesuai")
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
				merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number FROM merchant_new_recruitments as m WHERE customer_code = ?", customerCode).Scan(&dataMerchant).Error
			}
			// else if merchantPhone != "" {
			// 	merchErr = Dbcon.Raw("SELECT m.id, m.name, m.customer_code, m.village_id, m.phone_number FROM merchant_new_recruitments as m WHERE phone_number = ?", merchantPhone).Scan(&dataMerchant).Error
			// }

			todolist.MerchantNewRecruitmentID = dataMerchant.ID
		} else {
			merchErr = Dbcon.Raw("SELECT m.id, m.name, m.merchant_id, m.village_id, m.phone_number FROM merchants as m WHERE merchant_id = ?", merchantID).Scan(&dataMerchant).Error
		}
		catErr := Dbcon.Raw("SELECT c.name, c.id FROM todolist_categories as c WHERE code = ?", category).Scan(&dataCategory).Error
		subCatErr := Dbcon.Raw("SELECT sc.name, sc.id, sc.code FROM todolist_sub_categories as sc WHERE sc.code IN (?)", subCategoryIds).Scan(&subCategories).Error

		if salesPhone != "" {
			salesErr := Dbcon.Where("phone_number = ?", salesPhone).Find(&sales).Error
			if salesErr != nil {
				errorMessages = append(errorMessages, "Sales with phone "+salesPhone+" not found")
			}
		}

		if merchErr != nil {
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

		fmt.Println("before task date ---------", taskDate)
		fmt.Println("task date ---------", newTaskDate)
		todolist.Mid = dataMerchant.MerchantID
		todolist.MerchantName = dataMerchant.Name
		todolist.TodolistCategoryID = dataCategory.ID
		todolist.VillageID = dataMerchant.VillageID
		todolist.Notes = notes
		todolist.Status = "Open"
		todolist.TaskDate = newTaskDate
		todolist.SalesPhone = salesPhone
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

// contains ..
func contains(subCategories []dbmodels.TodolistSubCategory, str string) bool {
	for _, sub := range subCategories {
		if sub.Code == str {
			return true
		}
	}
	return false
}

// CreateBulkTask ..
func CreateBulkTask(tasks []interface{}, gormchan chan error) {
	err := gormbulk.BulkInsert(Dbcon, tasks, 3000, "ActionDate")

	gormchan <- err
}

// generateUrlErrorFile ..
func generateUrlErrorFile(dataErrors []models.DataErrorByRow, todoCount string, dataErrCount string) string {
	csvContent, err := gocsv.MarshalString(&dataErrors) // Get all clients as CSV string
	// fmt.Println("Export ------ ", csvContent)
	if err != nil {
		panic(err)
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(csvContent))
	errorCSVContent := "data:text/csv;base64," + encodedString
	errorFileURL := utils.UploadFileError(errorCSVContent, "todolist")
	fmt.Println("=====>>> URL --->", errorFileURL)

	if errorFileURL != "" {
		bulkError := dbmodels.BulkErrorFile{
			ErrorFile: errorFileURL,
			BulkType:  "todolist",
			Message:   todoCount + " data berhasil diproses. " + dataErrCount + " data gagal diproses.",
		}

		bulkErr := Dbcon.Create(&bulkError).Error

		if bulkErr != nil {
			fmt.Println("Failed to create file error : ", bulkErr)
		}
	}

	return errorFileURL
}
