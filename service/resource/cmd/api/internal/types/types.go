// Code generated by goctl. DO NOT EDIT.
package types

type Oss struct {
	TenantId   string `json:"tenantId,optional"`   // 租户ID
	Category   int    `json:"category,optional"`   // 所属分类
	OssCode    string `json:"ossCode,optional"`    // 资源编号
	Endpoint   string `json:"endpoint,optional"`   // 资源地址
	AccessKey  string `json:"accessKey,optional"`  // accessKey
	SecretKey  string `json:"secretKey,optional"`  // secretKey
	BucketName string `json:"bucketName,optional"` // 空间名
	AppId      string `json:"appId,optional"`      // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`     // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`     // 所属分类
	Status     int    `json:"status,optional"`     // 状态
}

type OssCreate struct {
	TenantId   string `json:"tenantId,optional"`        // 租户ID
	Category   int    `json:"category,options=1|2|3|4"` // 所属分类
	OssCode    string `json:"ossCode,optional"`         // 资源编号
	Endpoint   string `json:"endpoint,optional"`        // 资源地址
	AccessKey  string `json:"accessKey,optional"`       // accessKey
	SecretKey  string `json:"secretKey,optional"`       // secretKey
	BucketName string `json:"bucketName,optional"`      // 空间名
	AppId      string `json:"appId,optional"`           // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`          // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`          // 所属分类
	Status     int    `json:"status,optional"`          // 状态
}

type OssUpdate struct {
	Id         uint   `json:"id"`                                // 主键ID
	TenantId   string `json:"tenantId,optional"`                 // 租户ID
	Category   int    `json:"category,options=1|2|3|4,optional"` // 所属分类
	OssCode    string `json:"ossCode,optional"`                  // 资源编号
	Endpoint   string `json:"endpoint,optional"`                 // 资源地址
	AccessKey  string `json:"accessKey,optional"`                // accessKey
	SecretKey  string `json:"secretKey,optional"`                // secretKey
	BucketName string `json:"bucketName,optional"`               // 空间名
	AppId      string `json:"appId,optional"`                    // 应用ID TencentCOS需要
	Region     string `json:"region,optional"`                   // 地域简称 TencentCOS需要
	Remark     string `json:"remark,optional"`                   // 所属分类
	Status     int    `json:"status,optional"`                   // 状态
}

type OssDelete struct {
	Id uint `json:"id"` // 主键ID
}

type OssListReq struct {
	PageInfo
	Oss
}

type MakeBucketReq struct {
	BucketName string `json:"bucketName"` // 存储桶名称
}

type EmptyResult struct {
}

type BaseResult struct {
	Id uint `json:"id"` // 主键ID
}

type TenantResult struct {
	BaseResult
	TenantId uint `json:"tenantId"` // 租户ID
}

type PageInfo struct {
	Page     int `json:"page,range=[1:100],default=1"`     // 页码
	PageSize int `json:"pageSize,range=[1:100],default=1"` // 每页大小
}

type PageResult struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`     // 页码
	PageSize int         `json:"pageSize"` // 每页大小
}
