package jobmanagements

import (
	"encoding/json"
	"fmt"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	ottosfanotif "ottosfa-api-web/hosts/ottosfanotif"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"time"

	"github.com/opentracing-contrib/go-zap/log"
)

// Create ..
func (svc *ServiceJobManagements) Create(token string, req models.ReqCreateJobManagement, res *models.Response) {
	fmt.Println(">>> Create - ServiceJobManagements <<<")

	senderId, err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}
	var (
		// createdAt time.Time
		deadline       time.Time
		assignmentDate time.Time
		// acceptDate time.Time
		// deliverDate time.Time
		// completeDate time.Time
		// resendDate time.Time
		// cancelDate time.Time

	)

	assignmentDate = time.Now()

	if req.Deadline != "" {
		deadline, _ = time.Parse("2006-01-02", req.Deadline)
	} else {
		deadline = time.Now()
	}

	for i := 0; i < len(req.RecipientId); i++ {
		check := svc.CheckSenderRecipient(int64(senderId), req.RecipientId[i])
		if !check {
			res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
			res.Meta.Message = "Anda tidak dapat memberikan tugas ke Atasan"
			return
		}
		var data dbmodels.JobManagements
		// data.Id = req.Id
		data.Name = req.Name
		data.JobCategoryId = req.JobCategoryId
		data.SenderId = int64(senderId)
		data.RecipientId = req.RecipientId[i]
		data.JobPriority = req.JobPriority
		if req.Status != 0 {
			data.Status = req.Status
		} else {
			data.Status = 1
		}
		data.StatusStorage = req.StatusStorage
		data.JobDescriptions = req.JobDescriptions
		data.CreatedAt = time.Now()
		data.UpdatedAt = time.Now()
		// data.AcceptDate = acceptDate
		// data.DeliverDate = deliverDate
		data.AssignmentDate = assignmentDate
		// data.CompleteDate = completeDate
		// data.ResendDate = resendDate
		// data.CancelDate = cancelDate
		data.Deadline = deadline

		dataByte, _ := json.Marshal(data)
		fmt.Println("data send to db : ", string(dataByte))
		if data.JobDescriptions[0].Id != 0 {
			data.JobDescriptions[0].Id = 0
		}
		id, err := svc.Database.SaveJobManagement(data)
		if err != nil {
			res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
			res.Meta.Message = err.Error()
			return
		}

		sender, _ := svc.Database.FindAdminById(int64(senderId))

		notif := ottosfanotif.ReqCreateNotif{
			SenderName:  sender.FirstName + " " + sender.LastName,
			ObjectId:    id,
			ObjectType:  "JobManagement",
			RecipientId: req.RecipientId[i],
			SenderId:    int64(senderId),
			Message:     "Anda Mendapatkan 1 Tugas Baru dari " + sender.FirstName + " " + sender.LastName,
			NotifCode:   0,
		}

		if !req.StatusStorage {
			err = ottosfanotif.SendNotif(&notif, token)
			if err != nil {
				log.Info("Error send notifikasi : ")
				fmt.Println(err)
			}
		}
	}
	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

}

func (svc *ServiceJobManagements) CheckSenderRecipient(senderId, recipientId int64) bool {
	sender, err := svc.Database.FindAdminById(senderId)
	if err != nil {
		return false
	}
	recipient, err := svc.Database.FindAdminById(recipientId)
	if err != nil {
		return false
	}
	fmt.Println("sender role : ", sender.AssignmentRole, "recipient role : ", recipient.AssignmentRole)
	// hq bsm rsm tl
	switch sender.AssignmentRole {
	case "rsm":
		if recipient.AssignmentRole == "hq" {
			return false
		}
	case "bsm":
		if recipient.AssignmentRole == "hq" || recipient.AssignmentRole == "rsm" {
			return false
		}
	case "tl":
		if recipient.AssignmentRole == "hq" || recipient.AssignmentRole == "bsm" || recipient.AssignmentRole == "rsm" {
			return false
		}
	}

	return true
}

// CheckPriority ..
func (svc *ServiceJobManagements) CheckPriority(priority string) bool {
	if priority == "High" || priority == "Low" || priority == "Medium" {
		return true
	}

	return false
}
