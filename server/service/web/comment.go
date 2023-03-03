package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/web"
	webReq "github.com/flipped-aurora/gin-vue-admin/server/model/web/request"
	"gorm.io/gorm"
)

type CommentService struct {
}

// CreateComment 创建Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) CreateComment(comment web.Comment) (err error) {
	err = global.GVA_DB.Create(&comment).Error
	return err
}

// DeleteComment 删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteComment(comment web.Comment) (err error) {
	err = global.GVA_DB.Delete(&comment).Error
	return err
}

// DeleteCommentByIds 批量删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) DeleteCommentByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]web.Comment{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateComment 更新Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) UpdateComment(comment web.Comment) (err error) {
	err = global.GVA_DB.Save(&comment).Error
	return err
}

// GetComment 根据id获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetComment(id uint) (comment web.Comment, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&comment).Error
	return
}

// GetCommentInfoList 分页获取Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) GetCommentInfoList(info webReq.CommentSearch) (list []web.Comment, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&web.Comment{})
	var comments []web.Comment
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ToId != nil {
		db = db.Where("to_id = ?", info.ToId)
	}
	if info.ToType != nil {
		db = db.Where("to_type = ?", info.ToType)
	}
	if info.UserId != nil {
		db = db.Where("userId = ?", info.UserId)
	}
	if info.SortBy != "" {
		if info.SortBy == "hot" {
			db = db.Order("voteUpNum desc")
		} else {
			db = db.Order("id desc")
		}
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&comments).Error
	return comments, total, err
}

// DeleteCommentByIds 批量删除Comment记录
// Author [piexlmax](https://github.com/piexlmax)
func (commentService *CommentService) ApprovalComment(commentId int, approval bool) (err error) {

	var comment web.Comment
	if approval {
		return global.GVA_DB.First(&comment, commentId).Update("VoteUpNum", gorm.Expr("VoteUpNum + ?", 1)).Error
	} else {
		return global.GVA_DB.Where("id = ? AND VoteDownNum > ?", commentId, 0).First(&comment).Update("VoteDownNum", gorm.Expr("VoteDownNum - ?", 1)).Error
	}
}
