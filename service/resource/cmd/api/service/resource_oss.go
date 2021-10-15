package service

import (
	"go-zero-resource/common/api"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/model/gormx"
)

var (
	cacheResourceOssIdPrefix = "cache:resourceOss:id:"
)

type ResourceOssService struct {
}

var ResourceOssApp = new(ResourceOssService)

func (resourceOssService *ResourceOssService) CreateResourceOss(resourceOss gormx.ResourceOss) (err error) {
	err = svc.CachedDb.Db.Create(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOss(resourceOss gormx.ResourceOss) (err error) {
	err = svc.CachedDb.Db.Delete(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOssByIds(ids api.IdsReq) (err error) {
	err = svc.CachedDb.Db.Delete(&[]gormx.ResourceOss{}, "id in ?", ids.Ids).Error
	return err
}

func (resourceOssService *ResourceOssService) UpdateResourceOss(resourceOss gormx.ResourceOss) (err error) {
	err = svc.CachedDb.Db.Save(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) GetResourceOss(id uint) (err error, resourceOss gormx.ResourceOss) {
	// 修改成带缓存
	//err = svc.CachedDb.Db.Where("id = ?", id).First(&resourceOss).Error
	return
}

func (resourceOssService *ResourceOssService) GetResourceOssInfoList(info gormx.ResourceOssSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := svc.CachedDb.Db.Model(&gormx.ResourceOss{})
	var resourceOsss []gormx.ResourceOss
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&resourceOsss).Error
	return err, resourceOsss, total
}
