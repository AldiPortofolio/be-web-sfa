package merchantsNewRecruitment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services/merchantNewRecruitment"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// DownloadTemplate ..
// Merchant New Recruitment - Download Template godoc
// @Summary Merchant New Recruitment - Download Template
// @Description Merchant New Recruitment - Download Template
// @ID Merchant New Recruitment - Download Template
// @Tags Merchant New Rec
// @Router /ottosfa/v2/merchant-new-recruitment/download-template [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{data=string} "Merchant New Recruitment - Download Template EXAMPLE"
func DownloadTemplate(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - Download Template - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).DownloadTemplate(token, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-DownloadTemplate Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
