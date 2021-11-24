package activitasSalesmen

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.3/services/activitasSalesmen"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// List ..
// Activitas Salesmen - List Detail Todo List godoc
// @Summary Activitas Salesmen - List Detail Todo List
// @Description Activitas Salesmen - List Detail Todo List
// @ID Activitas Salesmen - List Detail Todo List
// @Tags Activitas Salesmen V2.3
// @Router /ottosfa/v2.3/activitas-salesmen/lits-detail-todolist [POST]
// @Accept json
// @Produce json
// @Param Body body dbmodels.DetailListActivitasTodolistReq{} true "Body"
// @Success 200 {object} models.Response{} "Activitas Salesmen - List Detail Todo List"
func ListDetailActivitasSalesmenTodoList(ctx *gin.Context) {
	fmt.Println(">>> Activitas Salesmen - list - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	res := models.ResponsePagination{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.DetailListActivitasTodolistReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	activitasSalesmen.InitiateServicActivitasSalesmen(log).ListDetailActivitasSalesmenTodolist(req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Activitas Salesmen-List Controller",
		log.AddField("ResponseBody", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
