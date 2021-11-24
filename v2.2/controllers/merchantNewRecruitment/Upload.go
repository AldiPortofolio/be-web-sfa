package merchantsNewRecruitment

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2.2/services/merchantNewRecruitment"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Upload ..
// Merchant New Recruitment - Upload godoc
// @Summary Merchant New Recruitment - Upload
// @Description Merchant New Recruitment - Upload
// @ID Merchant New Recruitment - Upload v2.2
// @Tags Merchant New Rec v2.2
// @Router /ottosfa/v2.2/merchant-new-recruitment/bulk [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param file formData file true "Body with file csv"
// @Success 200 {object} models.Response{data=string} "Merchant New Recruitment - Upload EXAMPLE"
func Upload(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - Upload - Controller <<<")

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

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).Upload(token, fileBytes, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-Upload Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
