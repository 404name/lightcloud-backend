// 自动生成模板ActivityRecord
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ActivityRecord 结构体
type ActivityRecord struct {
	global.GVA_MODEL
	ActivityId *int   `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动id;size:10;"`
	Code       string `json:"code" form:"code" gorm:"column:code;comment:填写的内推码;size:255;"`
	Status     string `json:"status" form:"status" gorm:"column:status;comment:参与状态;size:255;"`
	UserId     *int   `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id;size:10;"`

	Note     string    `json:"note" form:"note" gorm:"column:note;comment:;size:255;"`
	Score    *int      `json:"score" form:"score" gorm:"column:score;comment:;size:10;"`
	Activity *Activity `json:"activity,omitempty"`
}

// TableName ActivityRecord 表名
func (ActivityRecord) TableName() string {
	return "web_activity_record"
}
