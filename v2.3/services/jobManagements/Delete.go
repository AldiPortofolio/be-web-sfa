package jobmanagements

import (
	"fmt"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (svc *ServiceJobManagements) Delete(ID string, res *models.Response) {
	fmt.Println(">>> Draft - ServicJobManagement<<<")

	JobID, _ := strconv.Atoi(ID)

	err := svc.Database.DeleteJobManagement(int64(JobID))
	if err != nil {
		res.Meta = utils.GetMessageFailedError(401, err)
		res.Data = gin.H{}
		return
	}

	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true

}
