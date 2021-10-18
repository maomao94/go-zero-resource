package oss

import (
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/cmd/api/internal/types"
	"go-zero-resource/service/resource/model/gormx"
	"mime/multipart"

	"github.com/minio/minio-go"
)

var (
	enable  = 1
	disable = 2
)

type OssTemplate interface {
	MakeBucket(bucketName string) error               // 创建存储桶
	RemoveBucket(bucketName string) error             // 删除存储桶
	BucketExists(bucketName string) error             // 存储桶是否存在
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
	enabled    bool                   // 是否启用
	name       string                 // 对象存储名称
	tenantMode bool                   // 是否开启租户模式
	endpoint   string                 // 对象存储服务的URL
	appId      string                 // 应用ID TencentCOS需要
	region     string                 // 区域简称 TencentCOS需要
	accessKey  string                 // Access key就像用户ID，可以唯一标识你的账户
	secretKey  string                 // Secret key是你账户的密码
	bucketName string                 // 默认的存储桶名称
	args       map[string]interface{} // 自定义属性
}

type OssBuilder struct{}

func (*OssBuilder) getOss(tenantId, code string) (oss *gormx.ResourceOss) {
	resourceOssQuery := gormx.ResourceOss{
		TenantId: tenantId,
	}
	if len(code) != 0 {
		resourceOssQuery.OssCode = code
	} else {
		resourceOssQuery.Status = enable
	}
	svc.CachedDb.Db.Where(&resourceOssQuery).First(&oss)
	if svc.CachedDb.Db.RowsAffected == 0 {
		// todo 使用默认配置
		return nil
	} else {
		return
	}
}

func (*OssBuilder) template(code string) {

}

type MinioTemplate struct {
	OssTemplate
	client        *minio.Client // Minio客户端
	ossProperties OssProperties // 配置参数
}

func (*MinioTemplate) Template(oss types.Oss) *MinioTemplate {
	ossProperties := OssProperties{
		endpoint:   oss.Endpoint,
		accessKey:  oss.AccessKey,
		secretKey:  oss.SecretKey,
		bucketName: oss.BucketName,
		args:       nil,
	}
	// 初使化 minio client对象。
	minioClient, _ := minio.New(oss.Endpoint, oss.AccessKey, oss.SecretKey, false)
	return &MinioTemplate{
		client:        minioClient,
		ossProperties: ossProperties,
	}
}

func (m *MinioTemplate) MakeBucket(bucketName string) error {
	return m.client.MakeBucket(bucketName, "")
}

func (m *MinioTemplate) PutFile(file *multipart.FileHeader) (File, error) {
	return File{}, nil
}
