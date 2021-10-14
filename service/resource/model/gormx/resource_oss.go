// 自动生成模板ResourceOss
package gormx

import (
	"go-zero-resource/common/api"
	"go-zero-resource/common/orm"
)

// ResourceOss 结构体
// 如果含有time.Time 请自行import time包
type ResourceOss struct {
	orm.GVA_MODEL
	TenantId   string `json:"tenantId" form:"tenantId" gorm:"column:tenant_id;comment:租户ID;type:varchar(12);default:000000"`
	Category   *int   `json:"category" form:"category" gorm:"column:category;comment:分类;type:int"`
	OssCode    string `json:"ossCode" form:"ossCode" gorm:"column:oss_code;comment:资源编号;type:varchar(32);"`
	Endpoint   string `json:"endpoint" form:"endpoint" gorm:"column:endpoint;comment:资源地址;type:varchar(255);"`
	AccessKey  string `json:"accessKey" form:"accessKey" gorm:"column:access_key;comment:accessKey;type:varchar(255);"`
	SecretKey  string `json:"secretKey" form:"secretKey" gorm:"column:secret_key;comment:secretKey;type:varchar(255);"`
	BucketName string `json:"bucketName" form:"bucketName" gorm:"column:bucket_name;comment:空间名;type:varchar(255);"`
	AppId      string `json:"appId" form:"appId" gorm:"column:app_id;comment:应用ID;type:varchar(255);"`
	Region     string `json:"region" form:"region" gorm:"column:region;comment:地域简称;type:varchar(255);"`
	Remark     string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;type:varchar(255);"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;type:int"`
	IsDeleted  *int   `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否已删除;type:int"`
}

// TableName ResourceOss 表名
func (ResourceOss) TableName() string {
	return "resource_oss"
}

type ResourceOssSearch struct {
	ResourceOss
	api.PageInfo
}
