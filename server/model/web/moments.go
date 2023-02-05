// 自动生成模板Moments
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Moments 结构体
type Moments struct {
	global.GVA_MODEL
	Content  string `json:"content" form:"content" gorm:"column:content;comment:动态内容;size:255;"`
	Cover    string `json:"cover" form:"cover" gorm:"column:cover;comment:封面;size:255;"`
	FromId   *int   `json:"fromId" form:"fromId" gorm:"column:from_id;comment:消息发起id;size:10;"`      // 跳转过去id
	FromType *int   `json:"fromType" form:"fromType" gorm:"column:from_type;comment:消息发起方;size:10;"` // 跳转过去类别
	IsTop    *bool  `json:"isTop" form:"isTop" gorm:"column:is_top;comment:是否置顶（0否, 1是）;"`
	Title    string `json:"title" form:"title" gorm:"column:title;comment:标题(AXX了你/通知);"`
	Type     string `json:"type" form:"type" gorm:"column:type;comment:类别(通知，动态，);size:255;"`
	UserId   *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"` // 谁发起的
}

// TableName Moments 表名
func (Moments) TableName() string {
	return "web_moments"
}
