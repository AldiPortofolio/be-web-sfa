package utils

import (
	"encoding/json"
	"io/ioutil"
	"ottosfa-api-web/constants"
	"ottosfa-api-web/models"

	"github.com/astaxie/beego/logs"
)

var (
	// ListErrorCode ..
	ListErrorCode []models.MappingErrorCodes
)

func init() {
	RegisterErrorCode()
}

// RegisterErrorCode registering code
func RegisterErrorCode() bool {
	logs.Info("Register error code from json file")

	var b []byte
	var err error
	b, err = ioutil.ReadFile("error_code.json") // just pass the file name
	if err != nil {
		b, err = ioutil.ReadFile("../error_code.json") // just pass the file name
		if err != nil {
			logs.Error("Failed to read file error code json ", err)
		}
	}

	if json.Unmarshal(b, &ListErrorCode) != nil {
		logs.Error("Unmarshal [%v] or JSONErrorCode Failed : [%d]", err)
		return false
	}

	return true
}

// GetMetaResponse ..
func GetMetaResponse(key string) models.Meta {
	logs.Info("Get response by key", key)

	var meta models.Meta

	if key == constants.KeyResponseDefault {
		meta.Code = 400
		meta.Message = "Failed"
		meta.Status = false
	}

	if key == constants.KeyResponseSuccessful {
		meta.Code = 200
		meta.Message = "OK"
		meta.Status = true
	}

	if key == constants.KeyResponseFailed {
		meta.Code = 400
		meta.Message = "Failed"
		meta.Status = false
	}

	if key == constants.KeyResponseInvalidToken {
		meta.Code = 401
		meta.Message = "Invalid Token"
		meta.Status = false
	}

	for _, element := range ListErrorCode {
		if element.Key == key {
			meta.Status = element.Content.Status
			if element.Content.Status != true {
				meta.Status = false
			}
			meta.Code = element.Content.Code
			meta.Message = element.Content.Message

			return meta
		}
	}

	return meta
}

//func GetMessageFailedErrorReplace(res *models.Response, code int, err error) {
//	res.Meta.Code = code
//	res.Meta.Status = false
//	res.Meta.Message = err.Error()
//}
