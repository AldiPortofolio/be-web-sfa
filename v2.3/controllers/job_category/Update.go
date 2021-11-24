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

// Update godoc
// @Summary Jobcategories - Update
// @Description Jobcategories - Update
// @ID Jobcategories - Update
// @Tags Jobcategories
// @Router /ottosfa/v2.3/jobcategories/update [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.JobCategories true "Body"
// @Success 200 {object} models.Response{} "jobcategoriess - Update EXAMPLE"
func Update(ctx *gin.Context) {
	fmt.Println(">>> jobcategories - Update - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.JobCategories{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("jobcategories-Update  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	jobcategory.InitiateServiceJobCategories(log).Update(token,req,&res)

	resBytes, _ := json.Marshal(res)
	log.Info("jobcategories-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
