package jobmanagements

import (
	"encoding/csv"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/database/dbmodels"
	"ottosfa-api-web/hosts/ottosfanotif"
	mail "ottosfa-api-web/mailers"
	"ottosfa-api-web/models"
	"ottosfa-api-web/utils"
	"strconv"
	"strings"
	"time"
	
	"github.com/gocarina/gocsv"
)

// Upload ..
func (svc *ServiceJobManagements) Upload(token string, fileBytes []byte, res *models.Response) {
	fmt.Println(">>> Upload - ServiceJobManagements <<<")

	userId , err := svc.Database.CheckAdminToken(token)
	if err != nil {
		res.Meta = utils.GetMetaResponse(constants.KeyResponseInvalidToken)
		return
	}

	go svc.BulkUploadJob(fileBytes, userId, token)


	res.Meta.Message = "Upload Tugas sedang di proses dan hasil upload akan dikirimkan melalui email."
	res.Meta.Status = true
	res.Meta.Code = 200

}

// BulkUploadJob ..
func (svc *ServiceJobManagements) BulkUploadJob (fileBytes []byte, idSender uint, token string) {
	fmt.Println(">>> BulkUploadJob - ServiceJobManagements <<<")
	// sugarLogger := database.General.OttoZaplog
	var JobUploadCsv []models.ReqBulkUploadJobManagement

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use pipe as delimiter
	})

	err := gocsv.UnmarshalBytes(fileBytes, &JobUploadCsv)
	if err != nil {
		log.Println("Failed to extract file : ", err)
		return
	}

	var dataresult []models.ReqBulkUploadJobManagement

	for i, jobUpload := range JobUploadCsv {
		fmt.Println("csv job ------>>", jobUpload)
		
		no := strconv.Itoa(i+1)
		recipients := strings.Split(jobUpload.RecipientEmail, "|")
		if len(recipients) == 1 {
			job, err := svc.PackMessageJob(jobUpload, recipients[0], idSender)
			check := svc.CheckSenderRecipient(int64(idSender),job.RecipientId)
			checkPriority := svc.CheckPriority(job.JobPriority)
			fmt.Println("hasil check : ", check, checkPriority)
			jobByte, _ := json.Marshal(job)
			fmt.Println(string(jobByte))
			if err != nil || !check || !checkPriority {
				var reason string
				if err != nil {
					reason = err.Error()
				} else if !check { 
					reason = "Anda tidak dapat memberikan tugas ke Atasan" 
				} else if !checkPriority {
					reason = `Prioritas harus "High", "Low", atau "Medium"`
				} 
				dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
					No: no,
					Name: jobUpload.Name,
					JobCategory: jobUpload.JobCategory,
					RecipientEmail: jobUpload.RecipientEmail,
					JobPriority: jobUpload.JobPriority,
					Deadline: jobUpload.Deadline,
					DeskripsiTugas: jobUpload.DeskripsiTugas,
					LabelAttachment: jobUpload.LabelAttachment,
					LinkAttachment: jobUpload.LinkAttachment,
					Keterangan: reason ,
				})
				continue
			}
			idJob, err := svc.Database.SaveJobManagement(job)
			if err != nil {
				dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
					No: no,
					Name: jobUpload.Name,
					JobCategory: jobUpload.JobCategory,
					RecipientEmail: jobUpload.RecipientEmail,
					JobPriority: jobUpload.JobPriority,
					Deadline: jobUpload.Deadline,
					DeskripsiTugas: jobUpload.DeskripsiTugas,
					LabelAttachment: jobUpload.LabelAttachment,
					LinkAttachment: jobUpload.LinkAttachment,
					Keterangan: err.Error(),
				})
				continue
			} else {
				dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
					No: no,
					Name: jobUpload.Name,
					JobCategory: jobUpload.JobCategory,
					RecipientEmail: jobUpload.RecipientEmail,
					JobPriority: jobUpload.JobPriority,
					Deadline: jobUpload.Deadline,
					DeskripsiTugas: jobUpload.DeskripsiTugas,
					LabelAttachment: jobUpload.LabelAttachment,
					LinkAttachment: jobUpload.LinkAttachment,
					Keterangan: "Sukses",
				})
			}
			
			sender, _ := svc.Database.FindAdminById(int64(idSender))
			notif := ottosfanotif.ReqCreateNotif{
				SenderName:  sender.FirstName + " " + sender.LastName,
				ObjectId:    idJob,
				ObjectType:  "JobManagement",
				RecipientId: job.RecipientId,
				SenderId:    job.SenderId,
				Message:     "Anda Mendapatkan 1 Tugas Baru dari " + sender.FirstName + " " + sender.LastName,
				NotifCode:   0,
			}
			err = ottosfanotif.SendNotif(&notif, token)
			if err != nil {
				log.Println("Error send notifikasi : ", err)
			}

		} else {
			for _ , recipient := range recipients {
				job, err := svc.PackMessageJob( jobUpload, recipient, idSender )
				check := svc.CheckSenderRecipient(int64(idSender),job.RecipientId)
				checkPriority := svc.CheckPriority(job.JobPriority)
				fmt.Println("hasil check : ", check, checkPriority)
				jobByte, _ := json.Marshal(job)
				fmt.Println(string(jobByte))
				if err != nil || !check || !checkPriority {
					var reason string
					if err != nil {
						reason = err.Error()
					} else if !check { 
						reason = "Anda tidak dapat memberikan tugas ke Atasan" 
					} else if !checkPriority {
						reason = `Prioritas harus "High", "Low", atau "Medium"`
					} 
					dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
						No: no,
						Name: jobUpload.Name,
						JobCategory: jobUpload.JobCategory,
						RecipientEmail: recipient,
						JobPriority: jobUpload.JobPriority,
						Deadline: jobUpload.Deadline,
						DeskripsiTugas: jobUpload.DeskripsiTugas,
						LabelAttachment: jobUpload.LabelAttachment,
						LinkAttachment: jobUpload.LinkAttachment,
						Keterangan: reason,
					})
					continue
				}
				idJob , err := svc.Database.SaveJobManagement(job)
				if err != nil {
					dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
						No: no,
						Name: jobUpload.Name,
						JobCategory: jobUpload.JobCategory,
						RecipientEmail: recipient,
						JobPriority: jobUpload.JobPriority,
						Deadline: jobUpload.Deadline,
						DeskripsiTugas: jobUpload.DeskripsiTugas,
						LabelAttachment: jobUpload.LabelAttachment,
						LinkAttachment: jobUpload.LinkAttachment,
						Keterangan: err.Error(),
					})
					continue
				} else {
					dataresult = append(dataresult, models.ReqBulkUploadJobManagement{
						No: no,
						Name: jobUpload.Name,
						JobCategory: jobUpload.JobCategory,
						RecipientEmail: recipient,
						JobPriority: jobUpload.JobPriority,
						Deadline: jobUpload.Deadline,
						DeskripsiTugas: jobUpload.DeskripsiTugas,
						LabelAttachment: jobUpload.LabelAttachment,
						LinkAttachment: jobUpload.LinkAttachment,
						Keterangan: "Sukses",
					})
				}
				sender, _ := svc.Database.FindAdminById(int64(idSender))

				notif := ottosfanotif.ReqCreateNotif{
					SenderName:  sender.FirstName + " " + sender.LastName,
					ObjectId:    idJob,
					ObjectType:  "JobManagement",
					RecipientId: job.RecipientId,
					SenderId:    job.SenderId,
					Message:     "Anda Mendapatkan 1 Tugas Baru dari " + sender.FirstName + " " + sender.LastName,
					NotifCode:   0,
				}
				err = ottosfanotif.SendNotif(&notif, token)
				if err != nil {
					log.Println("Error send notifikasi : ", err)
				}
			}
		}
	
	}

	url := utils.UploadFileUploadJob(dataresult)
	sender, _ := svc.Database.FindAdminById(int64(idSender))
	fmt.Println( "ini sender :",idSender, sender)
	
	mail.UploadJobMailer("",url, sender.Email)
	

}

func (svc *ServiceJobManagements) PackMessageJob (job models.ReqBulkUploadJobManagement, email string, senderId uint ) (dbmodels.JobManagements, error) {
	var res dbmodels.JobManagements
	var err error

	if job.Deadline == "" || job.DeskripsiTugas == "" || job.JobPriority == "" || job.Name == "" {
		err = errors.New("name, deadline, deskripsi, atau prioritas, tidak boleh kosong")
		return res, err
	}

	if !svc.CheckDateFormat(job.Deadline) {
		err = errors.New("tanggal akhir penugasan tidak boleh sebelum hari ini")
		return res, err
	}

	JobCategory, err := svc.Database.FindByName(job.JobCategory)
	if err != nil {
		return res, errors.New("job category tidak terdaftar")
	}

	recipient, err := svc.Database.FindAdminByEmail(email)
	if err != nil {
		return res, errors.New("user penerima tugas tidak terdaftar")
	}
	// var listLink []map[string]string
	// listLink = append(listLink, map[string]string{
	// 	"fileName": job.LabelAttachment,
	// 	"url":      job.LinkAttachment,
	// 	"size":     "",
	// })
	labels := strings.Split(job.LabelAttachment, "|")
	links := strings.Split(job.LinkAttachment, "|")
	alink := ""
	for i:=0 ; i < len(labels) ; i ++ {
		alink += `<a href="` + links[i] + `" target="_blank">`+labels[i]+`</a><br> `
	}
	description := dbmodels.JobDescriptions {
		Description: `<p>` + job.DeskripsiTugas +  `</p><br>` + alink ,
	}

	deadline, _ := time.Parse("02/01/2006", job.Deadline)

	res.AssignmentDate = time.Now()
	res.Deadline = deadline
	res.CreatedAt = time.Now()
	res.JobCategoryId = JobCategory.ID
	res.JobPriority = job.JobPriority
	res.Name = job.Name
	res.RecipientId = recipient.Id
	res.SenderId = int64(senderId)
	res.UpdatedAt = time.Now()
	res.Status = 1
	res.StatusStorage = false
	res.JobDescriptions = []dbmodels.JobDescriptions{description}

	return res, nil

}


func (svc *ServiceJobManagements) CheckDateFormat (date string ) bool {
	dateFormat , err := time.Parse("02/01/2006 15:04:05", date + " 23:59:59")

	if diff := time.Until(dateFormat).Hours(); diff < 0 {
		return false
	}

	if err != nil {
		return false
	}

	return true

}