package http

import (
	"crypto/tls"
	"net/http"
	"ottosfa-api-web/constants"
	"time"

	"github.com/astaxie/beego"
	"github.com/parnurzeal/gorequest"
)

var (
	debugClientHTTP bool
	instuition      string
	timeout         string
	retrybad        int
)

// Env ..
type Env struct {
	DebugClient bool   `envconfig:"DEBUG_CLIENT" default:"true"`
	Timeout     string `envconfig:"TIMEOUT" default:"60s"`
	RetryBad    int    `envconfig:"RETRY_BAD" default:"1"`
}

var (
	httpEnv Env
)


func init() {
	debugClientHTTP = beego.AppConfig.DefaultBool("debugClientHTTP", true)
	timeout = beego.AppConfig.DefaultString("timeout", "60s")
	retrybad = beego.AppConfig.DefaultInt("retrybad", 1)
}

// HTTPGet func
func HTTPGet(url string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Get(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPost func
func HTTPPost(url string, jsondata interface{}) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPUploadMinio ..
func HTTPUploadMinio(url string, jsondata interface{}, key string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	_, body, errs := reqagent.
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		Type("form-data").
		Send(jsondata).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}

	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader(url string, jsondata interface{}, token string) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(debugClientHTTP)
	timeout, _ := time.ParseDuration(timeout)
	//_ := errors.New("Connection Problem")
	if url[:5] == "https" {
		request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	}
	reqagent := request.Post(url)
	reqagent.Header.Set("Content-Type", "application/json")
	reqagent.Header.Add("Authorization", token)
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(retrybad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// HTTPPostWithHeader func
func HTTPPostWithHeader2(url string, jsondata interface{}, header http.Header) ([]byte, error) {
	request := gorequest.New()
	request.SetDebug(httpEnv.DebugClient)
	timeout, _ := time.ParseDuration(httpEnv.Timeout)
	//_ := errors.New("Connection Problem")
	// if url[:5] == "https" {
	// 	request.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	// }
	reqagent := request.Post(url)
	reqagent.Header = header
	_, body, errs := reqagent.
		Send(jsondata).
		Timeout(timeout).
		Retry(httpEnv.RetryBad, time.Second, http.StatusInternalServerError).
		End()
	if errs != nil {
		return []byte(body), errs[0]
	}
	return []byte(body), nil
}

// SendHttpRequest ..
func SendHttpRequest(method string, url string, header http.Header, body interface{}) ([]byte, error) {
	var data []byte
	var err error
	switch method {
	case constants.HttpMethodGet:
		data, err = HTTPGet(url)
		
	case constants.HttpMethodPost:
		data, err = HTTPPostWithHeader2(url, body, header)
		
	}
	return data, err
}
