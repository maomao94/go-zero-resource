package utils

var (
	IdVerify       = Rules{"Id": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
