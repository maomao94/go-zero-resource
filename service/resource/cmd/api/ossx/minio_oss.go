package ossx

import (
	"go-zero-resource/service/resource/model/gormx"
	"mime/multipart"

	"github.com/minio/minio-go"
)

type MinioTemplate struct {
	OssTemplate
	client        *minio.Client // Minio客户端
	ossProperties OssProperties // 配置参数
}

func (m *MinioTemplate) MakeBucket(bucketName string) error {
	return m.client.MakeBucket(bucketName, "us-east-1")
}

func (m *MinioTemplate) BucketExists(bucketName string) (bool, error) {
	return m.client.BucketExists(bucketName)
}

func (m *MinioTemplate) PutFile(file *multipart.FileHeader) (File, error) {
	return File{}, nil
}

func NewMinioTemplate(Oss *gormx.ResourceOss) *MinioTemplate {
	ossProperties := OssProperties{
		Endpoint:   Oss.Endpoint,
		AccessKey:  Oss.AccessKey,
		SecretKey:  Oss.SecretKey,
		BucketName: Oss.BucketName,
		Args:       nil,
	}
	// 初使化 minio client对象。
	minioClient, _ := minio.New(Oss.Endpoint, Oss.AccessKey, Oss.SecretKey, false)
	return &MinioTemplate{
		client:        minioClient,
		ossProperties: ossProperties,
	}
}
