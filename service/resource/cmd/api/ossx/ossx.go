package ossx

import (
	"go-zero-resource/common/errorx"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/model/gormx"
	"mime/multipart"
)

var (
	Enable  = 1
	Disable = 2

	Category_Minio   = 1
	Category_Qiniu   = 2
	Category_Ali     = 3
	Category_Tencent = 4
)

type OssTemplate interface {
	MakeBucket(bucketName string) error               // 创建存储桶
	RemoveBucket(bucketName string) error             // 删除存储桶
	BucketExists(bucketName string) (bool, error)     // 存储桶是否存在
	PutFile(file *multipart.FileHeader) (File, error) // 上传文件
	RemoveFile(bucketName string) error               // 删除文件
	RemoveFiles(bucketName []string) error            // 批量删除文件
}

type File struct {
	Link         string // 文件地址
	Domain       string // 域名地址
	Name         string // 文件名
	OriginalName string // 初始文件名
	AttachId     string // 附件表ID
}

type OssProperties struct {
	Enabled    bool                   // 是否启用
	name       string                 // 对象存储名称
	TenantMode bool                   // 是否开启租户模式
	Endpoint   string                 // 对象存储服务的URL
	AppId      string                 // 应用ID TencentCOS需要
	Region     string                 // 区域简称 TencentCOS需要
	AccessKey  string                 // Access key就像用户ID，可以唯一标识你的账户
	SecretKey  string                 // Secret key是你账户的密码
	BucketName string                 // 默认的存储桶名称
	Args       map[string]interface{} // 自定义属性
}

func getOss(tenantId, code string) (oss *gormx.ResourceOss, err error) {
	resourceOssQuery := gormx.ResourceOss{
		TenantId: tenantId,
	}
	if len(code) != 0 {
		resourceOssQuery.OssCode = code
	} else {
		resourceOssQuery.Status = Enable
	}
	tx := svc.CachedDb.Db.Where(&resourceOssQuery).First(&oss)
	if tx.RowsAffected == 0 {
		// todo 使用默认配置
		return nil, errorx.NewDefaultError("查找模板失败")
	} else {
		return
	}
}

func Template(TenantId, Code string) (ossTemplate OssTemplate, err error) {
	resourceOss, err := getOss(TenantId, Code)
	if err != nil {
		return nil, err
	} else {
		// todo 缓存template
		if resourceOss.Category == Category_Minio {
			ossTemplate = NewMinioTemplate(resourceOss)
		}
		return
	}
}
