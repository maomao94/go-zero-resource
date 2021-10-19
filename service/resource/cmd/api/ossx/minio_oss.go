package ossx

import (
	"go-zero-resource/service/resource/model/gormx"
	"mime/multipart"

	"github.com/minio/minio-go"
)

type MinioTemplate struct {
	client        *minio.Client // Minio客户端
	ossProperties OssProperties // 配置参数
	ossRule       OssRule
}

func (m MinioTemplate) MakeBucket(tenantId, bucketName string) error {
	return m.client.MakeBucket(m.ossRule.bucketName(tenantId, bucketName), "")
}

func (m MinioTemplate) RemoveBucket(tenantId, bucketName string) error {
	panic("implement me")
}

func (m MinioTemplate) BucketExists(tenantId, bucketName string) (bool, error) {
	return m.client.BucketExists(m.ossRule.bucketName(tenantId, bucketName))
}

func (m MinioTemplate) PutFile(tenantId string, file *multipart.FileHeader) (*File, error) {
	f, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer f.Close()
	fileName := m.ossRule.fileName(file.Filename)
	_, err = m.client.PutObject(m.ossRule.bucketName(tenantId, m.ossProperties.BucketName),
		fileName, f, file.Size, minio.PutObjectOptions{
			ContentType: file.Header.Get("content-type"),
		})
	if err != nil {
		return nil, err
	} else {
		return &File{
			Link:         m.fileLink(tenantId, m.ossProperties.BucketName, fileName),
			Domain:       m.getOssHost(tenantId, m.ossProperties.BucketName),
			Name:         fileName,
			OriginalName: file.Filename,
		}, nil
	}
}

func (m MinioTemplate) RemoveFile(tenantId, bucketName string) error {
	panic("implement me")
}

func (m MinioTemplate) RemoveFiles(tenantId string, bucketName []string) error {
	panic("implement me")
}

func (m MinioTemplate) getOssHost(tenantId, bucketName string) string {
	return m.ossProperties.Endpoint + "/" + m.ossRule.bucketName(tenantId, bucketName)
}

func (m MinioTemplate) fileLink(tenantId, bucketName, fileName string) string {
	return m.ossProperties.Endpoint + "/" + m.ossRule.bucketName(tenantId, bucketName) + "/" + fileName
}

func NewMinioTemplate(Oss gormx.ResourceOss, ossRule OssRule) *MinioTemplate {
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
		ossRule:       ossRule,
	}
}
