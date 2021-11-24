package jobmanagements

import (
	"fmt"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"

	"github.com/gin-gonic/gin"
)

func (svc *ServiceJobManagements) Detail(ID string, res *models.Response) {
	fmt.Println(">>> Draft - ServicJobManagement<<<")

	data, err := svc.Database.DetailJobManagement(ID)
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
		res.Data = gin.H{}
		return
	}

	res.Data = data
	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true

}
