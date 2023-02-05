package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type IntroduceService struct {
}

// CreateIntroduce 创建Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService) CreateIntroduce(introduce web.Introduce) (err error) {
	err = global.GVA_DB.Create(&introduce).Error
	return err
}

// DeleteIntroduce 删除Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService)DeleteIntroduce(introduce web.Introduce) (err error) {
	err = global.GVA_DB.Delete(&introduce).Error
	return err
}

// DeleteIntroduceByIds 批量删除Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService)DeleteIntroduceByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Introduce{},"id in ?",ids.Ids).Error
	return err
}

// UpdateIntroduce 更新Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService)UpdateIntroduce(introduce web.Introduce) (err error) {
	err = global.GVA_DB.Save(&introduce).Error
	return err
}

// GetIntroduce 根据id获取Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService)GetIntroduce(id uint) (introduce web.Introduce, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&introduce).Error
	return
}

// GetIntroduceInfoList 分页获取Introduce记录
// Author [piexlmax](https://github.com/piexlmax)
func (introduceService *IntroduceService)GetIntroduceInfoList(info webReq.IntroduceSearch) (list []web.Introduce, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.Introduce{})
    var introduces []web.Introduce
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ActivityId != nil {
        db = db.Where("activity_id = ?",info.ActivityId)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&introduces).Error
	return  introduces, total, err
}
