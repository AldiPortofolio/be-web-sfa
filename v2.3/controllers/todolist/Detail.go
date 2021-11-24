package todolist

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	todolist "ottosfa-api-web/v2.3/services/todolist"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Detail ..
// Todolist - Detail godoc
// @Summary Todolist - Detail
// @Description Todolist - Detail
// @ID Todolist - Detail V2.3
// @Tags Todolist V2.3
// @Router /ottosfa/v2.3/todolist/detail/:todolist_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{} "Todolists - Create EXAMPLE"
func Detail(ctx *gin.Context) {
	fmt.Println(">>> Todolist - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	todolistID := ctx.Params.ByName("todolist_id")

	reqBytes, _ := json.Marshal(todolistID)
	log.Info("Todolist-Detail  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	todolist.InitiateServiceTodolist(log).Detail(token, todolistID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Detail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
