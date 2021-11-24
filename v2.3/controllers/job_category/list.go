package job_category

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/jobcategory.go"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// JobCategories - list godoc
// @Summary JobCategories - Filter
// @Description JobCategories - Filter
// @ID JobCategories - Filter
// @Tags JobCategories
// @Router /ottosfa/v2.3/jobcategories/filter [POST]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.ReqFilterJobCategories true "Body"
// @Success 200 {object} models.ResponsePagination{}
func Filter(ctx *gin.Context) {
	fmt.Println(">>> JobCategories - Filter - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.ResponsePagination{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.ReqFilterJobCategories{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	jobcategory.InitiateServiceJobCategories(log).FilterJobCategories(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("JobCategories-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
