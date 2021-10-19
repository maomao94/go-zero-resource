package utils

var (
	IdVerify       = Rules{"Id": {NotEmpty()}}
	TenantIdVerify = Rules{"TenantId": {NotEmpty()}}
	PageInfoVerify = Rules{"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
)
