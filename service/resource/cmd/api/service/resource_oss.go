package service

import (
	"fmt"
	"go-zero-resource/common/api"
	"go-zero-resource/common/errorx"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/model/gorm_model"

	"gorm.io/gorm"
)

var (
	enable = 1

	disable = 2

	cacheResourceOssIdPrefix = "cache:resourceOss:id:"

	cacheResourceOssCodePrefix = "cache:resourceOss:code:"
)

type ResourceOssService struct {
}

var ResourceOssApp = new(ResourceOssService)

func (resourceOssService *ResourceOssService) CreateResourceOss(resourceOss gorm_model.ResourceOss) (err error) {
	err = svc.CachedConn.Db.Create(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOss(resourceOss gorm_model.ResourceOss) (err error) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, resourceOss.ID)
	_, err = svc.CachedConn.Exec(func(db *gorm.DB) (int64, error) {
		tx := svc.CachedConn.Db.Delete(&resourceOss)
		return tx.RowsAffected, tx.Error
	}, resourceOssIdKey)
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOssByIds(ids api.IdsReq) (err error) {
	err = svc.CachedConn.Db.Delete(&[]gorm_model.ResourceOss{}, "id in ?", ids.Ids).Error
	return err
}

func (resourceOssService *ResourceOssService) UpdateResourceOss(resourceOss gorm_model.ResourceOss) (err error) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, resourceOss.ID)
	_, err = svc.CachedConn.Exec(func(db *gorm.DB) (int64, error) {
		tx := svc.CachedConn.Db.Updates(&resourceOss)
		return tx.RowsAffected, tx.Error
	}, resourceOssIdKey)
	return err
}

func (resourceOssService *ResourceOssService) GetResourceOss(id uint) (err error, resourceOss gorm_model.ResourceOss) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, id)
	err = svc.CachedConn.QueryRow(&resourceOss, resourceOssIdKey, func(db *gorm.DB, v interface{}) error {
		return svc.CachedConn.Db.Where("id = ?", id).First(&resourceOss).Error
	})
	// 格式化错误
	switch err {
	case nil:
		return nil, resourceOss
	case gorm.ErrRecordNotFound:
		return errorx.NewCodeError(errorx.NotFound), resourceOss
	default:
		return err, resourceOss
	}
}

func (resourceOssService *ResourceOssService) GetOss(tenantId, code string) (err error, resourceOss gorm_model.ResourceOss) {
	// 生成基于索引的key
	indexKey := fmt.Sprintf("%s%v-%v", cacheResourceOssCodePrefix, tenantId, code)
	// 使用缓存
	err = svc.CachedConn.QueryRowIndex(&resourceOss, indexKey,
		// 基于主键生成完整数据缓存的key
		func(primary interface{}) string {
			return fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, primary)
		},
		// 基于索引的DB查询方法
		func(db *gorm.DB, v interface{}) (interface{}, error) {
			resourceOssQuery := gorm_model.ResourceOss{
				TenantId: tenantId,
			}
			if len(code) != 0 {
				resourceOssQuery.OssCode = code
			} else {
				resourceOssQuery.Status = enable
			}
			tx := svc.CachedConn.Db.Where(&resourceOssQuery).First(&v)
			if tx.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			} else {
				return resourceOss.ID, nil
			}
		},
		// 基于主键的DB查询方法
		func(db *gorm.DB, v, primary interface{}) error {
			return svc.CachedConn.Db.Where("id = ?", primary).First(&resourceOss).Error
		})

	// 格式化错误
	switch err {
	case nil:
		return nil, resourceOss
	case gorm.ErrRecordNotFound:
		return errorx.NewCodeError(errorx.NotFound), resourceOss
	default:
		return err, resourceOss
	}
}

func (resourceOssService *ResourceOssService) GetResourceOssInfoList(info gorm_model.ResourceOssSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := svc.CachedConn.Db.Model(&gorm_model.ResourceOss{})
	var resourceOsss []gorm_model.ResourceOss
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Category != 0 {
		db = db.Where("category = ?", info.Category)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&resourceOsss).Error
	return err, resourceOsss, total
}
