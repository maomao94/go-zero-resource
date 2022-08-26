package ossx

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/minio/minio-go"
	"gtw/model"
	"mime/multipart"
	"path"
	"strings"
	"sync"
	"time"
)

var (
	Category_Minio   int64 = 1
	Category_Qiniu   int64 = 2
	Category_Ali     int64 = 3
	Category_Tencent int64 = 4

	templatePool = make(map[string]OssTemplate)
	ossPool      = make(map[string]*model.TOss)
	lock         sync.Mutex
)

type OssTemplate interface {
	MakeBucket(tenantId, bucketName string) error                                                    // 创建存储桶
	RemoveBucket(tenantId, bucketName string) error                                                  // 删除存储桶
	StatFile(tenantId, bucketName, filename string) (*OssFile, error)                                // 获取文件信息
	BucketExists(tenantId, bucketName string) (bool, error)                                          // 存储桶是否存在
	PutFile(tenantId, bucketName string, fileHeader *multipart.FileHeader) (*File, error)            // 上传文件
	PutFileStream(tenantId, bucketName, filename, contentType string, stream *[]byte) (*File, error) // 上传文件
	GetObject(tenantId, bucketName, filename string) (*minio.Object, error)                          // 上传文件
	RemoveFile(tenantId, bucketName, filename string) error                                          // 删除文件
	RemoveFiles(tenantId string, bucketName string, filenames []string) error                        // 批量删除文件
}

var _ OssTemplate = (*MinioTemplate)(nil)

type OssRule struct {
	tenantMode bool
}

func (o *OssRule) bucketName(tenantId, bucketName string) string {
	prefix := ""
	if o.tenantMode {
		prefix = tenantId + "-"
	}
	return prefix + bucketName
}

func (o *OssRule) filename(originalFileName string) string {
	u, _ := uuid.NewUUID()
	return "upload" + "/" + time.Now().Format("20060102") + "/" +
		strings.Replace(fmt.Sprintf("%s", u), "-", "", -1) +
		path.Ext(originalFileName)
}

type File struct {
	Link         string // 文件地址
	Domain       string // 域名地址
	Name         string // 文件名
	OriginalName string // 初始文件名
	AttachId     string // 附件表ID
}

type OssFile struct {
	Link        string    // 文件地址
	Name        string    // 文件名
	Length      int64     // 文件大小
	PutTime     time.Time // 文件上传时间
	ContentType string    // 文件contentType
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

type GetOssFn func(tenantId, code string) (oss *model.TOss, err error)

func Template(TenantId, Code string, tenantMode bool, getOss GetOssFn) (ossTemplate OssTemplate, err error) {
	oss, err := getOss(TenantId, Code)
	ossCached := ossPool[TenantId]
	ossTemplate = templatePool[TenantId]
	if err != nil {
		return nil, err
	} else {
		if ossCached == nil || ossTemplate == nil ||
			(oss.Endpoint != ossCached.Endpoint) ||
			(oss.AccessKey != ossCached.AccessKey) {
			lock.Lock()
			defer lock.Unlock()
			if ossCached == nil || ossTemplate == nil ||
				(oss.Endpoint != ossCached.Endpoint) ||
				(oss.AccessKey != ossCached.AccessKey) {
				ossRule := OssRule{}
				if tenantMode {
					ossRule = OssRule{
						tenantMode: true,
					}
				} else {
					ossRule = OssRule{
						tenantMode: false,
					}
				}
				if oss.Category == Category_Minio {
					ossTemplate = NewMinioTemplate(oss, ossRule)
				} else {
					return nil, errors.New("oss type error")
				}
				templatePool[TenantId] = ossTemplate
				ossPool[TenantId] = oss
				return
			}
		}
		return
	}
}
