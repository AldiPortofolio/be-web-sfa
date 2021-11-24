package utils

import (
	"encoding/base64"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"ottosfa-api-web/hosts/minio"
	"ottosfa-api-web/models"
	"ottosfa-api-web/models/miniomodels"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/gocarina/gocsv"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var (
	// RedisKeyAuth ..
	RedisKeyAuth string
)

const (
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

func init() {
	RedisKeyAuth = beego.AppConfig.DefaultString("redis.key.auth", "OTTO-SFA-TOKEN :")
}

// DecodeBearer ..
func DecodeBearer(str string) string {
	token := strings.Replace(str, "Bearer ", "", 1)
	return token
}

// Response ..
func Response(key string) models.Response {
	return models.Response{
		Meta: GetMetaResponse(key),
	}
}

// StatusAccount ..
func StatusAccount(Status int) string {
	var statusacc string

	switch Status {
	case 0:
		statusacc = "Unregistered"
		break
	case 1:
		statusacc = "Registered"
		break
	case 2:
		statusacc = "Verified"
		break
	case 3:
		statusacc = "Inactive"
		break
	case 4:
		statusacc = "Pending"
		break
	}

	return statusacc
}

// Rand64String ..
func Rand64String(n int) string {
	//todo bisa di pindahkan di global variable
	var src = rand.NewSource(time.Now().UnixNano())

	b := make([]byte, n)
	l := len(letterBytes)
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < l {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return string(b)
}

// HashPassword ..
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// UploadImage2Minio ..
func UploadImage2Minio(imageBase64 string, nameFile string, spanID string) (miniomodels.UploadRes, error) {
	random := rand.Intn(100000000)

	req := miniomodels.UploadReq{
		BucketName:  "ottodigital",
		Data:        imageBase64,
		NameFile:    nameFile + "-" + strconv.Itoa(random) + ".jpeg",
		ContentType: "image/jpeg",
	}

	res, errMinio := minio.Send(req, spanID)
	if errMinio != nil {
		fmt.Println("Failed to connect to minio:", errMinio)
		return res, errMinio
	}
	return res, nil
}

// UniqueString ..
func UniqueString(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

// ConvertDateFormat ..
func ConvertDateFormat(reqDate string) string {
	reqDateLayout := "2020-10-16"
	reqDate = "2020-10-16T00:00:00Z"

	date, _ := time.Parse(reqDateLayout, reqDate)

	return date.Format("2006-01-02")
}

// ConvertTime ..
func ConvertTime(t time.Time) string {
	return t.Format("2006-01-02")
}

// GetMessageFailedError ..
func GetMessageFailedError(code int, err error) models.Meta {
	return models.Meta{
		Status:  false,
		Code:    code,
		Message: err.Error(),
	}
}

// GetToken ..
func GetToken(r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	return reqToken
}

// UploadFileError ..
func UploadFileError(fileError string, bulkType string) string {
	random := rand.Intn(100000000)

	//upload Error File
	reqFile := miniomodels.UploadReq{
		BucketName:  "ottodigital",
		Data:        fileError,
		NameFile:    "errorFile_" + bulkType + "-" + strconv.Itoa(random) + ".csv",
		ContentType: "text/csv",
	}
	uid := fmt.Sprintf("%v", uuid.New())
	errorFile, err := minio.Send(reqFile, uid)
	if err != nil {
		fmt.Println("Gagal upload file Error bulk "+bulkType, err)
		return ""
	}

	fmt.Println("=====>>> URL ", errorFile.Url)
	return errorFile.Url
}

// UploadFileEdukasi ..
func UploadFileEdukasi(fileEdukasi string, merchName string) string {
	random := rand.Intn(100000000)

	//upload File edukasi
	reqEdu := miniomodels.UploadReq{
		BucketName:  "ottodigital",
		Data:        fileEdukasi,
		NameFile:    "Edukasi_" + merchName + "-" + strconv.Itoa(random) + ".pdf",
		ContentType: "application/pdf",
	}
	uid := fmt.Sprintf("%v", uuid.New())
	edukasi, err := minio.Send(reqEdu, uid)
	if err != nil {
		log.Println("Gagal upload file edukasi", err)
		return ""
	}

	fmt.Println("=====>>> URL ", edukasi.Url)
	return edukasi.Url
}

// StatusAttendance ..
func StatusAttendance(Status int) string {
	var statusAttendance string

	switch Status {
	case 0:
		statusAttendance = "-"
		break
	case 1:
		statusAttendance = "Validasi"
		break
	case 2:
		statusAttendance = "Diterima"
		break
	case 3:
		statusAttendance = "Ditolak"
		break
	default:
		statusAttendance = "-"
		break
	}

	return statusAttendance
}

// UploadFileJobUpload ..
func UploadFileJobUpload(fileUpload string) string {
	random := rand.Intn(100000000)

	//upload Error File
	reqFile := miniomodels.UploadReq{
		BucketName:  "ottodigital",
		Data:        fileUpload,
		NameFile:    "uploadJob_" + strconv.Itoa(random) + ".csv",
		ContentType: "text/csv",
	}
	uid := fmt.Sprintf("%v", uuid.New())
	errorFile, err := minio.Send(reqFile, uid)
	if err != nil {
		fmt.Println("Gagal upload file uploadJob ", err)
		return ""
	}

	fmt.Println("=====>>> URL ", errorFile.Url)
	return errorFile.Url
}

// UploadFileUploadJob ..
func UploadFileUploadJob(data []models.ReqBulkUploadJobManagement) string {
	csvContent, err := gocsv.MarshalString(&data) // Get all clients as CSV string

	fmt.Println("Export ------ ", csvContent)
	if err != nil {
		log.Println(err)
	}
	encodedString := base64.StdEncoding.EncodeToString([]byte(csvContent))
	bulkUploadJob := "data:text/csv;base64," + encodedString
	// errorFileURL := ""
	resultUploadFileURL := UploadFileJobUpload(bulkUploadJob)
	log.Println("=====>>> URL --->", resultUploadFileURL)
	return resultUploadFileURL
}
