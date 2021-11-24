package job_category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/jobcategory.go"
	"strconv"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Delete godoc
// @Summary Jobcategories - Delete
// @Description Jobcategories - Delete
// @ID Jobcategories - Delete
// @Tags Jobcategories
// @Router /ottosfa/v2.3/jobcategories/delete [delete]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.JobCategories true "Body"
// @Success 200 {object} models.Response{} "jobcategoriess - Delete EXAMPLE"
func Delete(ctx *gin.Context) {
	fmt.Println(">>> jobcategories - Update - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	id := ctx.Params.ByName("id")

	categoryID, _ := strconv.Atoi(id)

	jobcategory.InitiateServiceJobCategories(log).Delete(token, int64(categoryID), &res)

	resBytes, _ := json.Marshal(res)
	log.Info("jobcategories-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
