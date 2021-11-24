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

// List ..
// Merchant New Recruitment - List godoc
// @Summary Merchant New Recruitment - List
// @Description Merchant New Recruitment - List
// @ID Merchant New Recruitment - List
// @Tags Merchant New Rec
// @Router /ottosfa/v2/merchant-new-recruitment/list [post]
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer 7n1cdTNMK_EeoXr_0D0luYJ68NaYESWxyPXd"
// @Param Body body models.MerchantNewRecruitmentListReq true "Body"
// @Success 200 {object} models.ResponsePagination{data=[]dbmodels.MerchantNewRecruitments} "Merchant New Recruitment - List Response EXAMPLE"
func List(ctx *gin.Context) {
	fmt.Println(">>> Merchant New Recruitment - List - Controller <<<")

	log := ottologger.InitLogs(ctx.Request)

	page := ctx.DefaultQuery("page", "1")
	token := utils.GetToken(ctx.Request)

	res := models.ResponsePagination{
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
	log.Info("MerchantNewRecruitment-List  Controller",
		log.AddField("RequestBody:", string(reqBytes)))

	merchantNewRecruitment.InitiateServiceMerchantNewRecruitment(log).List(token, req, &res)

	resBytes, _ := json.Marshal(res)
	log.Info("MerchantNewRecruitment-List Controller",
		log.AddField("ResponseBody: ", string(resBytes)))

	ctx.JSON(http.StatusOK, res)
}
