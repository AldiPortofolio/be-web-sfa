package ottosfanotif

import (
	"encoding/json"
	"fmt"
	https "ottosfa-api-web/utils/http"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
	ottologger "ottodigital.id/library/logger"
)

// Env ..
type Env struct {
	Host                string `envconfig:"OTTO_SFA_NOTIF_HOST" default:"http://34.101.141.240:8046/ottosfa-api-notif/v1"`
	EndpointCreateNotif string `envconfig:"OTTO_SFA_NOTIF_ENDPOINT_CREATE_NOTIF" default:"/notification/create"`
}

var (
	sfaEnv Env
)

// init ..
func init() {
	err := envconfig.Process("SFA API NOTIF", &sfaEnv)
	if err != nil {
		fmt.Println("Failed to get SFA API NOTIF env:", err)
	}
}

// SendNotif ..
func SendNotif(req *ReqCreateNotif, token string) error {
	sugarLogger := ottologger.GetLogger()
	var res Response

	urlSvr := sfaEnv.Host + sfaEnv.EndpointCreateNotif
	auth := "Bearer " + token
	data, err := https.HTTPPostWithHeader(urlSvr, req, auth)
	if err != nil {
		sugarLogger.Error("Level: Error", zap.String("Error: ", err.Error()))
		return err
	}

	fmt.Println(string(data))
	if err = json.Unmarshal(data, &res); err != nil {
		sugarLogger.Error("Level: Error", zap.String("Failed to unmarshaling response: ", err.Error()))
		return err
	}

	return nil
}
