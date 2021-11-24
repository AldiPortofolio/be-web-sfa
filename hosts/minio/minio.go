package minio

import (
	"encoding/json"
	"fmt"
	"net/http"
	ottoutils "ottodigital.id/library/utils"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"
	httputils "ottosfa-api-web/utils/http"
)

type MinioHost struct {
	Authorization string
	DeviceID      string
	models.GeneralModel
}

func InitMinioHost()  *MinioHost{
	return &MinioHost{
		Authorization:"",
		DeviceID:"",
	}
}

var (
	host               string
	endpointUpload	   string
)

func init() {
	host = ottoutils.GetEnv("MINIO_HOST", "http://13.228.25.85:8312")
	endpointUpload = ottoutils.GetEnv("MINIO_ENDPOINT_UPLOAD", "/upload")
}

// Send ..
func (head *MinioHost) Send(msgReq interface{}, typeTrans string) ([]byte, error) {

	header := make(http.Header)
	header.Add("Accept", "*/*")
	header.Add("Content-Type", "application/json")

	fmt.Println("header:", header)

	urlSvr := ""
	method := constants.HttpMethodPost
	switch typeTrans {

	case constants.MinioUpload:
		fmt.Println(constants.MinioUpload)
		urlSvr = host + endpointUpload
		method = constants.HttpMethodPost
		header.Add("Content-Type", "application/json")
		break

	}

	dataReq, _ := json.Marshal(msgReq)

	fmt.Println("req --> ", string(dataReq))

	data, err := httputils.SendHttpRequest(method, urlSvr, header, msgReq)
	fmt.Println()

	fmt.Println("res -->", string(data))

	return data, err
}