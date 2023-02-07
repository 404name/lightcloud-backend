// 自动生成模板Comment
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Comment 结构体
type Comment struct {
	global.GVA_MODEL
	Content     string `json:"content" form:"content" gorm:"column:content;comment:内容;"`
	Describe    string `json:"describe" form:"describe" gorm:"column:describe;comment:描述(a评论了xxx);size:255;"`
	IsAnonymous *bool  `json:"isAnonymous" form:"isAnonymous" gorm:"column:is_anonymous;comment:是否匿名（0否，1是）;"`
	Rate        *int   `json:"rate" form:"rate" gorm:"column:rate;comment:评分;size:10;"`
	ReferId     *int   `json:"referId" form:"referId" gorm:"column:referId;comment:引用回复;size:10;"`
	Tags        string `json:"tags" form:"tags" gorm:"column:tags;comment:标签(已毕业、社团成员);size:255;"`
	ToId        *int   `json:"toId" form:"toId" gorm:"column:to_id;comment:评论对象id;size:10;"`
	ToType      *int   `json:"toType" form:"toType" gorm:"column:to_type;comment:评论对象类别(文章、活动、用户、留言);size:10;"`
	UserId      *int   `json:"userId" form:"userId" gorm:"column:userId;comment:发表者id;size:10;"`
	VoteDownNum *int   `json:"voteDownNum" form:"voteDownNum" gorm:"column:voteDownNum;comment:;size:10;"`
	VoteUpNum   *int   `json:"voteUpNum" form:"voteUpNum" gorm:"column:voteUpNum;comment:;size:10;"`
	Extra       string `json:"extra" form:"extra" gorm:"column:extra;comment:附加信息;size:255;"`
}

// TableName Comment 表名
func (Comment) TableName() string {
	return "web_comment"
}
