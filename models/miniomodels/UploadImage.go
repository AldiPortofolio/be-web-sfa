package miniomodels

// UploadReq ..
type UploadReq struct {
	BucketName  string
	Data        string
	NameFile    string
	ContentType string
}

// UploadRes ..
type UploadRes struct {
	Url      string
	NameFile string
	Rc       string
	Message  string
}
