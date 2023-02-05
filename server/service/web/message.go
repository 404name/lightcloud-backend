package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
)

type MessageService struct {
}

// CreateMessage 创建Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService) CreateMessage(message web.Message) (err error) {
	err = global.GVA_DB.Create(&message).Error
	return err
}

// DeleteMessage 删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)DeleteMessage(message web.Message) (err error) {
	err = global.GVA_DB.Delete(&message).Error
	return err
}

// DeleteMessageByIds 批量删除Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)DeleteMessageByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Message{},"id in ?",ids.Ids).Error
	return err
}

// UpdateMessage 更新Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)UpdateMessage(message web.Message) (err error) {
	err = global.GVA_DB.Save(&message).Error
	return err
}

// GetMessage 根据id获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)GetMessage(id uint) (message web.Message, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&message).Error
	return
}

// GetMessageInfoList 分页获取Message记录
// Author [piexlmax](https://github.com/piexlmax)
func (messageService *MessageService)GetMessageInfoList(info webReq.MessageSearch) (list []web.Message, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&web.Message{})
    var messages []web.Message
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
    if info.FromId != nil {
        db = db.Where("from_id = ?",info.FromId)
    }
    if info.FromType != nil {
        db = db.Where("from_type = ?",info.FromType)
    }
    if info.Type != "" {
        db = db.Where("type = ?",info.Type)
    }
    if info.UserId != nil {
        db = db.Where("user_id = ?",info.UserId)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&messages).Error
	return  messages, total, err
}
