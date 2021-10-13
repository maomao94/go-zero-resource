package util

var (
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	IdVerify       = Rules{"ID": {NotEmpty()}}
)
