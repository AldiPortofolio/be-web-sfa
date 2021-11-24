package services

import (
	"encoding/json"
	"fmt"
	"log"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/hosts/minio"
	"ottosfa-api-web/models"
	"ottosfa-api-web/models/miniomodels"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"
)

// UploadImageService struct
type UploadImageService struct {
	General                   models.GeneralModel
	
}

// InitUploadImageService ...
func InitUploadImageService(gen models.GeneralModel) *UploadImageService {
	return &UploadImageService{
		General:          gen,
		
	}
}

// Upload ...
func (service *UploadImageService) Upload(req miniomodels.UploadReq) models.Response {
	fmt.Println(">>> UploadImage - Upload <<<")
	sugarLogger := service.General.OttoZaplog
	sugarLogger.Info("UploadImage: Upload",
		zap.Any("req", req))
	span, _ := opentracing.StartSpanFromContext(service.General.Context, "UploadImage: upload")
	defer span.Finish()

	var res models.Response

	data, err := minio.InitMinioHost().Send(req, constants.MinioUpload)
	if err != nil {
		log.Println("err -> ", err)
		res.Meta.Code = 400
		res.Meta.Message = constants.EC_FAIL_SEND_TO_HOST_DESC
		return res
	}

	resMinio := miniomodels.UploadRes{}
	json.Unmarshal(data, &resMinio)
	fmt.Println(string(data))

	res.Data = resMinio
	res.Meta.Code = 200
	res.Meta.Message = constants.ERR_SUCCESS_MSG

	sugarLogger.Info("Response", zap.Any("res", res))

	return res
}


