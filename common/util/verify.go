package utils

var (
	PageInfoVerify   = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	BucketNameVerify = Rules{"BucketName": {NotEmpty()}}
)
