package jobmanagements

import (
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
)

// Create ..
func (svc *ServiceJobManagements) RecipientList(token string,  req models.RecipientReq, res *models.Response) {
	fmt.Println(">>> Create - ServiceJobManagements <<<")

	adminID, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	admin, _ := svc.Database.FindAdminById(int64(adminID))

	var assigmentRoles []string
	switch admin.AssignmentRole {
    case "hq":
       assigmentRoles = []string{"rsm", "bsm", "tl"}
    case "rsm":
        assigmentRoles = []string{"rsm", "bsm", "tl"}
    case "bsm":
        assigmentRoles = []string{"bsm", "tl"}
	default:
		assigmentRoles = []string{"rsm", "bsm", "tl"}
	}

	recipients, _ := svc.Database.FindAdminByAssignmentRole(assigmentRoles, req)

	res.Data = recipients
	res.Meta.Code = 200
	res.Meta.Message = "success"
	res.Meta.Status = true
}
