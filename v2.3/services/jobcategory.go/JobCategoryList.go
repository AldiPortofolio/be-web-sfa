package jobcategory

import (
	"fmt"
	"math"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
)

// MerchantDetail ..
func (svc *ServiceJobCategories) FilterJobCategories(req models.ReqFilterJobCategories, res *models.ResponsePagination) {
	fmt.Println(">>> MerchantList - ServiceJobCategories <<<")

	data, total, err := svc.Database.FilterJobCategories(req)
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
		res.Data = gin.H{}
		return
	}

	meta := models.MetaPagination{
		CurrentPage: int64(req.Page),
		NextPage:    int64(req.Page + 1),
		PrevPage:    int64(req.Page - 1),
		TotalPages:  int64(math.Ceil(float64(total) / float64(25))),
		TotalCount:  int64(total),
	}

	res.Meta = meta
	res.Data = data
	res.ResponseCode = "200"
	res.Message = "success"
}
