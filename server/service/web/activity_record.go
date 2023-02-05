package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type ActivityRecordService struct {
}

// CreateActivityRecord 创建ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService) CreateActivityRecord(activityRecord web.ActivityRecord) (err error) {
	err = global.GVA_DB.Create(&activityRecord).Error
	return err
}

// DeleteActivityRecord 删除ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService)DeleteActivityRecord(activityRecord web.ActivityRecord) (err error) {
	err = global.GVA_DB.Delete(&activityRecord).Error
	return err
}

// DeleteActivityRecordByIds 批量删除ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService)DeleteActivityRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.ActivityRecord{},"id in ?",ids.Ids).Error
	return err
}

// UpdateActivityRecord 更新ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService)UpdateActivityRecord(activityRecord web.ActivityRecord) (err error) {
	err = global.GVA_DB.Save(&activityRecord).Error
	return err
}

// GetActivityRecord 根据id获取ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService)GetActivityRecord(id uint) (activityRecord web.ActivityRecord, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&activityRecord).Error
	return
}

// GetActivityRecordInfoList 分页获取ActivityRecord记录
// Author [piexlmax](https://github.com/piexlmax)
func (activityRecordService *ActivityRecordService)GetActivityRecordInfoList(info webReq.ActivityRecordSearch) (list []web.ActivityRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.ActivityRecord{})
    var activityRecords []web.ActivityRecord
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.ActivityId != nil {
        db = db.Where("activity_id = ?",info.ActivityId)
    }
    if info.Code != "" {
        db = db.Where("code = ?",info.Code)
    }
    if info.Status != "" {
        db = db.Where("status = ?",info.Status)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
        if info.StartScore != nil && info.EndScore != nil {
            db = db.Where("score BETWEEN ? AND ? ",info.StartScore,info.EndScore)
        }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&activityRecords).Error
	return  activityRecords, total, err
}
