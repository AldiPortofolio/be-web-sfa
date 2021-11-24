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
	"strconv"
)

// Export ..
// Merchant New Recruitment - Export godoc
// @Summary Merchant New Recruitment - Export
// @Description Merchant New Recruitment - Export
// @ID Merchant New Recruitment - Export
// @Tags Merchant New Rec
// @Router /ottosfa/v2/merchant-new-recruitment/export [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.MerchantNewRecruitmentListReq true "Body"
// @Success 200 {object} models.Response{data=string} "Merchant New Recruitment - Export EXAMPLE"
func Export(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - Export - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	page := ctx.DefaultQuery("page", "1")
	token := utils.GetToken(ctx.Request)

	res := models.Response{
		Meta: utils.GetMetaResponse(constants.KeyResponseDefault),
	}

	req := models.MerchantNewRecruitmentListReq{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		go log.Error(fmt.Sprintf("Body request error: %v", err))
		ctx.JSON(http.StatusBadRequest, res)
		return
	}

	if req.Page == 0 {
		requestPage, _ := strconv.Atoi(page)
		req.Page = int64(requestPage)
	}

	reqBytes, _ := json.Marshal(req)
	log.Info("MerchantNewRecruitment-Export  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).Export(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-Export Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
