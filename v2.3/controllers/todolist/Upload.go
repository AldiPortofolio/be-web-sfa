package todolist

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	todolist "ottosfa-api-web/v2.3/services/todolist"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Upload ..
// Todolist - Upload godoc
// @Summary Todolist - Upload
// @Description Todolist - Upload
// @ID Todolist - Upload V2.3
// @Tags Todolist
// @Router /ottosfa/v2.3/todolist/bulk [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param file formData file true "Body with file csv"
// @Success 200 {object} models.Response{data=string} "Todolist - Upload EXAMPLE"
func Upload(ctx *gin.Context) {
	fmt.Println(">>> Todolist - Upload - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	file, err := ctx.FormFile("file")
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	fmt.Println("file  ======= ", file.Filename)
	fileContent, _ := file.Open()
	fileBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	todolist.InitiateServiceTodolist(log).Upload(token, fileBytes, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Upload Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
