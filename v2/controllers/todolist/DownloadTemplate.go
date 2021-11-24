package todolist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services/todolist"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// DownloadTemplate ..
// Todolist - Download Template godoc
// @Summary Todolist - Download Template
// @Description Todolist - Download Template
// @ID Todolist - Download Template
// @Tags Todolist
// @Router /ottosfa/v2/todolist/download-template [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{data=string} "Todolist - Download Template EXAMPLE"
func DownloadTemplate(ctx *gin.Context) {
	fmt.Println(">>> Todolist - Download Template - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	todolist.InitiateServiceTodolist(log).DownloadTemplate(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-DownloadTemplate Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
