// 自动生成模板Introduce
package web

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// Introduce 结构体
type Introduce struct {
      global.GVA_MODEL
      ActivityId  *int `json:"activityId" form:"activityId" gorm:"column:activity_id;comment:活动id;size:10;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:内推码;size:255;"`
      Post  string `json:"post" form:"post" gorm:"column:post;comment:生成海报;size:255;"`
      UserId  *int `json:"userId" form:"userId" gorm:"column:user_id;comment:用户id(组织默认生成);size:10;"`
}


// TableName Introduce 表名
func (Introduce) TableName() string {
  return "web_introduce"
}

