package merchantsNewRecruitment

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"ottosfa-api-web/v2/services/merchantNewRecruitment"

	"github.com/gin-gonic/gin"
	ottologger "ottodigital.id/library/logger/v2"
)

// Create ..
// Merchant New Recruitment - Create godoc
// @Summary Merchant New Recruitment - Create
// @Description Merchant New Recruitment - Create
// @ID Merchant New Recruitment - Create
// @Tags Merchant New Rec
// @Router /ottosfa/v2/merchant-new-recruitment/create [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body dbmodels.MerchantNewRecruitments true "Body"
// @Success 200 {object} models.Response{} "Merchant New Recruitment - Create EXAMPLE"
func Create(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - Create - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := dbmodels.MerchantNewRecruitments{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("MerchantNewRecruitment-Create  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).Create(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-Create Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
