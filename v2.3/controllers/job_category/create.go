package job_category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/jobcategory.go"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Create ..
// JobCategories - Create godoc
// @Summary JobCategories - Create 
// @Description JobCategories - Create 
// @ID JobCategories - Create 
// @Tags JobCategories
// @Router /ottosfa/v2.3/jobcategories/save [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.JobCategories true "Body"
// @Success 200 {object} models.Response{} "JobCategoriess - Create EXAMPLE"
func Create(ctx *gin.Context) {
	fmt.Println(">>> JobCategories - Create - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.JobCategories{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		res.Meta.Code = 400
		res.Meta.Message = fmt.Sprintf("Body request error: %v", err)
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	if req.Name == "" || req.Description == "" {
		res.Meta.Code = 400
		res.Meta.Message = "Body request error: name or description can not null"
		res.Meta.Status = false
		ctx.JSON(http.StatusOK, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Todolist-Create  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	jobcategory.InitiateServiceJobCategories(log).Create(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Create Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
