package rose

import (
	"encoding/json"
	"fmt"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
	https "ottosfa-api-web/utils/http"
)

// Env ..
type Env struct {
	Host                    string `envconfig:"ROSE_API_SERVICE_HOST" default:"http://13.228.25.85:8899/rose-api-service/v0.0.1"`
	EndpointLookUpGroup     string `envconfig:"ROSE_API_SERVICE_ENDPOINT_LOOK_UP_GROUP" default:"/lookup/lookupgroup"`
	EndpointUserCategory    string `envconfig:"ROSE_API_SERVICE_ENDPOINT_LOOK_UP_GROUP" default:"/user-category/all"`
	EndpointMerchantGroup   string `envconfig:"ROSE_API_SERVICE_ENDPOINT_LOOK_UP_GROUP" default:"/merchant-group/find-by-type"`
	KeyAppId                string `envconfig:"ROSE_API_SERVICE_ENDPOINT_KEY_APP_ID" default:"3"`
	HealthCheckKey          string `envconfig:"ROSE_API_SERVICE_HEALTH_CHECK_KEY" default:"OTTOSFA-API-APK_HEALTH_CHECK:ROSE_API_SERVICE"`
}

var (
	roseEnv Env
)

// init ..
func init() {
	err := envconfig.Process("ROSE API SERVICE", &roseEnv)
	if err != nil {
		fmt.Println("Failed to get ROSE API SERVICE env:", err)
	}
}

// LookUpGroup ..
func LookUpGroup(lookupGroup string) ([]LookUpGroupResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var lookUpGroupResponse []LookUpGroupResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointLookUpGroup

	jsonData := map[string]string{"lookupGroup": lookupGroup}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return lookUpGroupResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return lookUpGroupResponse, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return lookUpGroupResponse, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &lookUpGroupResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return lookUpGroupResponse, err
	}

	return lookUpGroupResponse, nil
}

// UserCategory ..
func UserCategory() ([]UserCategoryResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var userCategoryResponse []UserCategoryResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointUserCategory

	data, err := https.HTTPGet(urlSvr)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return userCategoryResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return userCategoryResponse, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return userCategoryResponse, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &userCategoryResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return userCategoryResponse, err
	}

	return userCategoryResponse, nil
}

// MerchantGroupByMerchantType ..
func MerchantGroupByMerchantType(merchantType string) ([]MerchantGroupResponse, error) {
	sugarLogger := ottologger.GetLogger()
	var res Response
	var merchantGroupResponse []MerchantGroupResponse

	urlSvr := roseEnv.Host + roseEnv.EndpointMerchantGroup

	jsonData := map[string]string{"merchantType": merchantType}
	data, err := https.HTTPPost(urlSvr, jsonData)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return merchantGroupResponse, err
	}

	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return merchantGroupResponse, err
	}

	if res.Rc == "04" {
		sugarLogger.Error("Level: Error", zap.String("Failed to get data: ", errors.New(res.Msg).Error()))
		return merchantGroupResponse, errors.New(res.Msg)
	}

	bdata, _ := json.Marshal(res.Data)
	if err = json.Unmarshal(bdata, &merchantGroupResponse); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response data: ", err.Error()))
		return merchantGroupResponse, err
	}

	return merchantGroupResponse, nil
}