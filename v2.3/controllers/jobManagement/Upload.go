package jobManagement

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	jobManagement "ottosfa-api-web/v2.3/services/jobManagements"
	"strings"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Upload ..
// JobManagement - Upload godoc
// @Summary JobManagement - Upload
// @Description JobManagement - Upload
// @ID JobManagement - Upload
// @Tags JobManagement
// @Router /ottosfa/v2.3/job-management/upload [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param file formData file true "Body with file csv"
// @Success 200 {object} models.Response{} "JobManagement - Upload EXAMPLE"
func Upload(ctx *gin.Context) {
	fmt.Println(">>> JobManagement - Upload - Controller <<<")

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
	splitName := strings.Split(file.Filename, ".")
	if splitName[len(splitName)-1] != "csv" {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", errors.New("hanya diperbolehkan upload file csv")))
		return
	}
	fileContent, _ := file.Open()
	fileBytes, err := ioutil.ReadAll(fileContent)
	if err != nil {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("get form err: %s", err.Error()))
		return
	}

	jobManagement.InitiateServiceJobManagements(log).Upload(token, fileBytes, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("Todolist-Upload Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}