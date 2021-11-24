package constants

const (
	// KeyResponsePending ..
	KeyResponsePending = "pending"

	// KeyResponseDefault ..
	KeyResponseDefault = "default"

	// KeyResponseFailed ..
	KeyResponseFailed = "failed"

	// KeyResponseSuccessful ..
	KeyResponseSuccessful = "successful"

	// KeyResponseInvalidToken ..
	KeyResponseInvalidToken = "invalid-token"

	ERR_UNMARSHAL     = "03"
	ERR_UNMARSHAL_MSG = "Error, unmarshall body Request"
	MinioUpload = "MINIO_UPLOAD"
	EC_FAIL_SEND_TO_HOST = "08";
	EC_FAIL_SEND_TO_HOST_DESC = "FAILED SEND TO HOST ";
	ERR_SUCCESS_MSG = "SUCCESS.."
)

const (
	HttpMethodGet    = "GET"
	HttpMethodPost   = "POST"
	HttpMethodPut    = "PUT"
	HttpMethodDelete = "DELETE"
)