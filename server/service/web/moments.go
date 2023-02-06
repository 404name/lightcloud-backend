package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type MomentsService struct {
}

// CreateMoments 创建Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) CreateMoments(moments web.Moments) (err error) {
	err = global.GVA_DB.Create(&moments).Error
	return err
}

// DeleteMoments 删除Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) DeleteMoments(moments web.Moments) (err error) {
	err = global.GVA_DB.Delete(&moments).Error
	return err
}

// DeleteMomentsByIds 批量删除Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) DeleteMomentsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Moments{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateMoments 更新Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) UpdateMoments(moments web.Moments) (err error) {
	err = global.GVA_DB.Save(&moments).Error
	return err
}

// GetMoments 根据id获取Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) GetMoments(id uint) (moments web.Moments, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&moments).Error
	return
}

// GetMomentsInfoList 分页获取Moments记录
// Author [piexlmax](https://github.com/piexlmax)
func (momentsService *MomentsService) GetMomentsInfoList(info webReq.MomentsSearch) (list []web.Moments, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.Moments{})
	var momentss []web.Moments
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.FromId != nil {
		db = db.Where("from_id = ?", info.FromId)
	}
	if info.FromType != nil {
		db = db.Where("from_type = ?", info.FromType)
	}
	if info.Type != "" {
		db = db.Where("type = ?", info.Type)
	}
	if info.UserId != nil {
		db = db.Where("user_id = ?", info.UserId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Order("id desc").Limit(limit).Offset(offset).Find(&momentss).Error
	return momentss, total, err
}
