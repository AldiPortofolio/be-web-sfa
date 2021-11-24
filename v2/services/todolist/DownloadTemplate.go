package todolist

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// DownloadTemplate ..
func (svc *ServiceTodolist) DownloadTemplate(token string, res *models.Response) {
	fmt.Println(">>> DownloadTemplate - ServiceTodolist <<<")

	_, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	data, err := svc.Database.TemplateTodolist()
	if err != nil {
		res.Meta = utils.GetMetaResponse("todolist.download.failed")
		res.Meta.Message = err.Error()
		return
	}

	res.Meta = utils.GetMetaResponse("todolist.download.success")
	res.Data = data
	res.Meta.Message = "Success"
	res.Meta.Status = true

	return
}
