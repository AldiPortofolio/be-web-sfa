package postgres

import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"strconv"
	"strings"
	ottoutils "ottodigital.id/library/utils"
)

// SerializedTaskFollowUp ..
func SerializedTaskFollowUp(task dbmodels.Task, followUps []dbmodels.FollowUp, subCategory dbmodels.TodolistSubCategory) models.TaskDetail {
	var taskSer models.TaskDetail
	t, _ := json.Marshal(task)
	json.Unmarshal(t, &taskSer)
	taskSer.SubCategoryName = subCategory.Name
	taskSer.FollowUps = followUps

	return taskSer
}

// SerializedTask ..
func SerializedTask(task dbmodels.Task, subCategory dbmodels.TodolistSubCategory) models.TaskDetail {
	var taskSer models.TaskDetail
	t, _ := json.Marshal(task)
	json.Unmarshal(t, &taskSer)
	taskSer.SubCategoryName = subCategory.Name

	return taskSer
}

// SerializedTodolist ..
func SerializedTodolist(todolist dbmodels.Todolist, tasksAttributes []models.TaskDetail, customerCode, merchantPhone string) models.TodolistDetail {
	var todolistSer models.TodolistDetail

	todo, _ := json.Marshal(todolist)
	json.Unmarshal(todo, &todolistSer)

	todolistSer.TaskDate = todolist.TaskDate.Format("02-01-2006")
	todolistSer.Tasks = tasksAttributes
	todolistSer.CustomerCode = customerCode
	todolistSer.MerchantPhone = merchantPhone

	return todolistSer
}

// SerializedTodolistDetail ..
func SerializedTodolistDetail(todolist models.Todolist, merchant interface{}, tasks []models.TaskDetail, histories []dbmodels.TodolistHistory, salesmen []dbmodels.Salesman, sales dbmodels.Salesman) models.ShowTodolist {
	var todolistSer models.ShowTodolist
	var historiesSer []models.History

	todo, _ := json.Marshal(todolist)
	json.Unmarshal(todo, &todolistSer)
	todolistSer.Tasks = tasks
	todolistSer.MerchantDetail = merchant

	for _, history := range histories {
		todoHistory := models.History{
			ID:           history.ID,
			Description:  history.Description,
			FotoLocation: history.FotoLocation,
			NewTaskDate:  history.NewTaskDate,
			OldTaskDate:  history.OldTaskDate,
			Status:       history.Status,
			TodolistID:   history.TodolistID,
			Longitude:    history.Longitude,
			Latitude:     history.Latitude,
			Tasks:        GetTasks(history.CreatedAt.Format("01/02/2006"), tasks),
		}

		historiesSer = append(historiesSer, todoHistory)
	}

	todolistSer.Histories = historiesSer
	if todolist.SalesPhone == "" {
		for _, s := range salesmen {
			salesman := []string{s.FirstName, s.LastName}
			name := strings.Join(salesman, " ")
			salesData := []string{name, s.PhoneNumber}
			salesSer := models.SalesList{
				ID:         s.ID,
				LabelSales: strings.Join(salesData, " - "),
			}

			todolistSer.PossibleSales = append(todolistSer.PossibleSales, salesSer)
		}
	} else {
		salesman := []string{sales.FirstName, sales.LastName}
		name := strings.Join(salesman, " ")
		salesData := []string{name, sales.PhoneNumber}
		salesSer := models.SalesList{
			ID:         sales.ID,
			LabelSales: strings.Join(salesData, " - "),
		}
		todolistSer.PossibleSales = append(todolistSer.PossibleSales, salesSer)
	}
	return todolistSer
}

// GetTasks ..
func GetTasks(createdAt string, tasks []models.TaskDetail) []models.TaskDetail {
	var historyTasks []models.TaskDetail
	for _, task := range tasks {
		if task.ActionDate.Format("01/02/2006") == createdAt {
			historyTasks = append(historyTasks, task)
		}
	}
	return historyTasks
}

// SerializedSales ..
func SerializedSales(sales dbmodels.Salesman) models.SalesmanResponse {
	var salesSer models.SalesmanResponse
	imageHost := ottoutils.GetEnv("SFAADMIN_RUBY_IMAGE_URL", "https://adminsfa.ottopay.id/")

	data, _ := json.Marshal(sales)
	_ = json.Unmarshal(data, &salesSer)
	salesSer.Gender = dbmodels.GetGender(sales.Gender)
	salesSer.Status = dbmodels.GetStatus(sales.Status)
	salesSer.Positions = SerializePositions(sales.Positions)
	salesSer.CreatedAt = sales.CreatedAt
	salesSer.UpdatedAt = sales.UpdatedAt
	salesTypes := getSalesType(sales.SalesTypeID)
	salesSer.SalesType = salesTypes.Name

	if strings.Contains(sales.IDCard, "http") {
		salesSer.IDCardPic = sales.IDCard
	} else if sales.IDCard != "" {
		salesSer.IDCardPic = imageHost + "uploads/salesman/id_card/" + strconv.Itoa(int(sales.ID)) + "/" + sales.IDCard
	}

	if strings.Contains(sales.Photo, "http") {
		salesSer.Photo = sales.Photo
	} else if sales.Photo != "" {
		salesSer.Photo = imageHost + "uploads/salesman/photo/" + strconv.Itoa(int(sales.ID)) + "/" + sales.Photo
	}

	if sales.PhoneNumber[0:1] == "0" {
		salesSer.PhoneNumber = sales.PhoneNumber
	} else {
		salesSer.PhoneNumber = "0" + sales.PhoneNumber
	}

	salesSer.Dob = sales.Dob.Format("02-01-2006")
	salesSer.WorkDate = sales.WorkDate.Format("02-01-2006")

	return salesSer
}

// SerializePositions ..
func SerializePositions(positions []dbmodels.Position) []models.PositionsResponse {
	var positionsSer []models.PositionsResponse
	for _, position := range positions {
		var region dbmodels.Region
		var branch dbmodels.Branch
		var area dbmodels.Area
		var subArea dbmodels.SubArea

		postSer := models.PositionsResponse{
			ID:     position.ID,
			RoleID: position.SalesRoleID,
			Role:   position.RoleName,
		}

		switch position.RegionableType {
		case "Region":
			region = getRegionByID(position.RegionableID)
		case "Branch":
			branch = getBranchByID(position.RegionableID)
			region = getRegionByID(branch.RegionID)
		case "Area":
			area = getAreaByID(position.RegionableID)
			branch = getBranchByID(area.BranchID)
			region = getRegionByID(branch.RegionID)
		case "SubArea":
			subArea = getSubAreaByID(position.RegionableID)
			area = getAreaByID(subArea.AreaID)
			branch = getBranchByID(area.BranchID)
			region = getRegionByID(branch.RegionID)
		}

		postSer.RegionID = region.ID
		postSer.Region = region.Code + " - " + region.Name
		postSer.BranchID = branch.ID
		postSer.Branch = branch.Code + " - " + branch.Name
		postSer.BranchOffice = branch.BranchOffice
		postSer.AreaID = area.ID
		postSer.Area = area.Code + " - " + area.Name
		postSer.SubAreaID = subArea.ID
		postSer.SubArea = subArea.Code + " - " + subArea.Name
		positionsSer = append(positionsSer, postSer)
	}

	return positionsSer
}

// getRegionByID ..
func getRegionByID(regionID uint) dbmodels.Region {
	var region dbmodels.Region
	regErr := Dbcon.Raw("select id, name, code from regions where id = ?", regionID).Scan(&region).Error
	if regErr != nil {
		logs.Error("Failed to get region ---- serializer position ", regErr)
	}
	return region
}

// getBranchByID ..
func getBranchByID(branchID uint) dbmodels.Branch {
	var branch dbmodels.Branch
	brErr := Dbcon.Raw("select id, name, code, branch_office, region_id from branches where id = ?", branchID).Scan(&branch).Error
	if brErr != nil {
		logs.Error("Failed to get branch ---- serializer position ", brErr)
	}
	return branch
}

// getAreaByID ..
func getAreaByID(areaID uint) dbmodels.Area {
	var area dbmodels.Area
	areaErr := Dbcon.Raw("select id, name, code, branch_id from areas where id = ?", areaID).Scan(&area).Error
	if areaErr != nil {
		logs.Error("Failed to get area ---- serializer position ", areaErr)
	}
	return area
}

// getSubAreaByID ..
func getSubAreaByID(subAreaID uint) dbmodels.SubArea {
	var subArea dbmodels.SubArea
	subErr := Dbcon.Raw("select id, name, code, area_id from sub_areas where id = ?", subAreaID).Scan(&subArea).Error
	if subErr != nil {
		logs.Error("Failed to get sub area ---- serializer position ", subErr)
	}
	return subArea
}

// getSalesType ..
func getSalesType(salesTypeID uint) dbmodels.SalesType {
	var salesType dbmodels.SalesType
	salesTypeErr := Dbcon.Raw("select id, name, description from sales_types where id = ?", salesTypeID).Scan(&salesType).Error
	if salesTypeErr != nil {
		logs.Error("Failed to get sales type ---- serializer sales", salesTypeErr)
	}
	return salesType
}

