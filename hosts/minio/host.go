package minio

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	"ottosfa-api-web/models/miniomodels"
	https "ottosfa-api-web/utils/http"
)

// Env ..
type Env struct {
	Host                string `envconfig:"MINIO_HOST" default:"http://13.228.25.85:8312"`
	EndpointUploadImage string `envconfig:"MINIO_ENDPOINT_UPLOAD_IMAGE" default:"/upload"`
	HealthCheckKey      string `envconfig:"MINIO_HEALTH_CHECK_KEY" default:"OTTO-SFA-API_HEALTH_CHECK:MINIO"`
	Name                string `envconfig:"MINIO_NAME" default:"UPLOAD"`
}

var (
	minioEnv Env
)

func init() {
	err := envconfig.Process("MINIO", &minioEnv)
	if err != nil {
		fmt.Println("Failed to get MINIO env:", err)
	}
}

// Send ..
func Send(req miniomodels.UploadReq, spanID string) (miniomodels.UploadRes, error) {
	fmt.Println("<<< Upload Image to Minio >>>")
	sugarLogger := ottologger.GetLogger()

	var res miniomodels.UploadRes

	urlSvr := minioEnv.Host + minioEnv.EndpointUploadImage
	jsonData := map[string]string{"BucketName": req.BucketName, "Data": req.Data, "NameFile": req.NameFile, "ContentType": req.ContentType}

	data, err := https.HTTPUploadMinio(urlSvr, jsonData, minioEnv.HealthCheckKey)

	sugarLogger.Info("Level: Info",
		zap.String("SpanID: ", spanID),
		zap.String("Response: ", string(data)))

	json.Unmarshal(data, &res)

	return res, err
}
