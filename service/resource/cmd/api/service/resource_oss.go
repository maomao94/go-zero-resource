package service

import (
	"fmt"
	"go-zero-resource/common/api"
	"go-zero-resource/common/errorx"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/model/gormx"

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

func (resourceOssService *ResourceOssService) CreateResourceOss(resourceOss gormx.ResourceOss) (err error) {
	err = svc.CachedDb.Db.Create(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOss(resourceOss gormx.ResourceOss) (err error) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, resourceOss.ID)
	_, err = svc.CachedDb.Exec(func(db *gorm.DB) (int64, error) {
		tx := svc.CachedDb.Db.Delete(&resourceOss)
		return tx.RowsAffected, tx.Error
	}, resourceOssIdKey)
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOssByIds(ids api.IdsReq) (err error) {
	err = svc.CachedDb.Db.Delete(&[]gormx.ResourceOss{}, "id in ?", ids.Ids).Error
	return err
}

func (resourceOssService *ResourceOssService) UpdateResourceOss(resourceOss gormx.ResourceOss) (err error) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, resourceOss.ID)
	_, err = svc.CachedDb.Exec(func(db *gorm.DB) (int64, error) {
		tx := svc.CachedDb.Db.Updates(&resourceOss)
		return tx.RowsAffected, tx.Error
	}, resourceOssIdKey)
	return err
}

func (resourceOssService *ResourceOssService) GetResourceOss(id uint) (err error, resourceOss gormx.ResourceOss) {
	// 使用缓存
	resourceOssIdKey := fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, id)
	err = svc.CachedDb.QueryRow(&resourceOss, resourceOssIdKey, func(db *gorm.DB, v interface{}) error {
		return svc.CachedDb.Db.Where("id = ?", id).First(&resourceOss).Error
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

func (resourceOssService *ResourceOssService) GetOss(tenantId, code string) (err error, resourceOss gormx.ResourceOss) {
	// 生成基于索引的key
	indexKey := fmt.Sprintf("%s%v-%v", cacheResourceOssCodePrefix, tenantId, code)
	// 使用缓存
	err = svc.CachedDb.QueryRowIndex(&resourceOss, indexKey,
		// 基于主键生成完整数据缓存的key
		func(primary interface{}) string {
			return fmt.Sprintf("%s%v", cacheResourceOssIdPrefix, primary)
		},
		// 基于索引的DB查询方法
		func(db *gorm.DB, v interface{}) (interface{}, error) {
			resourceOssQuery := gormx.ResourceOss{
				TenantId: tenantId,
			}
			if len(code) != 0 {
				resourceOssQuery.OssCode = code
			} else {
				resourceOssQuery.Status = enable
			}
			tx := svc.CachedDb.Db.Where(&resourceOssQuery).First(&v)
			if tx.RowsAffected == 0 {
				return nil, gorm.ErrRecordNotFound
			} else {
				return resourceOss.ID, nil
			}
		},
		// 基于主键的DB查询方法
		func(db *gorm.DB, v, primary interface{}) error {
			return svc.CachedDb.Db.Where("id = ?", primary).First(&resourceOss).Error
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

func (resourceOssService *ResourceOssService) GetResourceOssInfoList(info gormx.ResourceOssSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := svc.CachedDb.Db.Model(&gormx.ResourceOss{})
	var resourceOsss []gormx.ResourceOss
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Category != 0 {
		db = db.Where("category = ?", info.Category)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&resourceOsss).Error
	return err, resourceOsss, total
}
