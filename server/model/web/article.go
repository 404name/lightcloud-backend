// 自动生成模板Article
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Article 结构体
type Article struct {
	global.GVA_MODEL
	CategoryId *int   `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类;size:10;"`
	Content    string `json:"content" form:"content" gorm:"type:text;column:content;comment:内容;"`
	Cover      string `json:"cover" form:"cover" gorm:"column:cover;comment:封面;size:255;"`
	IsTop      *bool  `json:"isTop" form:"isTop" gorm:"column:is_top;comment:是否置顶;"`
	Status     *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	Title      string `json:"title" form:"title" gorm:"column:title;comment:标题;size:255;"`
	Uid        *int   `json:"uid" form:"uid" gorm:"column:uid;comment:用户ID;size:10;"`
	Extra      string `json:"extra" form:"extra" gorm:"column:extra;comment:附加信息;size:255;"`
	ViewNum    *int   `json:"viewNum" form:"viewNum" gorm:"column:view_num;comment:阅读数;size:10;"`
}

// TableName Article 表名
func (Article) TableName() string {
	return "web_article"
}
