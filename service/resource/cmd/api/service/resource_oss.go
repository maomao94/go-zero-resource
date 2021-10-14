package service

import (
	"go-zero-resource/common/api"
	"go-zero-resource/service/resource/cmd/api/internal/svc"
	"go-zero-resource/service/resource/mode"
)

type ResourceOssService struct {
}

func (resourceOssService *ResourceOssService) CreateResourceOss(resourceOss mode.ResourceOss) (err error) {
	err = svc.DB.Create(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOss(resourceOss mode.ResourceOss) (err error) {
	err = svc.DB.Delete(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) DeleteResourceOssByIds(ids api.IdsReq) (err error) {
	err = svc.DB.Delete(&[]mode.ResourceOss{}, "id in ?", ids.Ids).Error
	return err
}

func (resourceOssService *ResourceOssService) UpdateResourceOss(resourceOss mode.ResourceOss) (err error) {
	err = svc.DB.Save(&resourceOss).Error
	return err
}

func (resourceOssService *ResourceOssService) GetResourceOss(id uint) (err error, resourceOss mode.ResourceOss) {
	err = svc.DB.Where("id = ?", id).First(&resourceOss).Error
	return
}

func (resourceOssService *ResourceOssService) GetResourceOssInfoList(info mode.ResourceOssSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := svc.DB.Model(&mode.ResourceOss{})
	var resourceOsss []mode.ResourceOss
	// 如果有条件搜索 下方会自动创建搜索语句
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&resourceOsss).Error
	return err, resourceOsss, total
}
