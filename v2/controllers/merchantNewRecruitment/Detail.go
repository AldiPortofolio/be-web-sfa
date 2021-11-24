package merchantsNewRecruitment

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	ottologger "ottodigital.id/library/logger/v2"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services/merchantNewRecruitment"
)

// Detail ..
// Merchant New Recruitment - Detail godoc
// @Summary Merchant New Recruitment - Detail
// @Description Merchant New Recruitment - Detail
// @ID Merchant New Recruitment - Detail
// @Tags Merchant New Rec
// @Router /ottosfa/v2/merchant-new-recruitment/detail/:merchant_new_rec_id [get]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Success 200 {object} models.Response{data=models.MerchantNewRecruitmentExportRes} "Merchant New Recruitment - Detail EXAMPLE"
func Detail(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - Detail - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)
	merchantNewRecID := ctx.Params.ByName("merchant_new_rec_id")

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	log.Info("MerchantNewRecruitment-Detail  Controller",
		log.AddField("RequestBody-ID:", merchantNewRecID))

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).Detail(token, merchantNewRecID, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-Detail Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
