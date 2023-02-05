package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
    "gorm.io/gorm"
)

type ActivityService struct {
}

// CreateActivity 创建Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService) CreateActivity(activity web.Activity) (err error) {
	err = global.GVA_DB.Create(&activity).Error
	return err
}

// DeleteActivity 删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService)DeleteActivity(activity web.Activity) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&web.Activity{}).Where("id = ?", activity.ID).Update("deleted_by", activity.DeletedBy).Error; err != nil {
              return err
        }
        if err = tx.Delete(&activity).Error; err != nil {
              return err
        }
        return nil
	})
	return err
}

// DeleteActivityByIds 批量删除Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService)DeleteActivityByIds(ids request.IdsReq,deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
	    if err := tx.Model(&web.Activity{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
            return err
        }
        if err := tx.Where("id in ?", ids.Ids).Delete(&web.Activity{}).Error; err != nil {
            return err
        }
        return nil
    })
	return err
}

// UpdateActivity 更新Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService)UpdateActivity(activity web.Activity) (err error) {
	err = global.GVA_DB.Save(&activity).Error
	return err
}

// GetActivity 根据id获取Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService)GetActivity(id uint) (activity web.Activity, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&activity).Error
	return
}

// GetActivityInfoList 分页获取Activity记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityService *ActivityService)GetActivityInfoList(info webReq.ActivitySearch) (list []web.Activity, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.Activity{})
    var activitys []web.Activity
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.Title != "" {
        db = db.Where("title LIKE ?","%"+ info.Title+"%")
    }
    if info.Type != nil {
        db = db.Where("type = ?",info.Type)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&activitys).Error
	return  activitys, total, err
}
