// 自动生成模板Activity
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

// Activity 结构体
type Activity struct {
	global.GVA_MODEL
	BeginTime *utils.JSONTime `json:"beginTime" form:"beginTime" gorm:"column:begin_time;comment:活动开始时间;"`
	Content   string          `json:"content" form:"content" gorm:"type:text;column:content;comment:内容;"`
	Cover     string          `json:"cover" form:"cover" gorm:"column:cover;comment:封面;size:255;"`
	EndTime   *utils.JSONTime `json:"endTime" form:"endTime" gorm:"column:end_time;comment:活动结束时间;"`
	Status    string          `json:"status" form:"status" gorm:"column:status;comment:活动状态(json);size:255;"`
	Title     string          `json:"title" form:"title" gorm:"column:title;comment:活动名称;size:255;"`
	Type      *int            `json:"type" form:"type" gorm:"column:type;comment:活动类型(0招新，1待开发);size:10;"`
	UserId    *int            `json:"userId" form:"userId" gorm:"column:user_id;comment:发起组织id;size:10;"`
	Extra     string          `json:"extra" form:"extra" gorm:"column:extra;comment:附加信息;size:255;"`
	CreatedBy uint            `gorm:"column:created_by;comment:创建者"`
	UpdatedBy uint            `gorm:"column:updated_by;comment:更新者"`
	DeletedBy uint            `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Activity 表名
func (Activity) TableName() string {
	return "web_activity"
}
