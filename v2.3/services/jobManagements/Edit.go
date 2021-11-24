package jobmanagements

import (
	"fmt"
	"ottosfa-api-web/constants"
	"strconv"

	// "ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/hosts/ottosfanotif"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"time"

	"github.com/opentracing-contrib/go-zap/log"
)

// Edit ..
func (svc *ServiceJobManagements) Edit(token string, req models.ReqEditJobManagement, res *models.Response) {
	fmt.Println(">>> Create - ServiceJobManagements <<<")

	id, err := svc.Database.CheckAdminToken(token)
	sender, _ := svc.Database.FindAdminById(int64(id))
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}
	var (
		acceptDate   time.Time
		deliverDate  time.Time
		completeDate time.Time
		resendDate   time.Time
		cancelDate   time.Time
		recipient    int64
		message      string
		notifCode    int
	)

	dataOld, _, _ := svc.Database.FilterJobManagementEdit(models.ReqFilterJobManagements{ID: req.Id, Page: 1, Limit: 10})

	if len(dataOld) == 0 {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		res.Meta.Message = "Data not Found"
		return
	}

	switch req.Status {
	case 1:
		deadline, _ := time.Parse("2006-01-02", req.Deadline)
		dataOld[0].Deadline = deadline
		dataOld[0].Name = req.Name
		dataOld[0].JobCategoryId = req.JobCategoryId
		categoryID := strconv.Itoa(int(req.JobCategoryId))
		dataCategory, err := svc.Database.DetailJobCategory(categoryID)
		if err != nil {
			res.Meta = utils.GetMetaResponse("job.categories.update.failed")
			res.Meta.Message = err.Error()
			return
		}
		dataOld[0].JobCategoryName = dataCategory
		dataOld[0].JobPriority = req.JobPriority
		dataOld[0].RecipientId = req.RecipientId
		if !req.StatusStorage {
			recipient = req.RecipientId
			message = "Anda Mendapatkan 1 Tugas Baru dari " + sender.FirstName + " " + sender.LastName
			notifCode = 0
		}
	case 2:
		acceptDate = time.Now()
		cancelDate = dataOld[0].CancelDate
		resendDate = dataOld[0].ResendDate
		completeDate = dataOld[0].CompleteDate
		deliverDate = dataOld[0].DeliverDate
	case 3:
		deliverDate = time.Now()
		acceptDate = dataOld[0].AcceptDate
		cancelDate = dataOld[0].CancelDate
		resendDate = dataOld[0].ResendDate
		completeDate = dataOld[0].CompleteDate
		recipient = req.SenderId
		message = sender.FirstName + " " + sender.LastName + " telah menjalankan 1 tugas"
		notifCode = 1
	case 4:
		completeDate = time.Now()
		acceptDate = dataOld[0].AcceptDate
		deliverDate = dataOld[0].DeliverDate
		cancelDate = dataOld[0].CancelDate
		resendDate = dataOld[0].ResendDate
		recipient = req.RecipientId
		message = "Selamat Tugas Anda telah disetujui"
		notifCode = 2
	case 5:
		cancelDate = time.Now()
		acceptDate = dataOld[0].AcceptDate
		deliverDate = dataOld[0].DeliverDate
		cancelDate = dataOld[0].CancelDate
		completeDate = dataOld[0].CompleteDate
	case 6:
		resendDate = time.Now()
		completeDate = dataOld[0].CompleteDate
		acceptDate = dataOld[0].AcceptDate
		deliverDate = dataOld[0].DeliverDate
		cancelDate = dataOld[0].CancelDate
		recipient = req.RecipientId
		message = "Anda mendapatkan 1 penugasan ulang dari " + sender.FirstName + " " + sender.LastName
		notifCode = 3
	}

	// if req.AcceptDate != "" {
	// 	acceptDate, _ = time.Parse("2006-01-02", req.AcceptDate)
	// }

	// if req.CompleteDate != "" {
	// 	completeDate, _ = time.Parse("2006-01-02", req.CompleteDate)
	// }

	// if req.DeliverDate != "" {
	// 	deliverDate, _ = time.Parse("2006-01-02", req.DeliverDate)
	// }

	// if req.ResendDate != "" {
	// 	resendDate, _ = time.Parse("2006-01-02", req.ResendDate)
	// }

	// if req.CancelDate != "" {
	// 	cancelDate, _ = time.Parse("2006-01-02", req.CancelDate)
	// }

	dataOld[0].Id = req.Id
	dataOld[0].Status = req.Status
	dataOld[0].JobDescriptions = req.JobDescriptions
	dataOld[0].UpdatedAt = time.Now()
	dataOld[0].AcceptDate = acceptDate
	dataOld[0].DeliverDate = deliverDate
	dataOld[0].CompleteDate = completeDate
	dataOld[0].ResendDate = resendDate
	dataOld[0].CancelDate = cancelDate
	dataOld[0].Reason = req.Reason
	dataOld[0].StatusStorage = req.StatusStorage

	_, err = svc.Database.SaveJobManagement(dataOld[0])
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseFailed)
		res.Meta.Message = err.Error()
		return
	}

	notif := ottosfanotif.ReqCreateNotif{
		SenderName:  sender.FirstName + " " + sender.LastName,
		ObjectId:    req.Id,
		ObjectType:  "JobManagement",
		RecipientId: recipient,
		SenderId:    int64(id),
		Message:     message,
		NotifCode:   notifCode,
	}

	if message != "" {
		err = ottosfanotif.SendNotif(&notif, token)
		if err != nil {
			log.Info("Error send notifikasi : ")
			fmt.Println(err)
		}
	}

	res.Meta = utils.GetMetaResponse(constants.KeyResponseSuccessful)

}
