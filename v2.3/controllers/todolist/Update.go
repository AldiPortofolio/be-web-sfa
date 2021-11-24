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

// Update godoc
// @Summary Todolist - Update V2.3
// @Description Todolist - Update V2.3
// @ID Todolist - Update V2.3
// @Tags Todolist V2.3
// @Router /ottosfa/v2.3/todolist/update [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.UpdateTodolistV2 true "Body"
// @Success 200 {object} models.Response{} "Todolists - Update EXAMPLE"
func Update(ctx *gin.Context) {
	fmt.Println(">>> Todolist - Update - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.UpdateTodolistV2{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("Todolist-Update  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	todolist.InitiateServiceTodolist(log).Update(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Update Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
